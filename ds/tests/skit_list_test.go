package tests

import (
	"fmt"
	"github.com/JameyWoo/goto/ds/skit_list"
	"math/rand"
	"testing"
	"time"
)

// int32 的最大最小值
const IntMax = int(^uint(0) >> 1)
const IntMin = ^IntMax

type myInt int

func (a *myInt) Compare(b skit_list.SkipListObj) bool {
	return *a < *b.(*myInt)
}

func (a *myInt) PrintObj() {
	fmt.Print(*a)
}

func TestCreateSkipList(t *testing.T) {
	var obj myInt
	obj = myInt(IntMin)
	s, err := skit_list.CreateSkipList(&obj, 10)
	if s == nil {
		fmt.Print(err)
		t.Errorf("create list failed")
	}
}

// 这个的索引建立是正常的, 数字
func TestOperations(t *testing.T) {
	var minObj, obj myInt
	minObj = myInt(IntMin)
	s, err := skit_list.CreateSkipList(&minObj, 10)
	if s == nil || err != nil {
		fmt.Print(err)
		t.Errorf("create skip list failed")
	}

	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		insertObj := new(myInt)
		*insertObj = myInt(rand.Intn(50))
		res, err := s.Insert(insertObj)
		if res == true {
			fmt.Println("insert obj ", obj, " success")
		} else {
			fmt.Print(err)
			t.Errorf("insert obj %d failed: ", obj)
		}
		//sleep 10ms
		time.Sleep(10000000)
		rand.Seed(time.Now().UnixNano())
		obj = myInt(rand.Intn(50))
		_, err = s.Search(&obj)
		_, err2 := s.RemoveNode(&obj)
		if err == nil && err2 != nil {
			fmt.Print(err)
			t.Errorf("remove obj %d failed: ", obj)
		} else {
			fmt.Printf("remove obj %d success\n", obj)
		}
	}
	fmt.Println("start print the skip list")
	s.Traverse()
}

// 自定义的string类型跳表
type myString struct {
	s string
}

func (a *myString) Compare(b skit_list.SkipListObj) bool {
	return (*a).s < (*b.(*myString)).s
}

func (a *myString) PrintObj() {
	fmt.Print(*a)
}

func TestStringSkipList(t *testing.T) {
	sl, err := skit_list.CreateSkipList(&myString{"   "}, 5)  // bixv
	if err != nil {
		panic(err)
	}
	rand.Seed(time.Now().UnixNano())
	sl.Insert(&myString{"fuck"})
	sl.Insert(&myString{"your"})
	sl.Insert(&myString{"mother"})
	sl.Insert(&myString{"and"})
	sl.Insert(&myString{"your"})
	sl.Insert(&myString{"father"})

	sl.Traverse()
}


// 自定义的string类型跳表
type myStr string

func (a *myStr) Compare(b skit_list.SkipListObj) bool {
	return *a < *b.(*myStr)
}

func (a *myStr) PrintObj() {
	fmt.Print(*a)
}

// 建立索引的时候有bug
// 经常是对整个行都加到了第二个索引上
func TestStringSkipList1(t *testing.T) {
	obj := myStr("000")
	sl, err := skit_list.CreateSkipList(&obj, 10)
	if err != nil {
		panic(err)
	}
	rand.Seed(time.Now().UnixNano())
	objs := []myStr{"fuck", "your", "mom", "and", "your", "fa"}
	for i := 0; i < len(objs); i++ {
		sl.Insert(&objs[i])
	}

	sl.Traverse()
}