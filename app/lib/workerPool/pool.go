package workerPool

type Pool struct {
	ExternalChain chan *Task
	taskChain     chan *Task //消息管道
	workNum       int        //协程数量
}

func (p *Pool) Run() {

	for i := 0; i < p.workNum; i++ {
		go p.work()
	}
	p.sendTask()

}

func (p *Pool) sendTask() {
	for task := range p.ExternalChain {
		p.taskChain <- task
	}

}

func (p *Pool) work() {
	for task := range p.taskChain {
		task.execute()
	}
}

func NewPool(workNum int) *Pool {
	p := Pool{
		ExternalChain: make(chan *Task),
		taskChain:     make(chan *Task),
		workNum:       workNum,
	}

	return &p
}
