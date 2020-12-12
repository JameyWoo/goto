/*
 * @Author: 姬小野
 * @Date: 2020-12-12 19:56:47
 * @LastEditTime: 2020-12-12 19:57:18
 * @LastEditors: 姬小野
 * @Description: 
 * @FilePath: \CXXImplement\playground\智能指针\unique_ptr_test2.cpp
 * @Hello, World!
 */
#include <iostream>
#include <memory>  //for smart pointer
#include <vector>

using namespace std;

class Widget {};

//返回值RVO优化：
unique_ptr<int> func() {
    unique_ptr<int> up(new int(100));
    return up;  //up是个左值，调用拷贝构造给返回值？ No。
                //C++标准要求当RVO被允许时，要么消除拷贝，要么隐式地把std::move用在要返回的局部
                //对象上去。这里编译器会直接在返回值位置创建up对象。因此根本不会发生拷贝构造，
                //unique_ptr本身也不能被拷贝构造。

    //return unique_ptr<int>(new int(100)); //右值，被移动构造。
}

void foo(std::unique_ptr<int> ptr) {
}

void myDeleter(int* p) {
    cout << "invoke deleter(void* p)" << endl;
    delete p;
}

int main() {
    //1. unique_ptr的初始化
    //1.1 通过裸指针创建unique_ptr（由于unique_ptr的构造函数是explicit的，必须使用直接初始化，不能做隐式类型转换）
    std::unique_ptr<Widget> ptr1(new Widget);  //ok; 直接初始化
    //std::unique_ptr<Widget> ptr1 = new Widget(); //error。不能隐式将Widget*转换为unqiue_ptr<Widget>类型。

    std::unique_ptr<int[]> ptr2(new int[10]);  //指向数组

    //1.2 通过移动构造
    //std::unique_ptr<Widget> ptr3 = ptr1;    //error，unique_ptr是独占型，不能复制构造
    std::unique_ptr<Widget> ptr3 = std::move(ptr1);  //ok，unique_ptr是个只移动类型，可以移动构造
    auto ptr4                    = std::move(ptr3);  //ok， ptr4为unique_ptr<Widget>类型

    //1.3 通过std::make_unique来创建
    auto ptr5 = std::make_unique<int>(10);

    //auto ptr6 = std::make_unique<vector<int>>({1,2,3,4,5}); //error，make_unique不支持初始化列表
    auto initList = {1, 2, 3, 4, 5};
    auto ptr6     = std::make_unique<vector<int>>(initList);

    //2. 传参和返回值
    int* px = new int(0);
    //foo(px); //error，px无法隐式转为unique_ptr。可防止foo函数执行完毕后，px会自动释放。
    //foo(ptr5); //error，智能指针不能被拷贝。因此，可以将foo的形参声明为引用，以避免所有权转移
    foo(std::move(ptr5));  //ok，通过移动构造

    auto ptr7 = func();  //移动构造

    //3.常用操作
    std::unique_ptr<Widget> upw1;  //空的unique_ptr
    upw1.reset(new Widget);
    std::unique_ptr<Widget> upw2(new Widget);

    cout << "before swap..." << endl;
    cout << "upw1.get() = " << hex << upw1.get() << endl;

    cout << "upw2.get() = " << hex << upw2.get() << endl;

    cout << "after swap..." << endl;
    upw1.swap(upw2);  //交换指针所指的对象
    cout << "upw1.get() = " << hex << upw1.get() << endl;
    cout << "upw2.get() = " << hex << upw2.get() << endl;

    //upw1.release(); //release放弃了控制权不会释放内存，丢失了指针
    Widget* pw = upw1.release();  //放弃对指针的控制
    delete pw;                    //需手动删除

    if (upw1) {  //unique_ptr重载了operator bool()
        cout << "upw1 owns resourse" << endl;
    } else {
        cout << "upw1 lost resourse" << endl;
    }

    upw1.reset(upw2.release());  //转移所有权
    cout << "upw1.get() = " << hex << upw1.get() << endl;
    cout << "upw2.get() = " << hex << upw2.get() << endl;

    //upw1 = nullptr; //释放upw1指向的对象，并将upw1置空
    //upw1.reset(nullptr);

    //4.unique_ptr的大小
    std::unique_ptr<int, decltype(&myDeleter)> upd1(new int(0), myDeleter);  //自定义删除器
    auto del = [](auto* p) { delete p; };
    std::unique_ptr<int, decltype(del)> upd2(new int(0), del);
    cout << sizeof(upw1) << endl;  //4字节，默认删除器
    cout << sizeof(upd1) << endl;  //8字节
    cout << sizeof(upd2) << endl;  //4字节

    return 0;
}