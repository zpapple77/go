package main

import (
	"fmt"
	"sync"
)

type Goods struct {
	v map[string]int
	m sync.Mutex
}

func (g *Goods) Inc(key string, id int) {
	g.m.Lock()
	defer g.m.Unlock()
	fmt.Printf("%d库存数量增加，已加锁")
	g.v[key]++
	fmt.Printf("%d库存数量增加完毕，已解锁")
}

func main() {

}
