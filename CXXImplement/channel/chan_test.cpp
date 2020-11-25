#include <iostream>
#include <thread>
#include <chrono>

#include "channel.hpp"

void in_test(Chan *chp, int val) {
    // std::this_thread::sleep_for(std::chrono::seconds(1));
    chp->in(val);
    // std::cout << "in test" << std::endl;
}

void out_test(Chan *chp) {
    // std::this_thread::sleep_for(std::chrono::seconds(1)); //sleep
    int out_val = chp->out();
    std::cout << "out: " << out_val << std::endl;
    // std::cout << "out test: " << chp->out() << std::endl;
}

int main() {
    Chan chan(0);

    std::thread th1s[10];
    for (int i = 0; i < 10; i++) {
        std::cout << "main: th1s: " << i << std::endl;
        std::thread t1([&chan, &i]() {
            in_test(&chan, i);
        });
    }

    std::thread th2s[10];
    for (int i = 0; i < 10; i++) {
        std::cout << "main: th2s: " << i << std::endl;
        std::thread t2([&chan]() {
            out_test(&chan);
        });
    }

    for (int i = 0; i < 10; i++) {
        th1s[i].join();
        th2s[i].join();
    }
}