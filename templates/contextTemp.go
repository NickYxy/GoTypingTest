package main

import (
	"context"
	"fmt"
	"time"
)

func step1(ctx context.Context) context.Context {
	child := context.WithValue(ctx, "name", "mmyang")
	return child
}

func step2(ctx context.Context) context.Context {
	child := context.WithValue(ctx, "age", 18)
	return child
}

func step3(ctx context.Context) {
	fmt.Printf("name:%s\n", ctx.Value("name"))
	fmt.Printf("age:%d\n", ctx.Value("age"))
}

func f1() {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*1000)
	defer cancel()
	t1 := time.Now()

	time.Sleep(time.Millisecond * 500)

	ctx2, cancel2 := context.WithTimeout(ctx, time.Millisecond*100)
	defer cancel2()
	t2 := time.Now()
	select {
	case <-ctx2.Done():
		err := ctx2.Err()
		t3 := time.Now()
		fmt.Println(t3.Sub(t1).Milliseconds(), t3.Sub(t2).Milliseconds())
		fmt.Println(err)
	}
}

func f2() {
	ctx, cancel := context.WithCancel(context.TODO())
	t0 := time.Now()
	go func() {
		time.Sleep(time.Millisecond * 100)
		cancel()
	}()

	select {
	case <-ctx.Done():
		t3 := time.Now()
		fmt.Println(t3.Sub(t0).Milliseconds())
		err := ctx.Err()
		fmt.Println(err)
	}
}
func main() {
	grandpa := context.TODO()
	parent := step1(grandpa)
	grandson := step2(parent)
	step3(grandson)
	f1()
	f2()
}
