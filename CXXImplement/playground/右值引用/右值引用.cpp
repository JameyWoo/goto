/*
 * @Author: your name
 * @Date: 2020-12-10 21:33:32
 * @LastEditTime: 2020-12-11 20:24:08
 * @LastEditors: 姬小野
 * @Description: In User Settings Edit
 * @FilePath: \CXXImplement\playground\右值引用\右值引用.cpp
 */

#include <algorithm>
#include <iostream>
#include <map>
#include <queue>
#include <unordered_map>
#include <vector>

using namespace std;

class Intvec {
   public:
    explicit Intvec(size_t num = 0)
        : m_size(num), m_data(new int[m_size]) {
        log("constructor");
    }

    ~Intvec() {
        log("destructor");
        if (m_data) {
            delete[] m_data;
            m_data = 0;
        }
    }

    Intvec(const Intvec& other)
        : m_size(other.m_size), m_data(new int[m_size]) {
        log("copy constructor");
        for (size_t i = 0; i < m_size; ++i)
            m_data[i] = other.m_data[i];
    }

    Intvec(Intvec&& other)
        : m_size(other.m_size), m_data(new int[m_size]) {
        log("move constructor");
        for (size_t i = 0; i < m_size; ++i)
            m_data[i] = other.m_data[i];
    }

    Intvec& operator=(const Intvec& other) {
        log("copy assignment operator");
        Intvec tmp(other);
        std::swap(m_size, tmp.m_size);
        std::swap(m_data, tmp.m_data);
        return *this;
    }

    Intvec& operator=(Intvec&& other) {
        log("move assignment operator");
        std::swap(m_size, other.m_size);
        std::swap(m_data, other.m_data);
        return *this;
    }

   private:
    void log(const char* msg) {
        cout << "[" << this << "] " << msg << "\n";
    }

    size_t m_size;
    int* m_data;
};

int main() {
    Intvec v1(20);
    Intvec v2;
    cout << "assigning rvalue...\n";
    // v2 = v1;
    // v2 = Intvec(33);  //const invevec &
    Intvec v3(v1);
    auto v4(move(v1));
    cout << "ended assigning rvalue...\n";
}