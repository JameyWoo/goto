/*
 * @Author: 姬小野
 * @Date: 2020-12-11 12:09:04
 * @LastEditTime: 2020-12-11 12:12:10
 * @LastEditors: 姬小野
 * @Description:  
 * @FilePath: \CXXImplement\playground\lambda测试\find_if_test.cpp
 * @Hello, World!
 */

#include <algorithm>
#include <iostream>
#include <map>
#include <queue>
#include <unordered_map>
#include <vector>

using namespace std;

int main() {
    vector<int> v{3, 4, 7, 1, 2, 9, 3, 4};
    auto res = find_if(v.begin(), v.end(), [](int val) {
        return val < 5;
    });
    while (res != v.end()) {
        cout << *res << endl;
        res++;
    }
    // for (const auto& x: res) {
    //     cout << x << endl;
    // }
}