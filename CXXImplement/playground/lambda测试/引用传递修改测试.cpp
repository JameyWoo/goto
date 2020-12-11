/*
 * @Author: 姬小野
 * @Date: 2020-12-11 18:07:13
 * @LastEditTime: 2020-12-11 18:10:18
 * @LastEditors: 姬小野
 * @Description:  
 * @FilePath: \CXXImplement\playground\lambda测试\引用传递修改测试.cpp
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
    [&](){a = 6;}();
    cout << a << endl;

    [a]()mutable{a=5;}();

    [=]()mutable{a=3;}();
}