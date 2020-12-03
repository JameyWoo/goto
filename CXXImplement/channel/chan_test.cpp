#include <iostream>
#include <thread>
#include <chrono>

#include "channel.hpp"
#include "../log/loguru.hpp"

std::mutex mtx, mtx2;

void in_test(Chan *chp, int val) {
//     std::this_thread::sleep_for(std::chrono::seconds(1));
    chp->in(val);
}

void out_test(Chan *chp) {
    // std::this_thread::sleep_for(std::chrono::seconds(1)); //sleep
    int out_val = chp->out();
    LOG_F(WARNING, "out: %d", out_val);
}

int main() {
    loguru::g_colorlogtostderr = false;
    const int chan_cap = 0;
    const int thread_count = 10;

    Chan chan(chan_cap);

    std::thread th1s[thread_count];
    for (int i = 0; i < thread_count; i++) {
        th1s[i] = std::thread([&chan, i]() {
            in_test(&chan, i);
        });
    }

    std::thread th2s[thread_count];
    for (int i = 0; i < thread_count; i++) {
        th2s[i] = std::thread([&chan]() {
            out_test(&chan);
        });
    }

    for (int i = 0; i < thread_count; i++) {
        th1s[i].join();
        th2s[i].join();
    }
}