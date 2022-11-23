package main

import (
	"design-pattern-go/singleton"
	"time"
)

func main() {
	for i := 0; i < 20; i++ {
		go singleton.GetInstance()
	}
	time.Sleep(30 * time.Second)
}
