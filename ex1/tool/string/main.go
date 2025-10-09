package main

import (
	"fmt"
	"strings"

	"github.com/atotto/clipboard"
)

func main() {
	all, err := clipboard.ReadAll()
	if err != nil {
		panic(err)
	}

	sb := strings.Builder{}
	list := strings.Split(all, "\n")
	for i, s := range list {
		sb.WriteString(fmt.Sprintf("### %s\n?\ndf", s))
		if i != len(list)-1 {
			sb.WriteString("\n\n")
		}
	}

	fmt.Println("kacker")

	clipboard.WriteAll(sb.String()[:])
}
