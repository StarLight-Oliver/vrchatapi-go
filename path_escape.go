package vrchatapi

import (
	"fmt"
	"strings"
)

// pathSegmentEscape escapes a string for use as a URL path segment per
// RFC 3986 pchar (unreserved + sub-delims + ":" + "@"). It is intentionally
// less aggressive than net/url's PathEscape, which over-escapes sub-delims
// like "(" and ")" that some upstream WAFs (notably Cloudflare) then reject
// as malformed. VRChat instance IDs depend on literal parens.
func pathSegmentEscape(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if isPchar(c) {
			b.WriteByte(c)
		} else {
			fmt.Fprintf(&b, "%%%02X", c)
		}
	}
	return b.String()
}

func isPchar(c byte) bool {
	switch {
	case c >= 'A' && c <= 'Z', c >= 'a' && c <= 'z', c >= '0' && c <= '9':
		return true
	}
	switch c {
	case '-', '.', '_', '~',
		'!', '$', '&', '\'', '(', ')', '*', '+', ',', ';', '=',
		':', '@':
		return true
	}
	return false
}
