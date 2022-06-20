package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoop(t *testing.T) {
	printCmd1 := &printCommand{
		arg: "smth",
	}
	printCmd2 := &printCommand{

		arg: "test2",
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
}

func TestParser(t *testing.T) {
	var commandStr string

	commandStr = "print hello"
	cmd := Parse(commandStr)
	assert.Equal(t, &printCommand{"hello"}, cmd)

	commandStr = "print"
	cmd = Parse(commandStr)
	assert.Equal(t, &printCommand{"Error: not enough arguments"}, cmd)

	commandStr = "wrongcommand hello"
	cmd = Parse(commandStr)
	assert.Equal(t, &printCommand{"err"}, cmd)

	commandStr = "delete hello l"
	cmd = Parse(commandStr)
	assert.Equal(t, &deleteCommand{str: "hello", symbol: "l"}, cmd)

	commandStr = "delete hello"
	cmd = Parse(commandStr)
	assert.Equal(t, &printCommand{"Error: not enough arguments for delete function"}, cmd)

	commandStr = "delete"
	cmd = Parse(commandStr)
	assert.Equal(t, &printCommand{"Error: not enough arguments"}, cmd)
}
