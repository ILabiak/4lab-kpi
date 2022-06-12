package engine

import (
	"fmt"
	"strings"
)


func Parse(commandLine string) {
	parts := strings.Fields(commandLine)

	if len(parts) < 2 {
		fmt.Println("Err")
		return
	}

	switch parts[0] {
	case "print":
		fmt.Println(parts[1])
		return
	}

	fmt.Println("SYNTAX ERROR: Unknown instruction")
}
