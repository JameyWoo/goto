#ifndef MUDUO_BASE_BOUNDEDBLOCKINGQUEUE_H
#define MUDUO_BASE_BOUNDEDBLOCKINGQUEUE_H


#include <cassert>
#include <mutex>
#include <condition_variable>
#include <vector>
#include <queue>


/*
 * 把生产者与消费者问题的buf 换成队列， 注意要线程安全
 * */

template<typename T>
class Chan {
public:
    Chan();
    Chan(int max_size);

public:
    void put(const T &x);
    void put(T &&x);
    T get();

    bool empty() const;
    bool full() const;
    size_t size() const;
    size_t capacity() const;

private:
    // 锁
    std::mutex mutex_{};
    // 条件变量
    std::condition_variable notEmpty;  // 在 get 中使用
    std::condition_variable notFull;  // 在 put 中使用

private:
    size_t capacity_;  // 容量
    size_t size_;  // 实际的大小
    std::vector<T> queue_;  // 存储数据
    int put_idx_;  // 放 下一个数据 的索引
    int get_idx_;  // 取 当前数据 的索引

    // 当缓冲为0时存储变量到队列Q中
    std::queue<T> Q;
    int cnt_get;  // 在等待的get的数量
    int cnt_put;  // 在等待的put的数量
};


template<typename T>
Chan<T>::Chan() : Chan(0) {}

template<typename T>
Chan<T>::Chan(int max_size) {
    capacity_ = max_size;
    queue_.resize(capacity_);

    size_ = 0;
    put_idx_ = 0;
    get_idx_ = 0;

    cnt_get = 0;
    cnt_put = 0;
}


template<typename T>
void Chan<T>::put(const T &x) {
    // 在unique_lock对象的声明周期内，它所管理的锁对象会一直保持上锁状态
    // unique_lock 具有 lock_guard 的所有功能，而且更为灵活
    std::unique_lock<std::mutex> lck(mutex_);
    // 当缓冲区大小为 0时
    if (capacity_ == 0) {
        Q.push(x);
        cnt_put += 1;
        if (cnt_get > 0) {
            notEmpty.notify_one();
            cnt_get -= 1;
            return;
        }
        while (cnt_get > 0) {
            notFull.wait(std::ref(lck));
        }
        // 被唤醒后, 没有阻塞的get了
        cnt_get -= 1;
        return;
    }
    while (size_ == capacity_) {
        // 传 引用
        notFull.wait(std::ref(lck));
    }
    assert(put_idx_ < capacity_ && put_idx_ >= 0);
    queue_[put_idx_] = x;
    put_idx_++;
    size_++;
    put_idx_ = put_idx_ % capacity_;
    // 告诉消费者可以消费了
    notEmpty.notify_one();
    // 此时可以解锁
    lck.unlock();
}

template<typename T>
void Chan<T>::put(T &&x) {
    put(x);
}

template<typename T>
T Chan<T>::get() {
    std::unique_lock<std::mutex> lck(mutex_);
    // 当缓冲区大小为 0时
    if (capacity_ == 0) {
        cnt_get += 1;
        if (cnt_put > 0) {
            notFull.notify_one();
            cnt_put -= 1;
            T tmp = Q.front();
            Q.pop();
            return tmp;
        }
        while (cnt_put == false) {
            notEmpty.wait(std::ref(lck));
        }
        // 被唤醒后, 没有阻塞的put了
        cnt_put -= 1;
        T tmp = Q.front();
        Q.pop();
        return tmp;
    }
    while (size_ == 0) {
        notEmpty.wait(std::ref(lck));
    }
    assert(get_idx_ < capacity_ && get_idx_ >= 0);
    T val = queue_[get_idx_];
    get_idx_++;
    size_--;
    get_idx_ = get_idx_ % capacity_;
    notFull.notify_one();
    lck.unlock();
    return val;
}

template<typename T>
bool Chan<T>::empty() const {
    return size_ == 0;
}

template<typename T>
bool Chan<T>::full() const {
    return size_ == capacity_;
}

template<typename T>
size_t Chan<T>::size() const {
    return size_;
}

template<typename T>
size_t Chan<T>::capacity() const {
    return capacity_;
}

#endif
