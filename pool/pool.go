package pool

type Task func() error

type Pool struct {
	taskNum, jobCnt int
	SetTask, getTask chan Task
}

func (p *Pool) Run() {
	for i := 0; i < p.taskNum; i++ {
		go p.work()
	}
	for t := range p.SetTask {
		p.getTask <- t
	}
}

// 执行工作
func (p *Pool) work() {
	for t := range p.getTask {
		t()
		p.jobCnt += 1
	}
}

// 添加任务
func (p *Pool) Add(task Task) {
	p.SetTask <- task
}

// 创建一个 协程池
func NewPool(cap int) *Pool {
	return &Pool{
		taskNum: cap,
		jobCnt: 0,
		SetTask: make(chan Task),
		getTask: make(chan Task),
	}
}