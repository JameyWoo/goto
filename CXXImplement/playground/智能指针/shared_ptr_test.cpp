/*
 * @Author: 姬小野
 * @Date: 2020-12-12 19:49:27
 * @LastEditTime: 2020-12-12 20:00:25
 * @LastEditors: 姬小野
 * @Description: 
 * @FilePath: \CXXImplement\playground\智能指针\shared_ptr_test.cpp
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
    Test(int x) : a(x) { cout << "arg Test(x)" << endl; }
    ~Test() {
        cout << "Test ~" << endl;
    }
};

void foo(shared_ptr<Test> p) {
    cout << "foo" << endl;
}

int main() {
    // auto p = make_shared<Test>();
    // shared_ptr<Test> p{new Test()};
    // shared_ptr<Test> p{new Test(7)};
    // auto p = make_shared<Test>(Test(7));  // * 这种方式用对象, 而不是用new的指针
    Test t(8);
    auto p = make_shared<Test>(t);

    auto p1 = p;
    cout << p.use_count() << endl;
    {
        auto p2 = p1;
        cout << p.use_count() << endl;
    }
    foo(move(p1));
    cout << p.use_count() << endl;
    cout << "p.a = " << p->a << endl;
}