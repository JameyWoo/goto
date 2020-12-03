cmake_minimum_required(VERSION 3.10)
project(chan_test)

set(CMAKE_CXX_STANDARD 14)

add_executable(
        chan_test
        channel/chan_test.cpp
        channel/channel.cpp
        channel/channel.hpp
        log/loguru.cpp
        log/loguru.hpp)

target_link_libraries(
        chan_test
        dl
        pthread
)