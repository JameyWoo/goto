// 博客: https://www.liwenzhou.com/posts/Go/singleton_in_go/

package main

/*
激进的加锁方式的问题在于, 判断nil这一步只会在创建的时候有, 创建好之后就不需要了.
而加了一个全局锁, 会导致即使实例已经被创建, 也需要参与获取锁

而用两段检查的方式, 如果实例已被创建, 那么直接返回即可
 */