## v0.1

先实现 int 的chan, 之后改成模板

参考:
1. C++11 中没有信号量，如何进行同步或互斥？ - 知乎
https://www.zhihu.com/question/31555101/answer/117537944
2. C++11 线程的使用: https://www.runoob.com/w3cnote/cpp-std-thread.html

该版本实现了少数in-out操作的正确性, 但一旦多了in或者out, 则会导致死锁


## v0.2

该版本如果 in 在前 out 在后, 则可以支持上千线程的并发正确

但是如果 out 在前 in 在后, 则即使是 10 个线程, 也会发生死锁 (需要修复的bug)

错误日志:
```
2020-12-03 13:49:04.700 (   0.001s) [        C67C0700]                   channel.cpp:34    INFO| out: I notify as in!
2020-12-03 13:49:04.700 (   0.001s) [        C5FB0700]                   channel.cpp:34    INFO| out: I notify as in!
2020-12-03 13:49:04.700 (   0.001s) [        C5FB0700]                 chan_test.cpp:18    WARN| out: 1
2020-12-03 13:49:04.699 (   0.001s) [        CFC70700]                   channel.cpp:85    INFO| in: I block!
2020-12-03 13:49:04.700 (   0.002s) [        B6FD0700]                   channel.cpp:34    INFO| out: I notify as in!
2020-12-03 13:49:04.701 (   0.002s) [        B5FB0700]                   channel.cpp:34    INFO| out: I notify as in!
2020-12-03 13:49:04.700 (   0.002s) [        C4F90700]                   channel.cpp:34    INFO| out: I notify as in!
2020-12-03 13:49:04.701 (   0.002s) [        B67C0700]                   channel.cpp:34    INFO| out: I notify as in!
2020-12-03 13:49:04.701 (   0.002s) [        B67C0700]                 chan_test.cpp:18    WARN| out: 7
2020-12-03 13:49:04.700 (   0.002s) [        C67C0700]                 chan_test.cpp:18    WARN| out: 0
2020-12-03 13:49:04.700 (   0.002s) [        B77E0700]                   channel.cpp:34    INFO| out: I notify as in!
2020-12-03 13:49:04.700 (   0.002s) [        B7FF0700]                   channel.cpp:34    INFO| out: I notify as in!
2020-12-03 13:49:04.700 (   0.002s) [        C57A0700]                   channel.cpp:34    INFO| out: I notify as in!
```
此处可看到, out提前notify了in, 导致后面有的in无法苏醒

描述过程:
```
对于 chan <- 操作(in):
    先判断len和cap:
        如果cap>len:
            判断等待队列waitQ是否有val (加锁)
                如果有:
                    用while将所有val添加到Q, 然后判断cap>len
                        如果是:
                            添加
                        否则:
                            添加到waitQ中
                            直接阻塞(或标记在后面阻塞)
                否则:
                    直接添加, 不阻塞
            // 此处阻塞, 避免死锁
            // 如果不阻塞, 判断是否有线程在等待, 唤醒
    否则
        添加到waitQ, 阻塞
        
对于 <- chan 操作(out):
    if len > 0:
        取出一个
    else:
        if waitQ.size > 0:  // 可能管道容量不够
            取出一个
            唤醒一个
        else:
            阻塞
            out_wait++
```
