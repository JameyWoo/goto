package DataStruct

import (
"fmt"
"testing"
"time"
)

func TestTime(t *testing.T) {
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(5 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	// 使得timer2强制退出计时, 但上面的go func里的计时无效
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(10 * time.Second)
}

