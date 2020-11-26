#if !defined(__CHANNEL_HPP__)
#define __CHANNEL_HPP__

#include <condition_variable>
#include <iostream>
#include <mutex>
#include <queue>

class Chan {
   private:
    int len;
    int cap;
    int wait_count;
    int in_wait_count;  // 因为 in操作 而导致阻塞的值的数量
    int out_wait_count;

   private:
    // 存放chan value的队列
    std::queue<int> Q;
    // 使用一个vector保存阻塞时进入到chan的值, 作为queue的后续
    std::queue<int> in_next;
    // 条件变量 + 锁, 控制线程的阻塞与唤醒
    std::mutex mtx;
    std::mutex rwmtx;
    // 关于in和out的阻塞线程, 需要用两个条件变量表示. 因为只能in唤醒out, out唤醒in
    std::condition_variable in_cv;
    std::condition_variable out_cv;

   private:
    // 内部方法
    void init(int, int, int, int, int);
    void debug(std::string s);

   public:
    // 构造方法
    Chan();
    Chan(int c);

   public:
    // 公开方法: in(chan<-) 和 out(<-chan)
    int out();
    void in(int in_val);
};

#endif  // __CHANNEL_HPP__
