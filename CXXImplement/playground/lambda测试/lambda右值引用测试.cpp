/*
 * @Author: 姬小野
 * @Date: 2020-12-11 15:56:01
 * @LastEditTime: 2020-12-11 16:04:44
 * @LastEditors: 姬小野
 * @Description:  
 * @FilePath: \CXXImplement\playground\lambda测试\lambda右值引用测试.cpp
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
    auto f = [x = move(a)]() mutable {
        int&& nx = move(x);
        nx = 18;
        x = 17;
    };
    cout << a << endl;
}