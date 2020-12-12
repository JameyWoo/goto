/*
 * @Author: 姬小野
 * @Date: 2020-12-12 19:35:35
 * @LastEditTime: 2020-12-12 19:49:14
 * @LastEditors: 姬小野
 * @Description: 
 * @FilePath: \CXXImplement\playground\智能指针\unique_ptr_test.cpp
 * @Hello, World!
 */
#include <algorithm>
#include <iostream>
#include <map>
#include <memory>
#include <queue>
#include <unordered_map>
#include <vector>

using namespace std;

class Test {
   public:
    int a;
    Test() {
        cout << "Test con" << endl;
    }
    ~Test() {
        cout << "Test ~" << endl;
    }
};

void foo(unique_ptr<Test> p) {
    cout << "over" << endl;
}

int main() {
    unique_ptr<Test> p = make_unique<Test>();
    p->a = 10;
    cout << p->a << endl;

    foo(move(p));
    cout << "after foo" << endl;
    // cout << p->a << endl;  // ! 不能再使用这个 p->a 了, 因为指针指向null了!
    cout << (p == nullptr) << endl;
}