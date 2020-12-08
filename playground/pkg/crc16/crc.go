package main

import (
	"fmt"
	"github.com/howeyc/crc16"
	_ "github.com/howeyc/crc16"
)

func main() {
	f := []byte("hello, world")
	r := crc16.ChecksumCCITT(f)
	fmt.Println(r)

	strs := []string{"1", "2", "3", "4", "5", "6", "7"}
	for _, s := range strs {
		fmt.Println(crc16.ChecksumCCITT([]byte(s)), crc16.ChecksumIBM([]byte(s)),
			crc16.ChecksumMBus([]byte(s)), crc16.ChecksumSCSI([]byte(s)))
	}
}