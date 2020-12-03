#include "channel.hpp"

Chan::Chan() {
    init(0, 0, 0, 0, 0);
}

Chan::Chan(int c) {
    init(0, c, 0, 0, 0);
}

void Chan::init(int l, int c, int wc, int iwc, int owc) {
    len = l;
    cap = c;
    wait_count = wc;
    in_wait_count = iwc;
    out_wait_count = owc;
}

int Chan::out() {
    int val;
    if (Q.size() > 0) {
        val = Q.front();
        Q.pop();
        LOG_F(INFO, "out: I notify!");
        in_cv.notify_one();
    } else if (len == 0) {
        if (in_wait_count > 0) {
            // 如果队列里没有了值, 检查 阻塞队列
            // rwmtx.lock();
            val = waitQ.front();
            waitQ.pop();
            in_wait_count--;
            LOG_F(INFO, "in_wait_count--;");
            // rwmtx.unlock();
            LOG_F(INFO, "out: I notify as in!");
            in_cv.notify_one();
        } else {
            // 如果都没有, 那么才阻塞这个线程
            std::unique_lock<std::mutex> lock(mtx);
            LOG_F(INFO, "out: I block!");
            out_wait_count++;
            LOG_F(INFO, "out_wait_count++;");
            // 调用wait的时候会解锁!
            out_cv.wait(lock);
            // 被唤醒之后, 需要取一个数据. 从Q或者是waitQ中
            LOG_F(INFO, "out: block ok!");
            // 加锁
            val = outQ.front();
            outQ.pop();
        }
    }
    return val;
}

void Chan::in(int in_val) {
//    LOG_F(INFO, "in start");
    // 先判断in wait队列是否有值, 加入到队列中
    // ! 注意加锁
    while (len < cap && in_wait_count > 0) {
        // rwmtx.lock();
        Q.push(waitQ.front());
        waitQ.pop();
        len++;
        in_wait_count--;
        // rwmtx.unlock();
    }
//    LOG_F(INFO, "cap = %d len = %d", cap, len);
    // 如果没有缓冲, 则阻塞
    if (cap == len) {
        // rwmtx.lock();
        // ! 将导致阻塞的值保存在 一个新的队列
        waitQ.push(in_val);
        in_wait_count++;
        LOG_F(INFO, "in_wait_count++");
        // rwmtx.unlock();
        std::unique_lock<std::mutex> lock(mtx);
        if (out_wait_count > 0) {
            out_wait_count--;
            LOG_F(INFO, "out_wait_count--;");
            in_wait_count--;
            LOG_F(INFO, "in_wait_count--;");
            outQ.push(waitQ.front());
            waitQ.pop();
            out_cv.notify_one();
        } else {
            // 阻塞
            LOG_F(INFO, "in: I block!");
            in_cv.wait(lock);
            LOG_F(INFO, "out: block ok!");
        }

    } else if (len < cap) {  // 如果队列没有满, 则加入到队列中
        // 将value值添加进队列.
        // go中是一个slice实现的队列, C++可以使用stl的队列
        // rwmtx.lock();
        Q.push(in_val);
        len++;
        // rwmtx.unlock();
        LOG_F(INFO, "in: I notify!");
        out_cv.notify_one();
    }
}