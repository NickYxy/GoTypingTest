package main

import (
	"fmt"
	"math/rand"
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
	go b.delMap(ip)
	return false
}

func (b *Ban) delMap(ip string) {
	timer := time.NewTicker(time.Second * 1)
	select {
	case <-timer.C:
		b.Lock()
		defer b.Unlock()
		delete(b.visitIPs, ip)
		fmt.Printf("我们删除了ip: %s \n", ip)
	}
}

func main() {
	var success int64 = 0
	ban := NewBan()
	wait := sync.WaitGroup{}
	wait.Add(1000 * 100)
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			go func(j int) {
				defer wait.Done()
				waitTime := time.Duration(rand.Int63n(2))
				time.Sleep(time.Second * waitTime)
				ip := fmt.Sprintf("192.168.0.%d", j)
				if !ban.visit(ip) {
					atomic.AddInt64(&success, 1)
				}
			}(j)
		}
	}
	wait.Wait()
	fmt.Printf("success is: %d", success)
}
