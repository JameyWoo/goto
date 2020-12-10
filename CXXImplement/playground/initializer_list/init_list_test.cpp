#include <iostream>

using namespace std;

class Vec {
public:
    Vec() {}
    Vec(initializer_list<int> nums) {
        for (auto& v: nums) {
            cout << v << ' ';
        }
        cout << endl;
    }
};

int main() {
    Vec v{1, 2, 3, 4, 5, 6.0};
}