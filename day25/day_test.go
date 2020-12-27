package day25

import (
	"log"
	"testing"
)

func findLoopSize(publicKey int64) int {
	// i := 0
	// for ; ; i++ {
	var p int64 = 1
	for j := 1; ; j++ {
		p = (p * 7) % 20201227
		log.Println(p*7, p)
		if p == publicKey {
			return j
			// break
		}
	}
	// if p > publicKey {
	// 	log.Panicln(i, p)
	// }
	// }
	// return i
}

func TestOne(t *testing.T) {
	//                        5764801, 17807724
	//                        9717666, 20089533
	loopSize1 := findLoopSize(9717666)
	print(loopSize1)
}
