package main

import (
	"github.com/sirupsen/logrus"
	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	db, err := leveldb.OpenFile("db0", nil)
	if err != nil {
		logrus.Error(err)
	}
	defer db.Close()
	logrus.Info("create db0")

	db.Put([]byte("bar3"), []byte("foo"), nil)
	db.Put([]byte("hello3"), []byte("world"), nil)

	value, err := db.Get([]byte("shit"), nil)
	if err != nil {
		// Error 是不会停止程序的
		logrus.Error(err)
		// 而Fatal 会停止程序
		//logrus.Fatal(err)
	}
	// 这些都是之前对这个数据库进行的操作. 因为可以持久化, 所以在这里可以显示出来
	value, err = db.Get([]byte("hello"), nil)
	logrus.Info(string(value))
	value, err = db.Get([]byte("bar2"), nil)
	logrus.Info(string(value))
	value, err = db.Get([]byte("hello2"), nil)
	logrus.Info(string(value))
}
