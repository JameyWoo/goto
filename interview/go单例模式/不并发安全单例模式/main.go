package main

type singleton struct {}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = &singleton{}   // 不是并发安全的
	}
	return instance
}

func main() {
	ins := GetInstance()
	_ = ins
}