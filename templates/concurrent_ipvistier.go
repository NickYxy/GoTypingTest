package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Ban struct {
	visitIPs map[string]time.Time
	sync.RWMutex
}

func NewBan() *Ban {
	return &Ban{visitIPs: make(map[string]time.Time)}
}

func (b *Ban) visit(ip string) bool {
	b.Lock()
	defer b.Unlock()
	if _, ok := b.visitIPs[ip]; ok {
		return true
	}
	b.visitIPs[ip] = time.Now()
	return false
}

func main() {
	var success int64 = 0
	ban := NewBan()
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			go func(j int) {
				ip := fmt.Sprintf("192.168.0.%d", j)
				if !ban.visit(ip) {
					atomic.AddInt64(&success, 1)
				}
			}(j)
		}
	}

	fmt.Printf("success is: %d", success)
}
