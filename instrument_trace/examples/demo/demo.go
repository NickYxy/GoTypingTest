package main

import "github.com/NickYxy/GoTypingTest/instrument_trace"

func foo() {
	defer instrument_trace.Trace()()
	bar()
}

func bar() {
	defer instrument_trace.Trace()()

}

func main() {
	defer instrument_trace.Trace()()
	foo()
}
