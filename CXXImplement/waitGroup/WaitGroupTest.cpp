//
// Created by 26220 on 2020/12/4.
//

#include <thread>
#include <chrono>

#include "WaitGroup.hpp"

int main() {
    WaitGroup wg;
    for (int i = 0; i < 10; i++) {
        wg.Add(1);
        std::thread([&wg, i]() {
            std::this_thread::sleep_for(std::chrono::seconds(1 + (i*i)%10));
            LOG_F(INFO, "i = %d", i);
            wg.Done();
        }).detach();
    }
    wg.Wait();
}