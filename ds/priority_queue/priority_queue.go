package priority_queue

// 定义优先队列的接口, 实现了这个接口的对象就可以使用这个优先队列
type PQNode interface {
	Less() bool
}

type PriorityQueue struct {

}

func (pq *PriorityQueue) Push(node PQNode) {

}

func (pq *PriorityQueue) Pop(node PQNode) PQNode {
	return nil
}

func (pq PriorityQueue) Top(node PQNode) PQNode {
	return nil
}