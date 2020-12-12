/*
 * @Author: 姬小野
 * @Date: 2020-12-12 16:11:42
 * @LastEditTime: 2020-12-12 16:16:16
 * @LastEditors: 姬小野
 * @Description: Effective C++ p71页了解到这个"隐式转换函数", 以前从来没听过. 现在理解了一些C++的转换规则了. 
 * @FilePath: \CXXImplement\playground\隐式转换函数测试.cpp
 * @Hello, World!
 */
#include <algorithm>
#include <iostream>
#include <map>
#include <queue>
#include <unordered_map>
#include <vector>

using namespace std;

class A {
   public:
    int a = 0;
    A(int x) : a(x) {}
    A() = default;
};

class B {
   public:
    A t;
    B(int x) : t(x) {}
    B() = default;
    operator A() const { return t; }
};

int main() {
    B b(110);
    A a = b;
    cout << a.a << endl;
}