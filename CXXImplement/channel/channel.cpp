#include "channel.hpp"

void debug(std::string s);

int Chan::out() {
    int val;
    if (len == 0) {
        if (in_wait_count > 0) {
            // 如果队列里没有了值, 检查 阻塞队列
            val = in_next.front();
            in_next.pop();
            in_wait_count--;
        } else {
            // 如果都没有, 那么才阻塞这个线程
            std::unique_lock<std::mutex> lock(mtx);
            std::cout << "out: I block!\n";
            // 调用wait的时候会解锁!
            out_cv.wait(lock);
            // 被唤醒之后, 需要取一个数据. 从Q或者是in_next中
            std::cout << "out: block ok!\n";
        }
    }
    if (len > 0) {
        val = Q.front();
        Q.pop();
    } else if (in_wait_count > 0) {
        val = in_next.front();
        in_next.pop();
        in_wait_count--;
    }

    std::cout << "out: I notify!\n";
    in_cv.notify_one();
    return val;
}

void Chan::in(int in_val) {
    debug("in start");
    // 先判断in wait队列是否有值, 加入到队列中
    // ! 注意加锁
    while (len < cap && in_wait_count > 0) {
        Q.push(in_next.front());
        len++;
        in_wait_count--;
        in_next.pop();
    }
    std::cout << "cap = " << cap << " len = " << len << std::endl;
    // 如果没有缓冲, 则阻塞
    if (cap == len) {
        // ! 将导致阻塞的值保存在 一个新的队列
        in_next.push(in_val);
        in_wait_count++;
        if (out_wait_count > 0) {
            out_cv.notify_one();
            out_wait_count--;
        } else {
            // 阻塞
            std::unique_lock<std::mutex> lock(mtx);
            std::cout << "in: I block!\n";
            in_cv.wait(lock);
            std::cout << "out: block ok!\n";
        }

    } else if (len < cap) {  // 如果队列没有满, 则加入到队列中
        // 将value值添加进队列.
        // go中是一个slice实现的队列, C++可以使用stl的队列
        Q.push(in_val);
        len++;
        std::cout << "int: I notify!\n";
        out_cv.notify_one();
    }
}

void debug(std::string s) {
    std::mutex mu;
    mu.lock();
    std::cout << s << std::endl;
    mu.unlock();
}