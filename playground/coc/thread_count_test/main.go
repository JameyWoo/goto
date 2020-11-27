package main

import "time"

/*
测试线程数量
用debug工具分析

创建x个协程然后让他们阻塞
然后看线程数量的变化
 */

func main() {
	for i := 0; i < 100; i++ {
		go func() {
			time.Sleep(5 * time.Second)
		}()
	}
	time.Sleep(8 * time.Second)
}