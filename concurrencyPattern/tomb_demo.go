package main

import (
	"bufio"
	"io"
	"strings"
	"fmt"
	"gopkg.in/tomb.v2"
	"time"
)

// 利用Tomb包控制协程的生命
// https://blog.labix.org/2011/10/09/death-of-goroutines-under-control

type LineReader struct {
	Ch chan string
	r  *bufio.Reader
}

func (lr *LineReader) loop() {
	i := 0
	for {
		i++
		time.Sleep(500 * time.Millisecond)
		line, err := lr.r.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			close(lr.Ch)
			return
		}
		fmt.Println("loop ", i, string(line))
		lr.Ch <- string(line)
	}
}

func NewLineReader(r io.Reader) *LineReader {
	lr := &LineReader{
		Ch: make(chan string),
		r:  bufio.NewReader(r),
	}
	go lr.loop()
	return lr
}

type LineReaderTomb struct {
	Ch chan string
	r  *bufio.Reader
	t  tomb.Tomb
}

func NewLineReaderTomb(r io.Reader) *LineReaderTomb {
	lr := &LineReaderTomb{
		Ch: make(chan string),
		r:  bufio.NewReader(r),
	}
	lr.t.Go(lr.loop)
	return lr
}

func (lr *LineReaderTomb) loop() error {
	i := 0
	for {
		i++
		time.Sleep(500 * time.Millisecond)
		line, err := lr.r.ReadSlice('\n')
		if err == io.EOF {
			return nil
		}
		if err != nil {
			close(lr.Ch)
			return err
		}
		fmt.Println("loop ", i, string(line))
		select {
		case lr.Ch <- string(line):
		case <-lr.t.Dying():
			close(lr.Ch)
			return nil
		}
	}
}

func (lr *LineReaderTomb) Stop() error {
	lr.t.Kill(nil)
	return lr.t.Wait()
}

func runWithoutTomb() {

	timeEnd := time.After(time.Second * 3)

	lr := NewLineReader(strings.NewReader(strings.Repeat("aaa\n", 10)))

	for {
		select {
		case d := <-lr.Ch:
			fmt.Printf("Ch data:%+v\n", d)
		case <-timeEnd:
			fmt.Printf("Time to end\n")
			close(lr.Ch) // 强制关闭
			return
		}
	}
}

func runWithTomb() {
	timeEnd := time.After(time.Second * 10)
	timeStop := time.After(time.Second * 5)
	lr := NewLineReaderTomb(strings.NewReader(strings.Repeat("aaa\n", 10)))

	for {
		select {
		case d := <-lr.Ch:
			fmt.Printf("Ch data:%+v\n", d)
		case <-timeStop:
			lr.Stop() // 强行结束任务
			return
		case <-timeEnd:
			fmt.Printf("Time to end\n")
			return
		}
	}
}

func main() {
	//runWithoutTomb()
	runWithTomb()
}
