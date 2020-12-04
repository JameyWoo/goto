//
// Created by 26220 on 2020/12/4.
//

#ifndef CHAN_TEST_WAITGROUP_HPP
#define CHAN_TEST_WAITGROUP_HPP

#include <iostream>
#include <mutex>
#include <atomic>
#include <condition_variable>

#include "../log/loguru.hpp"


class WaitGroup {
public:
    WaitGroup();
    void Add(int delta);
    void Done();  // Add(-1);
    void Wait();  // 等待所有线程执行结束

private:
    std::atomic<int> count;
    std::atomic<int> wait;
    // 互斥锁
    std::mutex mtx{};
    // 条件变量
    std::condition_variable cv;
};

WaitGroup::WaitGroup() {
    count = 0;
    wait = 0;
}

void WaitGroup::Add(int delta = 1) {
    std::atomic_fetch_add(&count, delta);
    if (std::atomic_load(&count) == 0) {
        cv.notify_all();
    }
}

void WaitGroup::Done() {
    Add(-1);
}

void WaitGroup::Wait() {
    std::unique_lock<std::mutex> lck(mtx);
    // 加锁的临界区, 不需要使用 atomic 了
    wait += 1;
    cv.wait(lck, [&]{return count <= 0;});
}

#endif //CHAN_TEST_WAITGROUP_HPP
