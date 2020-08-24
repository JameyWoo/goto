/*
一个简单的 int 类型 queue 实现
 */

package queue

type queue []int

// 获取size
func (q *queue) Size() int {
	return len(*q)
}

// 添加元素
func (q *queue) Push(num ...int) {
	*q = append(*q, num...)
}

// 弹出队首
func (q *queue) Pop() {
	*q = (*q)[1:]
}

// 是否为空
func (q *queue) Empty() bool {
	return len(*q) == 0
}

// 获得队首
func (q *queue) Front() int {
	return (*q)[0]
}