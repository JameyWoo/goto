/*
 * @Author: 姬小野
 * @Date: 2020-12-11 19:13:49
 * @LastEditTime: 2020-12-11 19:25:02
 * @LastEditors: 姬小野
 * @Description:  
 * @FilePath: \CXXImplement\playground\线程测试\atomic.cpp
 * @Hello, World!
 */
#include <algorithm>
#include <atomic>
#include <iostream>
#include <map>
#include <queue>
#include <thread>
#include <unordered_map>
#include <vector>

using namespace std;

int main() {
    int x = 0;
    atomic<int> y{0};
    vector<thread> ts;

    for (int i = 0; i < 1000; i++) {
        ts.emplace_back([&x, &y]() {
            for (int j = 0; j < 10000; j++) {
                x++;
                y++;
            }
        });
    }
    for_each(ts.begin(), ts.end(), [](thread& t) {
        t.join();
    });
    cout << "x = " << x << ", y = " << y << endl;
}