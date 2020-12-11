/*
 * @Author: 姬小野
 * @Date: 2020-12-11 15:48:51
 * @LastEditTime: 2020-12-11 15:50:25
 * @LastEditors: 姬小野
 * @Description:  
 * @FilePath: \CXXImplement\playground\lambda测试\成员函数const测试.cpp
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
    int a;
    void modify() {
        a = 10;
    }
    // 加了 const 的函数不可以修改类的成员变量
    void test() const {
        // a = 10;
        int b;
        b = 10;
    }
};

int main() {
    A a;
}