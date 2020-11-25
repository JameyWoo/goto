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
    std::queue<int> Q;  // 存放chan value的队列
    // 条件变量 + 锁, 控制线程的阻塞与唤醒
    std::mutex mtx;
    std::mutex io_mtx;
    std::condition_variable cv;

   public:
    Chan() {
        len = 0;
        cap = 0;
        wait_count = 0;
        // lock = std::unique_lock<std::mutex>(mtx);
    }
    Chan(int c) {
        len = 0;
        cap = c;
        wait_count = 0;
    }

   public:
    int out();
    void in(int in_val);
};

#endif  // __CHANNEL_HPP__
