#include <iostream>

using namespace std;

class Vec {
public:
    Vec() {}
    Vec(int) {
        cout << "int" << endl;
    }
    Vec(void*) {
        cout << "pointer" << endl;
    }
};

int main() {
    Vec v0(0);
//    Vec v1(NULL);
    Vec v2(nullptr);

    cout << NULL << endl;
//    cout << nullptr << endl;
}