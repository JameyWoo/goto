#include <iostream>
#include <thread>
#include <chrono>

#include "channel.hpp"
#include "../log/loguru.hpp"

void in_test(Chan<int> *chp, int val) {
//     std::this_thread::sleep_for(std::chrono::seconds(1));
    chp->put(val);
}

void out_test(Chan<int> *chp) {
    // std::this_thread::sleep_for(std::chrono::seconds(1)); //sleep
    int out_val = chp->get();
    LOG_F(WARNING, "out: %d", out_val);
}

int main() {
    const int chan_cap = 1;
    const int thread_count = 12;

    Chan<int> chan(chan_cap);

    std::thread th2s[thread_count];
    for (int i = 0; i < thread_count; i++) {
        th2s[i] = std::thread([&chan]() {
            out_test(&chan);
        });
    }

    std::thread th1s[thread_count];
    for (int i = 0; i < thread_count; i++) {
        th1s[i] = std::thread([&chan, i]() {
            in_test(&chan, i);
        });
    }

    for (int i = 0; i < thread_count; i++) {
        th1s[i].join();
        th2s[i].join();
    }
}