package main

import (
	"encoding/json"
	"fmt"
	"github.com/bytedance/sonic"
	"testing"
)

type Book struct {
	Name   string `json:"title"`
	Price  float64
	Tags   []string
	Press  string
	Arthur People
}

type People struct {
	Name    string
	Age     int
	School  string
	Company string
	Title   string
}

var (
	people = People{
		Name:    "mmyang",
		Age:     18,
		School:  "BJTU",
		Company: "BJTU TECH",
		Title:   "Senior Tech Lead",
	}

	book = Book{
		Name:   "Road to Zion",
		Price:  159,
		Tags:   []string{"Go", "Programming", "Zion"},
		Press:  "Tech Press",
		Arthur: people,
	}
)

func TestStdJson(t *testing.T) {
	if bs, err := json.Marshal(book); err != nil {
		fmt.Println(err)
		t.Fail()
	} else {
		fmt.Println(string(bs))

		var book2 Book
		if err := json.Unmarshal(bs, &book2); err != nil {
			fmt.Println(err)
			t.Fail()
		} else {
			fmt.Printf("%+v\n", book2)
		}
	}
}

func TestSonic(t *testing.T) {
	if bs, err := sonic.Marshal(book); err != nil {
		fmt.Println(err)
		t.Fail()
	} else {
		fmt.Println(string(bs))

		var book2 Book
		if err := sonic.Unmarshal(bs, &book2); err != nil {
			fmt.Println(err)
			t.Fail()
		} else {
			fmt.Printf("%+v\n", book2)
		}
	}
}

func BenchmarkStdJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bs, _ := json.Marshal(book)
		var book2 Book
		err := json.Unmarshal(bs, &book2)
		if err != nil {
			return
		}
	}
}

func BenchmarkSonic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bs, _ := sonic.Marshal(book)
		var book2 Book
		err := sonic.Unmarshal(bs, &book2)
		if err != nil {
			return
		}
	}
}

// go test -v .\sonic_speed_test.go -run=TestStdJson -count=1
// go test -v .\sonic_speed_test.go -run=TestSonic -count=1
// go test -v .\sonic_speed_test.go -bench=BenchmarkStdJson -run=none -count=1 -benchmem -benchtime=2s
// go test -v .\sonic_speed_test.go -bench=BenchmarkSonic -run=none -count=1 -benchmem -benchtime=2s