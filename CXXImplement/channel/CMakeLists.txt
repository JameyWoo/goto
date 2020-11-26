cmake_minimum_required(VERSION 3.17)
project(chan_test)

set(CMAKE_CXX_STANDARD 14)

add_executable(chan_test
        chan_test.cpp
        channel.cpp
        channel.hpp)