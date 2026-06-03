// Command migrate transforms the raw output of openapi-generator-cli (Go
// client) into the form vrchatapi-go actually compiles against.
//
// Six patches are applied:
//
//  1. Drop "decoder.DisallowUnknownFields()" from every model's UnmarshalJSON
//     so the client tolerates extra fields the VRChat API returns.
//  2. Prefix every enum const with its type name ("PUBLIC" -> "ReleaseStatus_PUBLIC")
//     across both declarations and cross-file default-value references — bare
//     names collide because multiple enums share values like "public" / "group".
//  3. Add the missing "time" import to api_miscellaneous.go (the generator
//     forgets it even though GetSystemTime returns *time.Time).
//  4. Delete the three duplicate *UserPersistence* Api types and their methods
//     from api_worlds.go (kept canonically in api_users.go); they exist in both
//     because the spec tags those operations with both "users" and "worlds".
//  5. Relax UnmarshalJSON on every enum type so unknown values (e.g. a new
//     GroupPermissions entry VRChat ships post-generation) are accepted as-is
//     instead of failing the whole decode. IsValid() and New<T>FromValue keep
//     strict semantics for callers that explicitly want validation.
//  6. Replace "url.PathEscape" with "pathSegmentEscape" in api_*.go files so
//     literal parens (legal per RFC 3986 sub-delims) in VRChat instance IDs
//     survive into the wire path and don't trip Cloudflare's WAF. The helper
//     itself lives in the hand-written path_escape.go at the repo root.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	srcDir := flag.String("src", "generation/src", "directory of fresh OpenAPI generator output")
	outDir := flag.String("out", "generation/out", "directory to write migrated files")
	override := flag.Bool("override", false, "write directly to repo root instead of -out")
	flag.Parse()

	dst := *outDir
	if *override {
		// Repo root, relative to the directory generate.sh is invoked from.
		dst = filepath.Join("..", "..")
	}
	if err := run(*srcDir, dst); err != nil {
		fmt.Fprintln(os.Stderr, "migrate:", err)
		os.Exit(1)
	}
}

func run(srcDir, dstDir string) error {
	if err := os.MkdirAll(dstDir, 0o755); err != nil {
		return fmt.Errorf("mkdir %s: %w", dstDir, err)
	}

	fset := token.NewFileSet()
	files := map[string]*ast.File{} // basename -> AST

	entries, err := os.ReadDir(srcDir)
	if err != nil {
		return fmt.Errorf("read %s: %w", srcDir, err)
	}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".go") {
			continue
		}
		path := filepath.Join(srcDir, e.Name())
		f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			return fmt.Errorf("parse %s: %w", path, err)
		}
		files[e.Name()] = f
	}

	enumTable := buildEnumTable(files)

	for name, f := range files {
		prefixEnums(f, enumTable)
		relaxEnumUnmarshal(f, enumTable)
		if strings.HasPrefix(name, "api_") {
			rewritePathEscapeCalls(f)
		}
		switch name {
		case "api_miscellaneous.go":
			addImport(f, "time")
		case "api_worlds.go":
			removePersistenceDuplicates(f)
		}
	}

	for name, f := range files {
		var buf bytes.Buffer
		if err := format.Node(&buf, fset, f); err != nil {
			return fmt.Errorf("format %s: %w", name, err)
		}
		content := buf.Bytes()

		if strings.HasPrefix(name, "model_") {
			content = dropDisallowUnknownFields(content)
		}

		outPath := filepath.Join(dstDir, name)
		if err := os.WriteFile(outPath, content, 0o644); err != nil {
			return fmt.Errorf("write %s: %w", outPath, err)
		}
	}

	fmt.Printf("migrated %d files to %s\n", len(files), dstDir)
	return nil
}

// buildEnumTable scans every file for enum definitions of the shape
//
//	type T string
//	const ( NAME T = "..." ... )
//
// and returns map[T]map[NAME]struct{} — the set of bare const names per
// enum type. We need this both to know which idents to rename and to scope
// the rename by the surrounding declared type (so a "PUBLIC" used as
// ReleaseStatus default doesn't get rewritten using the InstanceType table).
func buildEnumTable(files map[string]*ast.File) map[string]map[string]struct{} {
	stringTypes := map[string]struct{}{}
	for _, f := range files {
		for _, decl := range f.Decls {
			gd, ok := decl.(*ast.GenDecl)
			if !ok || gd.Tok != token.TYPE {
				continue
			}
			for _, spec := range gd.Specs {
				ts, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}
				if ident, ok := ts.Type.(*ast.Ident); ok && ident.Name == "string" {
					stringTypes[ts.Name.Name] = struct{}{}
				}
			}
		}
	}

	out := map[string]map[string]struct{}{}
	for _, f := range files {
		for _, decl := range f.Decls {
			gd, ok := decl.(*ast.GenDecl)
			if !ok || gd.Tok != token.CONST {
				continue
			}
			for _, spec := range gd.Specs {
				vs, ok := spec.(*ast.ValueSpec)
				if !ok || vs.Type == nil {
					continue
				}
				tIdent, ok := vs.Type.(*ast.Ident)
				if !ok {
					continue
				}
				if _, isEnum := stringTypes[tIdent.Name]; !isEnum {
					continue
				}
				if out[tIdent.Name] == nil {
					out[tIdent.Name] = map[string]struct{}{}
				}
				for _, n := range vs.Names {
					out[tIdent.Name][n.Name] = struct{}{}
				}
			}
		}
	}
	return out
}

