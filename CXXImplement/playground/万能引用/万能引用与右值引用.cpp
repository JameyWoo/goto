/*
 * @Author: 姬小野
 * @Date: 2020-12-11 11:10:46
 * @LastEditTime: 2020-12-11 11:36:32
 * @LastEditors: 姬小野
 * @Description: 万能引用与右值引用
 * @FilePath: \CXXImplement\playground\万能引用\万能引用与右值引用.cpp
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
void isR(T&& x) {
    cout << is_rvalue_reference<decltype(x)&&>::value << endl;
}

void isRTest() {
    cout << "isRTest--------------------" << endl;
    isR(1);
    int a;
    isR(a);
    isR(move(a));
    cout << "isRTest--------------------" << endl;
}

void t(int&& x) {
    cout << "int&& x" << endl;
}

void t(int& x) {
    cout << "int& x" << endl;
}

void test2() {
    cout << "test2--------------------" << endl;
    t(1);
    int a;
    t(a);
    cout << "test2--------------------" << endl;
}

void auto_test() {
    cout << "auto -----------------" << endl;

    auto&& x = 1;
    cout << is_rvalue_reference<decltype(x)&&>::value << endl;

    int a = 1;
    auto&& y = a;
    cout << is_rvalue_reference<decltype(y)&&>::value << endl;

    // ! 不是万能引用! 现在是右值! 严格按照这个格式
    const auto&& z = move(a);
    cout << is_rvalue_reference<decltype(z)&&>::value << endl;

    cout << "auto -----------------" << endl;
}

int main() {
    isRTest();

    test2();

    auto_test();

    cout << is_rvalue_reference<decltype(1)&&>::value << endl;

    int a;
    cout << is_rvalue_reference<decltype(a)&&>::value << endl;

    auto&& b = a;
    cout << is_rvalue_reference<decltype(b)&&>::value << endl;
    
    int&& c = move(a);
    cout << is_rvalue_reference<decltype(c)&&>::value << endl;
}