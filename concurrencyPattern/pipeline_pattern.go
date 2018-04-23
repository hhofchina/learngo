package main

import (
	"fmt"
	"sync"
)

// 把多个整数参数，推入chan，返回chan，当被chan中数据被消费完，会关闭chan
// 如同pushIn
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// pipeline 传递请求
func pipelineV1() {
	// Set up the pipeline.
	c := gen(2, 3)
	out := sq(c)

	// Consume the output.
	fmt.Println(<-out) // 4
	fmt.Println(<-out) // 9
}

// pipeline 输入输出参数类型相同，可以连接起来
func pipelineV2() {
	for n := range sq(sq(gen(2, 3))) {
		fmt.Println(n) // 16 then 81
	}
}

func mergeWithBlock(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	// chan中数据不被消费，会导致下一个要进入的数据被阻塞
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func runMergeWithBlock() {
	c1 := gen(2)
	c2 := gen(3)

	out := mergeWithBlock(c1, c2)
	fmt.Println(<-out)
	fmt.Println(<-out)
}

func runMergeWithBuffer() {
	c1 := gen(2)
	c2 := gen(3)

	out := mergeWithBuffer(c1, c2)
	fmt.Println(<-out)
	fmt.Println(<-out)
}

func mergeWithBuffer(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	// 增加chan缓存
	out := make(chan int, len(cs)) // enough space for the unread inputs
	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// 支持取消操作的
func mergeWithCancel(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int, len(cs))

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed or it receives a value
	// from done, then output calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			select {
			case out <- n:
			case <-done:
			}
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func runMergeWithCancel() {
	in := gen(2, 3)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(in)
	c2 := sq(in)

	// Consume the first value from output.
	done := make(chan struct{}, 2)
	out := mergeWithCancel(done, c1, c2)
	fmt.Println(<-out) // 4 or 9

	// Tell the remaining senders we're leaving.
	done <- struct{}{}
	done <- struct{}{}
}

func genWithDone(done <-chan struct{},nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
		close(out)
	}()
	return out
}

// 带有done参数的sq
func sqWithDone(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}
		}
	}()
	return out
}

func mergeWithDone(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c or done is closed, then calls
	// wg.Done.
	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func runMergeWithDone()  {
	// Set up a done channel that's shared by the whole pipeline,
	// and close that channel when this pipeline exits, as a signal
	// for all the goroutines we started to exit.
	done := make(chan struct{})
	defer close(done)

	in := genWithDone(done, 2, 3)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := sqWithDone(done, in)
	c2 := sqWithDone(done, in)

	// Consume the first value from output.
	out := mergeWithDone(done, c1, c2)
	for n := range out{ // 顺序不保证和mergeWithDone参数顺序一致
		fmt.Println(n)
	}
}

func main() {
	//pipelineV1()
	//pipelineV2()
	//runMergeWithBlock()
	//runMergeWithBuffer()
	runMergeWithCancel()
	//runMergeWithDone()
}
