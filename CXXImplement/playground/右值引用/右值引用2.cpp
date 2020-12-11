/*
 * @Author: your name
 * @Date: 2020-12-10 22:00:44
 * @LastEditTime: 2020-12-10 22:35:02
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \CXXImplement\playground\右值引用\std-move函数测试.cpp
 */

#include <iostream>
#include <type_traits>
#include <utility>

class A {
  public:
    A(size_t size): size(size), array((int*) malloc(size)) {
        std::cout 
          << "constructor called." 
          << std::endl;
    }
    ~A() {
        free(array);
    }

    A& operator=(A&& a) {
        a.array = nullptr;
        std::cout 
          << "operator= xvalue copied, memory at: " 
          << array 
          << std::endl;
        return *this;
    }

    A& operator=(const A& a) {
        std::cout 
          << "A& operator=(const A& a), memory at: " 
          << array 
          << std::endl;
        return *this;
    }

    A(A &&a) : array(a.array), size(a.size) {
        a.array = nullptr;
        std::cout 
          << "xvalue copied, memory at: " 
          << array 
          << std::endl;
    }
    A(const A &a) : size(a.size) {
        array = (int*) malloc(a.size);
        std::cout 
          << "normal copied, memory at: " 
          << array << std::endl;
    }

    size_t size;
    int *array;
};


int main (int argc, char **argv) {
    auto getTempA = [](size_t size = 100) -> A {
        auto tmp = A(size);
        std::cout << "Memory at: " << tmp.array << std::endl;

        return tmp;
    };

    std::cout 
      << std::is_rvalue_reference<decltype(getTempA())&&>::value 
      << std::endl;  // true;

    std::cout << "1 --- " << std::endl;
    auto x = getTempA(1000);
    std::cout << "2 --- " << std::endl;
    A z = A(100);
    std::cout << "3 --- " << std::endl;
    auto y = x;
    std::cout << "4 --- " << std::endl;

    return 0;   
}