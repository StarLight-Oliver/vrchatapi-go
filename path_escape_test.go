package vrchatapi

import (
	"strings"
	"testing"
)

func TestPathSegmentEscape(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "vrchat group instance id preserves parens",
			in:   "20848~group(grp_abc)~groupAccessType(public)~region(use)",
			want: "20848~group(grp_abc)~groupAccessType(public)~region(use)",
		},
		{
			name: "reserved non-pchar bytes are percent-encoded",
			in:   "a/b c#d",
			want: "a%2Fb%20c%23d",
		},
		{
			name: "control byte is percent-encoded uppercase hex",
			in:   "x\x01y",
			want: "x%01y",
		},
		{
			name: "unreserved tilde and sub-delims pass through",
			in:   "a~b!c$d&e'f*g+h,i;j=k:l@m",
			want: "a~b!c$d&e'f*g+h,i;j=k:l@m",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := pathSegmentEscape(tc.in)
			if got != tc.want {
				t.Fatalf("pathSegmentEscape(%q)\n  got:  %q\n  want: %q", tc.in, got, tc.want)
			}
		})
	}
}

func TestPathSegmentEscape_NoOverEscapedParens(t *testing.T) {
	in := "20848~group(grp_abc)~groupAccessType(public)~region(use)"
	got := pathSegmentEscape(in)
	if strings.Contains(got, "%28") || strings.Contains(got, "%29") {
		t.Fatalf("pathSegmentEscape over-escaped parens: %q", got)
	}
	if !strings.Contains(got, "(") || !strings.Contains(got, ")") {
		t.Fatalf("pathSegmentEscape stripped parens: %q", got)
	}
}
