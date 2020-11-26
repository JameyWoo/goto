#include <iostream>
#include <thread>
#include <chrono>

#include "channel.hpp"

std::mutex mtx, mtx2;

void in_test(Chan *chp, int val) {
     std::this_thread::sleep_for(std::chrono::seconds(1));
//    mtx2.lock();
//    std::cout << "in: " << val << std::endl;
//    mtx2.unlock();
    chp->in(val);
    // std::cout << "in test" << std::endl;
}

void out_test(Chan *chp) {
    // std::this_thread::sleep_for(std::chrono::seconds(1)); //sleep
    int out_val = chp->out();
    mtx.lock();
    std::cout << "out: " << out_val << std::endl;
    mtx.unlock();
    // std::cout << "out test: " << chp->out() << std::endl;
}

int main() {
    const int chan_cap = 0;
    const int thread_count = 2;

    Chan chan(chan_cap);

    std::thread th1s[thread_count];
    for (int i = 0; i < thread_count; i++) {
//        std::cout << "main: th1s: " << i << std::endl;
        th1s[i] = std::thread([&chan, i]() {
//            mtx.lock();
//            std::cout << "i = " << i << std::endl;
//            mtx.unlock();
            in_test(&chan, i);
        });
    }

    std::thread th2s[thread_count];
    for (int i = 0; i < thread_count; i++) {
//        std::cout << "main: th2s: " << i << std::endl;
        th2s[i] = std::thread([&chan]() {
            out_test(&chan);
        });
    }

    for (int i = 0; i < thread_count; i++) {
        th1s[i].join();
        th2s[i].join();
    }
}