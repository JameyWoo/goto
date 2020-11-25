#if !defined(__CHANNEL_HPP__)
#define __CHANNEL_HPP__

#include <iostream>
#include <queue>
#include <mutex>
#include <condition_variable>

class Chan {
   private:
    int len;
    int cap;
    int wait_count;
    int in_wait_count;  // 因为 in操作 而导致阻塞的值的数量
    int out_wait_count;
    std::queue<int> Q;  // 存放chan value的队列
    // 条件变量 + 锁, 控制线程的阻塞与唤醒
    std::mutex mtx;
    // std::mutex io_mtx;
    // 关于in和out的阻塞线程, 需要用两个条件变量表示. 因为只能in唤醒out, out唤醒in
    std::condition_variable in_cv;
    std::condition_variable out_cv;
    // 使用一个vector保存阻塞时进入到chan的值, 作为queue的后续
    std::queue<int> in_next;

   public:
    Chan() {
        len = 0;
        cap = 0;
        wait_count = 0;
        in_wait_count = 0;
        // lock = std::unique_lock<std::mutex>(mtx);
    }
    Chan(int c) {
        len = 0;
        cap = c;
        wait_count = 0;
        in_wait_count = 0;
    }

   public:
    int out();
    void in(int in_val);
};

#endif  // __CHANNEL_HPP__
