package graph

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseNaively(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "vuls",
			arg:  vulsInput,
			want: fmt.Sprintf(vulsOutput, "```", "```"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseNaively(tt.arg)
			if err != nil {
				t.Errorf("err: %s", err.Error())
			} else if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Hogefunc differs: (-got +want)\n%s", diff)
			}
		})
	}
}

const vulsInput = `github.com/future-architect/vuls cloud.google.com/go@v0.100.2
github.com/future-architect/vuls cloud.google.com/go/compute@v1.5.0
github.com/future-architect/vuls cloud.google.com/go/iam@v0.3.0
github.com/future-architect/vuls cloud.google.com/go/storage@v1.14.0
cloud.google.com/go@v0.100.2 cloud.google.com/go/compute@v0.1.0
cloud.google.com/go@v0.100.2 github.com/golang/protobuf@v1.5.2
cloud.google.com/go@v0.100.2 github.com/google/go-cmp@v0.5.6
`

// HACK: tell me how to escape backquote
const vulsOutput = `# GO MOD GRAPH

%smermaid
graph TD
    1("github.com/future-architect/vuls") --> 2("cloud.google.com/go@v0.100.2")
    1 --> 3("cloud.google.com/go/compute@v1.5.0")
    1 --> 4("cloud.google.com/go/iam@v0.3.0")
    1 --> 5("cloud.google.com/go/storage@v1.14.0")
    2 --> 6("cloud.google.com/go/compute@v0.1.0")
    2 --> 7("github.com/golang/protobuf@v1.5.2")
    2 --> 8("github.com/google/go-cmp@v0.5.6")
%s
`
