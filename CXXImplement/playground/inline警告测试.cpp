/*
 * @Author: 姬小野
 * @Date: 2020-12-12 16:58:23
 * @LastEditTime: 2020-12-12 17:03:34
 * @LastEditors: 姬小野
 * @Description: 
 * @FilePath: \CXXImplement\playground\inline警告测试.cpp
 * @Hello, World!
 */
#include <algorithm>
#include <iostream>
#include <map>
#include <queue>
#include <unordered_map>
#include <vector>

using namespace std;

// 无警告
inline void test() {
    cout << "hello" << endl;
}

inline void for_test() {
    for (int i = 0; i < 10000; i++) {
        // cout << "world" << endl;
    }
}

inline void dg_test(int n) {
    if (n == 0) return;
    dg_test(n - 1);
}

int main() {
    test();
    for_test();
    dg_test(10000);
}