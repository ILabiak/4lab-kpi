package engine

import "sync"

type Queue struct {
	me             sync.Mutex
	commands       []Command
	emptyFlag      bool
	notEmptySignal chan struct{}
}

func (cq *Queue) pull() Command {
	cq.me.Lock()
	defer cq.me.Unlock()

	if cq.empty() {
		cq.emptyFlag = true
		cq.me.Unlock()

		<-cq.notEmptySignal
		cq.me.Lock()
	}

	res := cq.commands[0]
	cq.commands[0] = nil
	cq.commands = cq.commands[1:]
	return res
}

func (cq *Queue) push(c Command) {
	cq.me.Lock()
	defer cq.me.Unlock()
	cq.commands = append(cq.commands, c)

	if cq.emptyFlag {
		cq.emptyFlag = false
		cq.notEmptySignal <- struct{}{}
	}
}

func (cq *Queue) empty() bool {
	//cq.me.Lock()
	//defer cq.me.Unlock()
	return len(cq.commands) == 0
}
