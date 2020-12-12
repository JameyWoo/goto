/*
 * @Author: 姬小野
 * @Date: 2020-12-12 20:56:56
 * @LastEditTime: 2020-12-12 20:58:06
 * @LastEditors: 姬小野
 * @Description: 
 * @FilePath: \CXXImplement\playground\智能指针\weak_ptr_test.cpp
 * @Hello, World!
 */
#include <iostream>
#include <memory>

using namespace std;

int main() {
    auto sp1 = make_shared<int>(10);
    weak_ptr<int> wp(sp1);  //通过shared_ptr初始化
    weak_ptr<int> wp1, wp2;
    wp1      = sp1;         //利用shared_ptr来赋值
    wp2      = wp;          //利用weak_ptr赋值
    auto sp2 = wp2.lock();  //sp2为shared_ptr类型

    cout << wp2.use_count() << endl;

    sp1 = nullptr;

    cout << wp2.use_count() << endl;  //1，强引用计数
    return 0;
}