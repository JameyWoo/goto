/*
 * @Author: 姬小野
 * @Date: 2020-12-12 20:42:30
 * @LastEditTime: 2020-12-12 20:54:49
 * @LastEditors: 姬小野
 * @Description: 
 * @FilePath: \CXXImplement\playground\智能指针\shared_ptr_deleter测试.cpp
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

class T {};

int main() {
    auto deleter = [](T* t) { cout << "deleter" << endl; delete t; };
    // auto p0       = make_shared<T>(T(), deleter);
    shared_ptr<T> p(new T(), deleter);
}