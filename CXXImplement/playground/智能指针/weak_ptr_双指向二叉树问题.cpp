/*
 * @Author: 姬小野
 * @Date: 2020-12-12 21:01:44
 * @LastEditTime: 2020-12-12 21:13:01
 * @LastEditors: 姬小野
 * @Description: 
 * @FilePath: \CXXImplement\playground\智能指针\weak_ptr_双指向二叉树问题.cpp
 * @Hello, World!
 */

#include <iostream>
#include <memory>

class Node {
    int value;

   public:
    std::shared_ptr<Node> leftPtr;
    std::shared_ptr<Node> rightPtr;
    // * 只需要把shared_ptr改为weak_ptr;
    // ! 但显然需要自己判断可能导致循环引用的部分
    std::weak_ptr<Node> parentPtr;
    Node(int val) : value(val) {
        std::cout << "Constructor" << std::endl;
    }
    ~Node() {
        std::cout << "Destructor" << std::endl;
    }
};

int main() {
    std::shared_ptr<Node> ptr = std::make_shared<Node>(4);
    ptr->leftPtr              = std::make_shared<Node>(2);
    ptr->leftPtr->parentPtr   = ptr;
    ptr->rightPtr             = std::make_shared<Node>(5);
    ptr->rightPtr->parentPtr  = ptr;
    std::cout << "ptr reference count = " << ptr.use_count() << std::endl;
    std::cout << "ptr->leftPtr reference count = " << ptr->leftPtr.use_count() << std::endl;
    std::cout << "ptr->rightPtr reference count = " << ptr->rightPtr.use_count() << std::endl;

    auto par = ptr->leftPtr->parentPtr.lock();
    std::cout << par.use_count() << std::endl;  // * 有两个子节点指向了它. 虚引用的个数
    if (par.get() == ptr.get()) {  // 都是父指针
        std::cout << "bingo\n";
    }
    std::cout << ptr->leftPtr->parentPtr.expired() << std::endl;
    
    auto wp = ptr->leftPtr->parentPtr;
    std::weak_ptr<Node> ww = ptr;
    std::cout << wp.use_count() << std::endl;
    return 0;
}