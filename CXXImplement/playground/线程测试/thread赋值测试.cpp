/*
 * @Author: 姬小野
 * @Date: 2020-12-11 19:25:29
 * @LastEditTime: 2020-12-11 19:33:11
 * @LastEditors: 姬小野
 * @Description:  
 * @FilePath: \CXXImplement\playground\线程测试\thread赋值测试.cpp
 * @Hello, World!
 */
#include <algorithm>
#include <iostream>
#include <thread>
#include <vector>

using namespace std;

int main() {
    thread t = thread([] { cout << "hello" << endl; });
    t.join();

    vector<thread> ts;
    ts.emplace_back([] { cout << "hello" << endl; });
    ts.emplace_back([] { cout << "world" << endl; });
    for_each(ts.begin(), ts.end(), [](thread& t) { t.join(); });
}
