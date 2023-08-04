package main

import "fmt"

func Trace(name string) func() {
	fmt.Println("enter:", name)
	return func() {
		fmt.Println("exit:", name)
	}
}

func foo() {
	defer Trace("foo")()
	bar()
}

func bar() {
	defer Trace("bar")()
}

func main() {
	defer Trace("main")()
	foo()
}
