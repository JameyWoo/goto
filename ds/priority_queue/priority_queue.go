package priority_queue

// 定义优先队列的接口, 实现了这个接口的对象就可以使用这个优先队列
type PQInterface interface {
	Less(i, j int) bool
}

type PriorityQueue struct {
	PQInterface
	nodes []interface{}
}

func (pq *PriorityQueue) Push(node interface{}) {

}

func (pq *PriorityQueue) Pop(node interface{}) interface{} {
	return nil
}

func (pq PriorityQueue) Top() interface{} {
	return nil
}