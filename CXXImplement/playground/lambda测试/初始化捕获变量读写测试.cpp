/*
 * @Author: 姬小野
 * @Date: 2020-12-11 15:51:07
 * @LastEditTime: 2020-12-11 16:03:41
 * @LastEditors: 姬小野
 * @Description:  
 * @FilePath: \CXXImplement\playground\lambda测试\初始化捕获变量读写测试.cpp
 * @Hello, World!
 */
#include <algorithm>
#include <iostream>
#include <map>
#include <queue>
#include <unordered_map>
#include <vector>

using namespace std;

int main() {
    int a = 10;
    auto f = [x = a]() {
        // ! error: assignment of read-only variable 'x'
        // x = 5;
        cout << x << endl;

        // * 用一个新的变量给他赋值然后读写
        int nx = x;
        nx = 7;
        cout << nx << endl;
    };
    f();

    auto f2 = [x = a]() mutable {
        x = 5;
        cout << x << endl;
    };
    f2();
}