// prefixEnums rewrites both the enum const declarations and any
// "var x T = NAME" default-value references in a file.
func prefixEnums(f *ast.File, enums map[string]map[string]struct{}) {
	prefix := func(typeName, base string) string {
		want := typeName + "_" + base
		if strings.HasPrefix(base, typeName+"_") {
			return base
		}
		return want
	}

	for _, decl := range f.Decls {
		gd, ok := decl.(*ast.GenDecl)
		if !ok || gd.Tok != token.CONST {
			continue
		}
		for _, spec := range gd.Specs {
			vs, ok := spec.(*ast.ValueSpec)
			if !ok || vs.Type == nil {
				continue
			}
			tIdent, ok := vs.Type.(*ast.Ident)
			if !ok {
				continue
			}
			if _, isEnum := enums[tIdent.Name]; !isEnum {
				continue
			}
			for _, n := range vs.Names {
				n.Name = prefix(tIdent.Name, n.Name)
			}
		}
	}

	ast.Inspect(f, func(n ast.Node) bool {
		vs, ok := n.(*ast.ValueSpec)
		if !ok || vs.Type == nil {
			return true
		}
		tIdent, ok := vs.Type.(*ast.Ident)
		if !ok {
			return true
		}
		members, ok := enums[tIdent.Name]
		if !ok {
			return true
		}
		for _, val := range vs.Values {
			ident, ok := val.(*ast.Ident)
			if !ok {
				continue
			}
			if _, isMember := members[ident.Name]; isMember {
				ident.Name = prefix(tIdent.Name, ident.Name)
			}
		}
		return true
	})
}

// relaxEnumUnmarshal rewrites UnmarshalJSON methods on enum types so unknown
// values are accepted as-is instead of erroring on a stale allowlist. The
// validating IsValid() and New<T>FromValue helpers are left untouched, so any
// caller that explicitly wants to check enum membership still can.
func relaxEnumUnmarshal(f *ast.File, enums map[string]map[string]struct{}) {
	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Name.Name != "UnmarshalJSON" {
			continue
		}
		recv := receiverTypeName(fn)
		if _, isEnum := enums[recv]; !isEnum {
			continue
		}
		fn.Body = relaxedEnumBody(recv)
	}
}

// relaxedEnumBody returns the AST of:
//
//	{
//	    var value string
//	    err := json.Unmarshal(src, &value)
//	    if err != nil {
//	        return err
//	    }
//	    *v = T(value)
//	    return nil
//	}
func relaxedEnumBody(typeName string) *ast.BlockStmt {
	return &ast.BlockStmt{List: []ast.Stmt{
		&ast.DeclStmt{Decl: &ast.GenDecl{
			Tok: token.VAR,
			Specs: []ast.Spec{&ast.ValueSpec{
				Names: []*ast.Ident{{Name: "value"}},
				Type:  &ast.Ident{Name: "string"},
			}},
		}},
		&ast.AssignStmt{
			Lhs: []ast.Expr{&ast.Ident{Name: "err"}},
			Tok: token.DEFINE,
			Rhs: []ast.Expr{&ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   &ast.Ident{Name: "json"},
					Sel: &ast.Ident{Name: "Unmarshal"},
				},
				Args: []ast.Expr{
					&ast.Ident{Name: "src"},
					&ast.UnaryExpr{Op: token.AND, X: &ast.Ident{Name: "value"}},
				},
			}},
		},
		&ast.IfStmt{
			Cond: &ast.BinaryExpr{
				X:  &ast.Ident{Name: "err"},
				Op: token.NEQ,
				Y:  &ast.Ident{Name: "nil"},
			},
			Body: &ast.BlockStmt{List: []ast.Stmt{
				&ast.ReturnStmt{Results: []ast.Expr{&ast.Ident{Name: "err"}}},
			}},
		},
		&ast.AssignStmt{
			Lhs: []ast.Expr{&ast.StarExpr{X: &ast.Ident{Name: "v"}}},
			Tok: token.ASSIGN,
			Rhs: []ast.Expr{&ast.CallExpr{
				Fun:  &ast.Ident{Name: typeName},
				Args: []ast.Expr{&ast.Ident{Name: "value"}},
			}},
		},
		&ast.ReturnStmt{Results: []ast.Expr{&ast.Ident{Name: "nil"}}},
	}}
}

