package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"sync/atomic"
	"time"
)

var totalQuery int32

func handler() {
	atomic.AddInt32(&totalQuery, 1)
	time.Sleep(50 * time.Millisecond)
}

func callHandler() {
	limiter := rate.NewLimiter(rate.Every(100*time.Millisecond), 1)
	//for {
	//	err := limiter.WaitN(context.Background(), 1)
	//	if err != nil {
	//		return
	//	}
	//	handler()
	//}
	for {
		reserve := limiter.ReserveN(time.Now(), 1)
		time.Sleep(reserve.Delay())
		handler()
	}
}

func main() {
	go callHandler()
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Printf("过去1秒钟接口调用了%d次\n", atomic.LoadInt32(&totalQuery))
		atomic.StoreInt32(&totalQuery, 0)
	}
}
