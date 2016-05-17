package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func genNum(n int) string {
	r := 9
	for i := 1; i < n; i++ {
		r = r*10 + 9
	}
	pad := fmt.Sprintf("%%0%ds", n)
	return fmt.Sprintf(pad, strconv.Itoa(rnd.Intn(r)))
}

func cls() {
	print("\033[H\033[2J")
}

func main() {
	var ex, inp string
	n := 3
	for {
		cls()
		ex = genNum(n)
		fmt.Println(ex)
		time.Sleep(time.Duration(n/2) * time.Second)
		cls()
		fmt.Scanln(&inp)
		cls()
		if inp == ex {
			n++
		} else {
			if n > 3 {
				n--
			}
		}
	}
}
