package main

import (
	"fmt"
	_"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"io"
)

//抓取内容
func fetch(url string, ch chan <-string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)//write to /dev/null, return bytes count
	//b ,err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s ; %v", url, err)
		return
	}
	//fmt.Printf("%s",b)//output content
	//nbytes=len(b)
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2f %7dbytes %s",secs, nbytes, url)
}



func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) //并发
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs second used", time.Since(start).Seconds())
	var input string
	fmt.Scanln(&input)
}