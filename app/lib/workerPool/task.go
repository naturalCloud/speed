package workerPool

type Task struct {
	f func() error
}

func (t *Task) execute() {
	t.f()
}

func (t *Task) SetHandel(f func() error) {
	t.f = f
}
