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
/*
 * ! hello
 * * hello
 * + hello
 * ? hello! hello
 */
/*
 * 你好
 * 我是谁
 */
// ! hello
// + world
// ? what?
// * attention
// TOOD: hello
int main() {
    Vec v0(0);
//    Vec v1(NULL);
    Vec v2(nullptr);

    cout << NULL << endl;
//    cout << nullptr << endl;
}