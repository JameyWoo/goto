cmake_minimum_required(VERSION 3.10)
project(chan_test)

set(CMAKE_CXX_STANDARD 14)

add_executable(
        chan
        channel/chan.cpp
        channel/channel.hpp
        log/loguru.cpp
        log/loguru.hpp)

target_link_libraries(
        chan
        dl
        pthread
)

add_executable(
        channel_failed
        channel_failed/chan_test.cpp
        channel_failed/channel.cpp
        channel_failed/channel.hpp
        log/loguru.cpp
        log/loguru.hpp)

target_link_libraries(
        channel_failed
        dl
        pthread
)


add_executable(
        block_queue
        block_queue/bq_test.cpp
        block_queue/bq.hpp
        log/loguru.cpp
        log/loguru.hpp)

target_link_libraries(
        block_queue
        dl
        pthread
)

add_executable(
        bq
        block_queue/bq.cpp
        block_queue/bq.hpp
        log/loguru.cpp
        log/loguru.hpp)

target_link_libraries(
        bq
        dl
        pthread
)