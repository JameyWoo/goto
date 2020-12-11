/*
 * @Author: your name
 * @Date: 2020-12-10 22:35:26
 * @LastEditTime: 2020-12-10 22:35:52
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \CXXImplement\playground\右值引用\std_move测试.cpp
 */

#include <algorithm>
#include <iostream>
#include <map>
#include <queue>
#include <unordered_map>
#include <vector>

using namespace std;

int main() {
    int&& a = move(2);
}