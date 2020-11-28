/*
一个简单的 int 类型 Queue 实现
 */

package queue

type Queue []int


// 获取size
func (q *Queue) Size() int {
	return len(*q)
}

// 添加元素
func (q *Queue) Push(num ...int) {
	*q = append(*q, num...)
}

// 弹出队首
func (q *Queue) Pop() {
	*q = (*q)[1:]
}

// 是否为空
func (q *Queue) Empty() bool {
	return len(*q) == 0
}

// 获得队首
func (q *Queue) Front() int {
	return (*q)[0]
}