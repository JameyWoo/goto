/*
 * @Author: 姬小野
 * @Date: 2020-12-13 20:05:28
 * @LastEditTime: 2020-12-13 20:10:11
 * @LastEditors: 姬小野
 * @Description: 
 * @FilePath: \CXXImplemente:\virtualShare\gopath3\src\github.com\Jameywoo\goto\playground\old\对比C++\fib.cpp
 * @Hello, World!
 */
#include <chrono>
#include <iostream>

using std::chrono::high_resolution_clock;
using std::chrono::milliseconds;

using namespace std;

using int64 = long long;

/*
func fib(n int64) int64 {
	if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
*/

int64 fib(int64 n) {
    if (n <= 2) {
        return 1;
    }
    return fib(n - 1) + fib(n - 2);
}

int main() {
    high_resolution_clock::time_point beginTime = high_resolution_clock::now();
    cout << fib(45) << endl;
    high_resolution_clock::time_point endTime   = high_resolution_clock::now();
    milliseconds timeInterval                   = std::chrono::duration_cast<milliseconds>(endTime - beginTime);
    std::cout << timeInterval.count() << "ms\n";
}