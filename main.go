package main

import (
	"bufio"
	"os"

	"github.com/ILabiak/4lab-kpi/engine"
)

func main() {
	eventLoop := new(engine.EventLoop)
	eventLoop.Start()

	if input, err := os.Open("test.txt"); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			commandLine := scanner.Text()
			cmd := engine.Parse(commandLine)
			eventLoop.Post(cmd)
		}
	}


	eventLoop.AwaitFinish()
}
