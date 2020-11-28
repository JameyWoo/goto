package main

import "fmt"
import "time"

//函数
func go_worker(name string) {
	for i := 0; i < 5; i++ {
		fmt.Println("我是一个go的协程，我的名字是 ", name)
		time.Sleep(1 * time.Second)
	}

	fmt.Println(name, "执行完毕")
}

//main,  go协程1，  go协程2
func main() {
	//协程 保证程序的并行

	//开辟一个go的协程 去执行 go_worker("xiaohei")
	go go_worker("小黑")
	//开辟一个go的协程 去执行 go_worker("xiaobai")
	go go_worker("小白")

	for i := 0; i < 10; i++ {
		fmt.Println("我是main...")
		time.Sleep(1 * time.Second)
	}
}
