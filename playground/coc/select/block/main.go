package main

// 观察阻塞的等待队列
func main() {
	ch := make(chan int)

	go func() {<-ch}()
	go func () {<-ch}()
	//runtime.Gosched()

	//fmt.Println("hello")
	//
	//fmt.Println("world")

	_ = 12
}
