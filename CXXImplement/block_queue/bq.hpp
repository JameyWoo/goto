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
class BoundedBlockingQueue {
public:
    // 默认
    explicit BoundedBlockingQueue() : BoundedBlockingQueue(1024) {}

    explicit BoundedBlockingQueue(int max_size) {
        capacity_ = max_size;
        queue_.resize(capacity_);

        size_ = 0;
        put_idx_ = 0;
        get_idx_ = 0;
        wait_get = 0;
        cnt_get = 0;
        cnt_put = 0;
    }


    void put(const T &x) {
        // 在unique_lock对象的声明周期内，它所管理的锁对象会一直保持上锁状态
        // unique_lock 具有 lock_guard 的所有功能，而且更为灵活
        std::unique_lock<std::mutex> lck(mutex_);
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
        if (wait_get > 0) {
            tmp_val = x;
            notEmpty.notify_one();
            return;
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

    void put(T &&x) {
        put(x);
    }

    T get() {
        std::unique_lock<std::mutex> lck(mutex_);
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

    bool empty() const {
        return size_ == 0;
    }

    bool full() const {
        return size_ == capacity_;
    }

    size_t size() const {
        return size_;
    }

    //
    size_t capacity() const {
        return capacity_;
    }

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
    int wait_get;  // 当空间为0时, 进行get的数量
    int wait_put;  // 当空间为0时, 进行put的数量
    T tmp_val;  // 当空间为0时, get和in进行通信的中间变量
    std::queue<T> Q;
    int cnt_get;
    int cnt_put;
};


#endif
