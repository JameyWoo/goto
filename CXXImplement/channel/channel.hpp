#ifndef MUDUO_BASE_BOUNDEDBLOCKINGQUEUE_H
#define MUDUO_BASE_BOUNDEDBLOCKINGQUEUE_H


#include <condition_variable>
#include <mutex>
#include <queue>


/*
 * 基于支持无缓冲模式的生产者消费者队列实现的channel
 * */

template<typename T>
class Chan {
public:
    Chan();
    Chan(int max_size);

public:
    void put(const T &x);  // chan T <-   // T 为 typename
    void put(T &&x);
    T get();  // <- chan T

private:
    // 锁
    std::mutex mutex_{};
    // 条件变量
    std::condition_variable notEmpty;  // 在 get 中使用
    std::condition_variable notFull;  // 在 put 中使用

private:
    size_t capacity;  // 容量
    std::queue<T> Queue;

    // 当缓冲为0时存储变量到队列Q中
    std::queue<T> Q;
    int cnt_get;  // 在等待的get的数量
    int cnt_put;  // 在等待的put的数量
};


template<typename T>
Chan<T>::Chan() : Chan(0) {}

template<typename T>
Chan<T>::Chan(int max_size) {
    capacity = max_size;

    cnt_get = 0;
    cnt_put = 0;
}

template<typename T>
void Chan<T>::put(const T &x) {
    // 在unique_lock对象的声明周期内，它所管理的锁对象会一直保持上锁状态
    // unique_lock 具有 lock_guard 的所有功能，而且更为灵活
    std::unique_lock<std::mutex> lck(mutex_);
    // 当缓冲区大小为 0时
    if (capacity == 0) {
        Q.push(x);
        cnt_put += 1;
        if (cnt_get > 0) {
            notEmpty.notify_one();
            cnt_get -= 1;
            return;
        }
        while (cnt_get == 0) {
            notFull.wait(lck);
        }
        // 被唤醒后, 没有阻塞的get了
        cnt_get -= 1;
        return;
    }

    while (Queue.size() == capacity) {
        // 传 引用
        notFull.wait(std::ref(lck));
    }

    Queue.push(x);

    // 告诉消费者可以消费了
    notEmpty.notify_one();
}

template<typename T>
void Chan<T>::put(T &&x) {
    put(x);
}

template<typename T>
T Chan<T>::get() {
    std::unique_lock<std::mutex> lck(mutex_);
    // 当缓冲区大小为 0时
    if (capacity == 0) {
        cnt_get += 1;
        if (cnt_put > 0) {
            notFull.notify_one();
            cnt_put -= 1;
            T tmp = Q.front();
            Q.pop();
            return tmp;
        }
        while (cnt_put == 0) {
            notEmpty.wait(std::ref(lck));
        }
        // 被唤醒后, 没有阻塞的put了
        cnt_put -= 1;
        T tmp = Q.front();
        Q.pop();
        return tmp;
    }

    while (Queue.size() == 0) {
        notEmpty.wait(std::ref(lck));
    }

    T val = Queue.front();
    Queue.pop();
    notFull.notify_one();

    return val;
}

#endif