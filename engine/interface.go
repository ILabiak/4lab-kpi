package engine

import (
	"fmt"
	"strings"
)

// Command represents actions that can be performed
// in a single event loop iteration.
type Command interface {
	Execute(handler Handler)
}

// Handler allows to send commands to an event loop
// it’s associated with.
type Handler interface {
	Post(cmd Command) error
}

type printCommand struct {
	arg string
}

func (p *printCommand) Execute(loop Handler) {
	fmt.Println(p.arg)
}

type deleteCommand struct {
	str    string
	symbol string
}

func (del *deleteCommand) Execute(loop Handler) {
	res := strings.ReplaceAll(del.str, del.symbol, "")
	loop.Post(&printCommand{arg: res})
}

type stopCommand struct{}

func (sc stopCommand) Execute(h Handler) {
	h.(*EventLoop).stop = true
}
