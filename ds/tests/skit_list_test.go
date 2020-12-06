package tests

import (
	"fmt"
	"github.com/JameyWoo/goto/ds/skit_list"
	"math/rand"
	"testing"
	"time"
)

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