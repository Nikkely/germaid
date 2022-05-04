package graph

import (
	"fmt"
	"strings"
)

const title = "# GO MOD GRAPH"
const header = "```mermaid\ngraph TD"
const footer = "```"

// Parse outputs mermaid from "go mod graph"
func ParseNaively(input string) (output string, err error) {
	output = fmt.Sprintf("%s\n\n%s\n", title, header)

	moduleMap := map[string]int{}
	cnt := 0
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		l := strings.TrimSpace(line)
		if l == "" {
			break // EOF
		}

		mods := strings.Split(l, " ")
		if len(mods) != 2 {
			err = fmt.Errorf("unexpected input line:%d\n", i+1)
			return
		}
		var a, b string
		if k, ok := moduleMap[mods[0]]; !ok {
			cnt++
			moduleMap[mods[0]] = cnt
			a = fmt.Sprintf(`%d("%s")`, cnt, mods[0])
		} else {
			a = fmt.Sprintf("%d", k)
		}
		if k, ok := moduleMap[mods[1]]; !ok {
			cnt++
			moduleMap[mods[1]] = cnt
			b = fmt.Sprintf(`%d("%s")`, cnt, mods[1])
		} else {
			b = fmt.Sprintf("%d", k)
		}
		output += fmt.Sprintf("    %s --> %s\n", a, b)
	}

	output += fmt.Sprintf("%s\n", footer)
	return
}
