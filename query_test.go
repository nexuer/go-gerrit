package gerrit

import (
	"fmt"
	"testing"
)

func TestQueryString(t *testing.T) {
	tests := []struct {
		input fmt.Stringer
		want  string
	}{
		{
			input: F("", ""),
			want:  "",
		},
		{
			input: F("status", ""),
			want:  `status:""`,
		},
		{
			input: F("status", "open"),
			want:  "status:open",
		},

		{
			input: F("message", "open and close"),
			want:  `message:"open and close"`,
		},
		{
			input: Or(),
			want:  "",
		},
		{
			input: Or(F("status", "open")),
			want:  "status:open",
		},
		{
			input: Or(
				F("status", "open"),
				F("status", "merged"),
			),
			want: "(status:open OR status:merged)",
		},

		{
			input: And(
				F("status", "open"),
			),
			want: "status:open",
		},
		{
			input: And(
				F("status", "open"),
				F("status", "merged"),
			),
			want: "(status:open AND status:merged)",
		},

		{
			input: Not(F("status", "open")),
			want:  "-status:open",
		},

		{
			input: Or(
				F("status", "open"),
				And(
					F("status", "merged"),
					F("status", "abandoned"),
				),
				Not(
					And(
						F("has", "draft"),
						F("project", "Foo"),
					),
				),
			),

			want: "(status:open OR (status:merged AND status:abandoned) OR -(has:draft AND project:Foo))",
		},
	}

	for _, tt := range tests {
		got := tt.input.String()
		if got != tt.want {
			t.Errorf("\ngot:\n%v\nwant:\n%v", got, tt.want)
		}
	}
}
