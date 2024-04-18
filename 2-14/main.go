package main

import (
	"fmt"
	"sync"
)

func main() {
	//sync.WaitGroup Add(n) Done(n) Wait()
	s := []string{"1", "2", "3", "4", "5", "6"}
	var wg sync.WaitGroup
	for _, item := range s {
		wg.Add(1)
		go SayNumber(item, &wg)
	}
	wg.Wait() //wg没有执行完之前进行阻塞
	fmt.Println("完成")
}

func SayNumber(num string, wg *sync.WaitGroup) {
	fmt.Println("num:" + num)
	wg.Done()
}
