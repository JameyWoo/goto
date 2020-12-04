//
// Created by 26220 on 2020/12/4.
//

#include <iostream>
#include <thread>
#include <chrono>

#include "channel.hpp"
#include "../log/loguru.hpp"


int main() {
    Chan<std::string> ch;

    std::thread t1 = std::thread([&ch]() {
        while (true) {
            std::this_thread::sleep_for(std::chrono::seconds(1)); //sleep
            std::string out_val = ch.get();
            LOG_F(WARNING, "out: %s", out_val.c_str());
        }
    });

    std::thread t2 = std::thread([&ch]() {
        while (true) {
            std::this_thread::sleep_for(std::chrono::seconds(3)); //sleep
            ch.put(std::string("hello, world"));
            LOG_F(WARNING, "put over~");
        }
    });

    // 使用detach(),main函数不用等待线程结束才能结束
    // 不想用join, 也不想出现异常
    t1.detach();
    t2.detach();

    std::this_thread::sleep_for(std::chrono::seconds(5));
}