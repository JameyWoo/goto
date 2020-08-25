package set

import "sync"

type inter interface {}

// 要有 Insert, Delete, Find, Size 方法
type set struct {
	m map[interface{}]bool
	sync.RWMutex
}

func NewSet() *set {
	return &set{
		m: map[interface{}]bool{},
	}
}

func (s *set) Size() int {
	return len(s.m)
}

func (s *set) Insert(str interface{}) {
	s.Lock()
	s.m[str] = true
	s.Unlock()
}

func (s *set) Delete(str interface{}) {
	s.Lock()
	if _, ok := s.m[str]; ok {
		delete(s.m, str)
	}
	s.Unlock()
}

func (s *set) Find(str interface{}) bool {
	s.RLock()
	defer s.RUnlock()
	return s.m[str]
}

// 迭代器, 获取所有的key
func (s *set) Iter() []interface{} {
	var iter []interface{}
	for x, _ := range s.m {
		iter = append(iter, x)
	}
	return iter
}