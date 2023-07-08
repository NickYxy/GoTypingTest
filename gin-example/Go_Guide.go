package main

import (
	"archive/zip"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Packet struct {
	PacketType    uint8
	PacketVersion uint8
	Data          *Data
}

type Data struct {
	Stat uint8
	Len  uint8
	Buf  [8]byte
}

var validate *validator.Validate

func validateVariable() {
	myEmail := "abc@tencent.com"
	errs := validate.Var(myEmail, "required, email")
	if errs != nil {
		fmt.Println(errs)
		return
	}
}
func (p *Packet) UnmarshalBinary(b []byte) error {
	if len(b) < 2 {
		return io.EOF
	}

	p.PacketType = b[0]
	p.PacketVersion = b[1]

	if len(b) > 2 {
		p.Data = new(Data)
	}

	return nil
}

func foo(c chan int) {
	defer close(c)
	_, err := parse(1, []byte("haha"))

	if err != nil {
		c <- 0
		return
	}
	c <- 1
}

func parse(lenControlByUser int, data []byte) ([]byte, error) {
	size := lenControlByUser
	if size > 64*1024*1024 {
		return nil, errors.New("value too large")
	}

	buffer := make([]byte, size)
	copy(buffer, data)
	return buffer, nil
}

func unZipGood(f string) bool {
	r, err := zip.OpenReader(f)
	if err != nil {
		fmt.Println("read zip file failed...")
		return false
	}

	for _, f := range r.File {
		if !strings.Contains(f.Name, "..") {
			p, _ := filepath.Abs(f.Name)
			err := os.WriteFile(p, []byte("present"), 0640)
			if err != nil {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

// good
func checkIllegal(cmdName string) bool {
	if strings.Contains(cmdName, "&") || strings.Contains(cmdName, "|") || strings.Contains(cmdName, ";") ||
		strings.Contains(cmdName, "$") || strings.Contains(cmdName, "'") || strings.Contains(cmdName, "`") ||
		strings.Contains(cmdName, "(") || strings.Contains(cmdName, ")") || strings.Contains(cmdName, "\"") {
		return true
	}
	return false
}

func doAuthReq(authReq *http.Request) *http.Response {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	}
	client := &http.Client{Transport: tr}
	res, _ := client.Do(authReq)

	return res
}

func server1() {
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		user := r.Form.Get("user")
		pw := r.Form.Get("password")

		log.Printf("Registering new user %s with %s.\n", user, pw)
	})

	http.ListenAndServe("80", nil)
}
func main() {
	packet := new(Packet)
	data := make([]byte, 2)

	if err := packet.UnmarshalBinary(data); err != nil {
		fmt.Println("Failed to unmarshall packet")
		return
	}

	if packet.Data == nil {
		return
	}

	fmt.Printf("Stat: %v\n", packet.Data.Stat)

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("strict-transport-security", "max-age=63072000; includeSubDomains")
		w.Write([]byte("This is an example server.\n"))
	})

	log.Fatal(http.ListenAndServeTLS("443", "yourCert.pem", "yourKey.pem", nil))
}
