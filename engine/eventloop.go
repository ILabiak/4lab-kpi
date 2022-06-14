package engine

type EventLoop struct {
	q *Queue

	stop       bool
	stopSignal chan struct{}
}

func (l *EventLoop) Start() {
	l.q = &Queue{
		notEmptySignal: make(chan struct{}),
	}
	l.stopSignal = make(chan struct{})
	go func() {
		for !l.stop || !l.q.empty() {
			cmd := l.q.pull()
			cmd.Execute(l)
		}
		l.stopSignal <- struct{}{}
	}()
}

func (l *EventLoop) Post(cmd Command) {
	l.q.push(cmd)
}

func (l *EventLoop) AwaitFinish() {
	l.Post(&stopCommand{})
	//l.stop = true
	<-l.stopSignal
}
