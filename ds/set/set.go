package set

import "sync"

type inter interface {}

// 要有 Insert, Delete, Find, Size 方法
type Set struct {
	m map[interface{}]bool
	sync.RWMutex
}

func NewSet() *Set {
	return &Set{
		m: map[interface{}]bool{},
	}
}

func (s *Set) Size() int {
	return len(s.m)
}

func (s *Set) Insert(str interface{}) {
	s.Lock()
	s.m[str] = true
	s.Unlock()
}

func (s *Set) Delete(str interface{}) {
	s.Lock()
	if _, ok := s.m[str]; ok {
		delete(s.m, str)
	}
	s.Unlock()
}

func (s *Set) Find(str interface{}) bool {
	s.RLock()
	defer s.RUnlock()
	return s.m[str]
}

// 迭代器, 获取所有的key
func (s *Set) Iter() []interface{} {
	var iter []interface{}
	for x, _ := range s.m {
		iter = append(iter, x)
	}
	return iter
}