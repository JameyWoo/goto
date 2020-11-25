#include "channel.hpp"

int Chan::out() {
    // std::cout << "--- out ---" << std::endl;
    // io_mtx.lock();
    
    if (len == 0) {
        std::unique_lock<std::mutex> lock(mtx);
        std::cout << "out: I block!\n";
        // ! 调用wait的时候会解锁! 所以要控制out和in的互斥需要另一个锁
        cv.wait(lock);
        std::cout << "out: block ok!\n";
    } else if (len > 0) {
        
    }
    int val = Q.front();
    Q.pop();
    std::cout << "out: I notify!\n";
    cv.notify_one();
    return val;
}

void Chan::in(int in_val) {
    // io_mtx.lock();
    // std::cout << "--- in ---" << std::endl;
    // 如果没有缓冲, 则阻塞
    if (cap == len) {
        // 阻塞
        // ? 将导致阻塞的值保存在哪里?
        std::unique_lock<std::mutex> lock(mtx);
        std::cout << "in: I block!\n";
        cv.wait(lock);
        std::cout << "out: block ok!\n";
    } else if (len < cap) {
        // 将value值添加进队列. 
        // go中是一个slice实现的队列, C++可以使用stl的队列
        Q.push(in_val);
        len++;
        std::cout << "int: I notify!\n";
        cv.notify_one();
    }
    io_mtx.unlock();
}