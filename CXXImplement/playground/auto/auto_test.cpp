#include <iostream>
#include <vector>
using namespace std;

int main() {
	auto f = [](auto& x) {
		cout << x << endl;
	};
	auto a = 12;
	f(a);
	auto s = "shit";
	f(s);

	auto f2 = [](auto& x, auto& y) {
	    cout << x + y << endl;
	};
	f2(a, s);

	vector<bool> bs({true, false, false});
	auto sec = bs[1];
	sec = true;
	cout << bs[1] << endl;
}