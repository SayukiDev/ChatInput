package tasks

type Handle func(c chan struct{})

type Runner struct {
	closeC chan struct{}
	handle Handle
}

func NewRunner(handle Handle) *Runner {
	return &Runner{
		closeC: make(chan struct{}),
		handle: handle,
	}
}

func (r *Runner) Start() {
	go r.handle(r.closeC)
}

func (r *Runner) Close() {
	close(r.closeC)
}