// rewritePathEscapeCalls replaces every "url.PathEscape(x)" call expression
// with "pathSegmentEscape(x)". The generator's Go template wraps every path
// parameter substitution in url.PathEscape, which percent-encodes parens
// even though they're legal RFC 3986 sub-delims — and VRChat instance IDs
// contain literal parens that Cloudflare's WAF rejects when escaped. The
// replacement helper lives in the hand-written path_escape.go at the repo
// root. The "net/url" import stays live because url.Values{} is still used
// for query parameters in the same files.
func rewritePathEscapeCalls(f *ast.File) {
	ast.Inspect(f, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}
		sel, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			return true
		}
		x, ok := sel.X.(*ast.Ident)
		if !ok {
			return true
		}
		if x.Name == "url" && sel.Sel.Name == "PathEscape" {
			call.Fun = &ast.Ident{Name: "pathSegmentEscape"}
		}
		return true
	})
}

// addImport inserts `path` into the file's first import block, keeping
// the block sorted by import path.
func addImport(f *ast.File, path string) {
	quoted := `"` + path + `"`
	for _, decl := range f.Decls {
		gd, ok := decl.(*ast.GenDecl)
		if !ok || gd.Tok != token.IMPORT {
			continue
		}
		for _, spec := range gd.Specs {
			if is, ok := spec.(*ast.ImportSpec); ok && is.Path != nil && is.Path.Value == quoted {
				return
			}
		}
		gd.Specs = append(gd.Specs, &ast.ImportSpec{
			Path: &ast.BasicLit{Kind: token.STRING, Value: quoted},
		})
		// Force the block to be parenthesized so the printer formats it as a list.
		if gd.Lparen == token.NoPos {
			gd.Lparen = gd.TokPos + 1
			gd.Rparen = gd.TokPos + 2
		}
		sort.SliceStable(gd.Specs, func(i, j int) bool {
			return gd.Specs[i].(*ast.ImportSpec).Path.Value < gd.Specs[j].(*ast.ImportSpec).Path.Value
		})
		return
	}
}

// removePersistenceDuplicates strips the three Api*UserPersistence* request
// types and their associated methods from api_worlds.go. The same operations
// are emitted into api_users.go (because the spec tags them under both
// "users" and "worlds") and the duplicates would cause redeclaration errors.
func removePersistenceDuplicates(f *ast.File) {
	dupTypes := map[string]bool{
		"ApiCheckUserPersistenceExistsRequest":   true,
		"ApiDeleteAllUserPersistenceDataRequest": true,
		"ApiDeleteUserPersistenceRequest":        true,
	}
	dupServiceMethods := map[string]bool{
		"CheckUserPersistenceExists":          true,
		"CheckUserPersistenceExistsExecute":   true,
		"DeleteAllUserPersistenceData":        true,
		"DeleteAllUserPersistenceDataExecute": true,
		"DeleteUserPersistence":               true,
		"DeleteUserPersistenceExecute":        true,
	}

	type span struct{ start, end token.Pos }
	var dropped []span
	kept := f.Decls[:0]
	for _, d := range f.Decls {
		drop := false
		switch dd := d.(type) {
		case *ast.GenDecl:
			if dd.Tok == token.TYPE {
				for _, spec := range dd.Specs {
					if ts, ok := spec.(*ast.TypeSpec); ok && dupTypes[ts.Name.Name] {
						drop = true
						break
					}
				}
			}
		case *ast.FuncDecl:
			recv := receiverTypeName(dd)
			if dupTypes[recv] {
				drop = true
			} else if recv == "WorldsAPIService" && dupServiceMethods[dd.Name.Name] {
				drop = true
			}
		}
		if drop {
			start, end := d.Pos(), d.End()
			// Doc comments precede d.Pos(); include them in the dropped span.
			switch dd := d.(type) {
			case *ast.GenDecl:
				if dd.Doc != nil {
					start = dd.Doc.Pos()
				}
			case *ast.FuncDecl:
				if dd.Doc != nil {
					start = dd.Doc.Pos()
				}
			}
			dropped = append(dropped, span{start, end})
			continue
		}
		kept = append(kept, d)
	}
	f.Decls = kept

	if len(dropped) > 0 {
		filtered := f.Comments[:0]
		for _, cg := range f.Comments {
			cgStart, cgEnd := cg.Pos(), cg.End()
			inside := false
			for _, s := range dropped {
				if cgStart >= s.start && cgEnd <= s.end {
					inside = true
					break
				}
			}
			if !inside {
				filtered = append(filtered, cg)
			}
		}
		f.Comments = filtered
	}
}

func receiverTypeName(fn *ast.FuncDecl) string {
	if fn.Recv == nil || len(fn.Recv.List) == 0 {
		return ""
	}
	switch t := fn.Recv.List[0].Type.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		if id, ok := t.X.(*ast.Ident); ok {
			return id.Name
		}
	}
	return ""
}

// dropDisallowUnknownFields removes any line that is exactly
// "decoder.DisallowUnknownFields()" (with surrounding whitespace) from a
// model file's source. Done as a string-level pass after the AST is printed
// because it's a single-line removal and we want to leave everything else
// the printer produced byte-identical.
func dropDisallowUnknownFields(content []byte) []byte {
	lines := bytes.Split(content, []byte("\n"))
	out := lines[:0]
	for _, line := range lines {
		if bytes.Equal(bytes.TrimSpace(line), []byte("decoder.DisallowUnknownFields()")) {
			continue
		}
		out = append(out, line)
	}
	return bytes.Join(out, []byte("\n"))
}
