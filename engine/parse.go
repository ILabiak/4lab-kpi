package engine

import (
	"strings"
)


func Parse(commandLine string) Command {
	parts := strings.Fields(commandLine)

	if len(parts) < 2 {
		return &printCommand{"Error: not enough arguments"}
	}

	switch parts[0] {
	case "print":
		return &printCommand{parts[1]}
	case "delete":
		if len(parts) < 3 {
			return &printCommand{"Error: not enough arguments for delete function"}
		}
		str := parts[1]
		symbol := parts[2]
		return &deleteCommand{str: str, symbol: symbol}
	}

	return &printCommand{"err"}
}
