/*
 * @Author: 姬小野
 * @Date: 2020-12-10 23:12:31
 * @LastEditTime: 2020-12-11 00:07:37
 * @LastEditors: 姬小野
 * @Description: 
 * @FilePath: \CXXImplement\playground\explicit关键字测试\explicit.cpp
 * @可以输入预定的版权声明、个性签名、空行等
 */

class Circle {
   public:
    explicit Circle(double r) : R(r) {}
    explicit Circle(int x, int y = 0) : X(x), Y(y) {}
    explicit Circle(const Circle& c) : R(c.R), X(c.X), Y(c.Y) {}

   private:
    double R;
    int X;
    int Y;
};

void test1() {
    //发生隐式类型转换
    //编译器会将它变成如下代码
    //tmp = Circle(1.23)
    //Circle A(tmp);
    //tmp.~Circle();

    // Circle A = 1.23;
    // //注意是int型的，调用的是Circle(int x, int y = 0)
    // //它虽然有2个参数，但后一个有默认值，任然能发生隐式转换
    // Circle B = 123;
    // //这个算隐式调用了拷贝构造函数
    // Circle C = A;

    Circle a(1);
    Circle A(a);
}

void test2() {
    Circle A(1);
}

int main() {
    test1();
    test2();

    return 0;
}