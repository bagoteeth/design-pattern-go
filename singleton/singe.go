package singleton

import (
	"fmt"
	"sync"
)

type single struct {
}

var mu sync.Mutex
var singleInstance *single

func GetInstance() *single {
	if singleInstance == nil {
		mu.Lock()
		defer mu.Unlock()
		if singleInstance == nil {
			fmt.Println("creating single")
			singleInstance = &single{}
		} else {
			fmt.Println("single exists")
		}
	} else {
		fmt.Println("single exists")
	}
	return singleInstance
}
