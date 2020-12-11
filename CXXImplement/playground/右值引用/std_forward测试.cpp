/*
 * @Author: 姬小野
 * @Date: 2020-12-11 11:01:30
 * @LastEditTime: 2020-12-11 11:05:17
 * @LastEditors: 姬小野
 * @Description: 
 * @FilePath: \CXXImplement\playground\右值引用\std_forward测试.cpp
 * @Hello, World!
 */
#include <algorithm>
#include <iostream>
#include <map>
#include <queue>
#include <unordered_map>
#include <vector>

using namespace std;

template<typename T>
void test(T&& p) {
    cout << std::is_rvalue_reference<decltype(p)&&>::value << '-';
    auto x = forward<T>(p);
    cout << std::is_rvalue_reference<decltype(x)&&>::value << endl;
}

int main() {
    int p = 1;
    test(p);

    int a = 1;
    p = a;
    test(p);

    test(1);
}