package xaes

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestXAes(t *testing.T) {
	//XAes()
	//fmt.Println(Rand32())
	//fmt.Println(len(Rand32()))
	//resp, err := decrypt("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", []byte("Di3cY&_0&7%A1#8-#C839M0-94221X2J"))
	//fmt.Println(resp)
	//fmt.Println(err)
	//Xaes2()

	cache := make(map[int]int, 0)
	for i := 0; i < 500000; i++ {
		cache[i] = i
	}
	start := time.Now().UnixMicro()
	for j := 0; j < 100; j++ {
		num := rand.Intn(5000)
		fmt.Println(cache[num])
	}
	end := time.Now().UnixMicro()
	fmt.Println(end - start)
}
