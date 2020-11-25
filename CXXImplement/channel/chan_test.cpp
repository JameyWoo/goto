#include <iostream>
#include <thread>
#include <chrono>

#include "channel.hpp"

void in_test(Chan *chp) {
    // std::this_thread::sleep_for(std::chrono::seconds(1));
    chp->in(1234567);
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
    // chan.out();
    // chan.in(10);

    std::thread t1([&chan]() {
        in_test(&chan);
    });

    std::thread t2([&chan]() {
        out_test(&chan);
    });

    t1.join();
    t2.join();
}