/*
 * @Author: 姬小野
 * @Date: 2020-12-12 20:33:58
 * @LastEditTime: 2020-12-12 20:37:14
 * @LastEditors: 姬小野
 * @Description: 
 * @FilePath: \CXXImplement\playground\智能指针\shared_ptr裸指针测试.cpp
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

class Widget {
    public:
    int x = 110;
};

void func(shared_ptr<Widget> sp) {}

int main() {
    // 它被用作智能指针那么不应该以左值出现! 只能右值!
    auto sp = new Widget;
    // ! 造成悬空指针!
    func(shared_ptr<Widget>(sp));  //慎用裸指针，sp将在func结束后被释放！
    cout << (sp == nullptr) << endl;
    // ! 悬空指针
    cout << sp->x << endl;
}