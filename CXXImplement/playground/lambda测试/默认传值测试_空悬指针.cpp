/*
 * @Author: 姬小野
 * @Date: 2020-12-11 13:14:57
 * @LastEditTime: 2020-12-11 13:30:31
 * @LastEditors: 姬小野
 * @Description:  
 * @FilePath: \CXXImplement\playground\lambda测试\默认传值测试_空悬指针.cpp
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
    auto test() {
        a = 10;
        auto f = [=]() {  // 会空悬
        // auto f = [a = a]() {  // 不会受空悬影响
            cout << "fun !\n";
            // a = 5;
            cout << a << endl;
            cout << &a << endl;
        };
        return f;
    }
};

void test() {
    A* a = new A();
    a->test()();
    a->a = 11;
    cout << &a->a << endl;
    auto f = a->test();
    // ! 造成空悬指针!
    delete a;
    f();
}

int main() {
    test();
}