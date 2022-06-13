package engine

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoop(t *testing.T) {
	printCmd1 := &printCommand{
		arg: "test 1",
	}
	printCmd2 := &printCommand{
		arg: "test 2",
	}
	deleteCmd := &deleteCommand{
		str:    "hello",
		symbol: "l",
	}

	eventLoop := new(EventLoop)
	eventLoop.Start()
	assert.Equal(t, false, eventLoop.stop)
	assert.Equal(t, 0, len(eventLoop.q.commands))

	eventLoop.Post(deleteCmd)
	eventLoop.Post(printCmd1)
	eventLoop.Post(printCmd2)

	assert.Equal(t, 3, len(eventLoop.q.commands))
	eventLoop.AwaitFinish()
	assert.Equal(t, true, eventLoop.stop)
	assert.Equal(t, 0, len(eventLoop.q.commands))

	var outputLines []string
	readFile, _ := os.Open("results.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		outputLines = append(outputLines, fileScanner.Text())
	}

	assert.Equal(t, printCmd1.arg, outputLines[0])
	assert.Equal(t, printCmd2.arg, outputLines[1])
	assert.Equal(t, strings.ReplaceAll(deleteCmd.str, deleteCmd.symbol, ""), outputLines[2])
}
