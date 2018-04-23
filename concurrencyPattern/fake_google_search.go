package main

import (
	"fmt"
	"time"
	"math/rand"
)

var (
	Web   = fakeSearch("Web")
	Image = fakeSearch("Image")
	Video = fakeSearch("Video")
	Map   = fakeSearch("Map")
)

type Result struct {
	string
}

type Search func(query string) Result

func NewResult(kind string, query string) Result {
	return Result{fmt.Sprintf("%s :search for %s", kind, query)}
}

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return NewResult(kind, query)
	}
}

// 顺序搜索
func GoogleV1(query string) (results []Result) {
	results = append(results, Web(query))
	results = append(results, Image(query))
	results = append(results, Map(query))
	results = append(results, Video(query))
	return
}

// 并发搜索，等待全部结果返回
func GoogleV2(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Map(query) }()
	go func() { c <- Video(query) }()
	for i := 0; i < 4; i++ {
		results = append(results, <-c)
	}
	return
}

// 限超时时间，忽略超时搜索结果
func GoogleV2_1(query string) (results []Result) {
	c := make(chan Result)
	timeout := time.After(30 * time.Millisecond)

	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Map(query) }()
	go func() { c <- Video(query) }()

	for i := 0; i < 4; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			return
		}
	}
	return
}

// 并发请求多个服务节点，保留最先返回的结果
func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	done := make(chan struct{}, len(replicas))
	defer close(c)

	searchReplica := func(searchTask Search) (err error) {
		// 判断c是否已经关闭
		defer func() {
			// 捕获后续服务向已关闭 chan 写数据的异常
			if p := recover(); p != nil {
				//fmt.Printf("channel closed:%+v\n", fmt.Errorf("panic error: %v", p))
				// demo如何返回捕获的panic，转为error返回
				err = fmt.Errorf("Error channel closed:%+v\n", fmt.Errorf("panic: %v", p))
			}
		}()
		r := searchTask(query)
		select {
		case <-done:
			return
		default:
			c <- r
		}
		return
	}

	for _, searchTask := range replicas {
		go func() {
			if err := searchReplica(searchTask); err != nil {
				fmt.Printf("Search err found:%v\n", err)
			}
		}()
	}

	first := <-c //最先返回的结果保留
	//通知其余的结束
	for i:=0;i<len(replicas)-1;i++ {
		done <- struct{}{}   // done
	}
	return first
}

// 并发请求多个服务节点，保留最先返回的结果
func GoogleFirst(query string) (results []Result) {
	results = append(results, First(query,
		fakeSearch("server1"),
		fakeSearch("server2"),
		fakeSearch("server3"),
		fakeSearch("server4"),
	))
	return
}

// 限超时时间,同类型服务节点保留最先返回的，忽略超时搜索结果
func GoogleV3(query string) (results []Result) {
	c := make(chan Result, 4)
	timeout := time.After(30 * time.Millisecond)

	go func() { c <- First(query, fakeSearch("web1"), fakeSearch("web2"), fakeSearch("web2")) }()
	go func() { c <- First(query, fakeSearch("Image1"), fakeSearch("Image2"), fakeSearch("Image3")) }()
	go func() { c <- First(query, fakeSearch("Video1"), fakeSearch("Video2"), fakeSearch("Video3")) }()
	go func() { c <- First(query, fakeSearch("Map1"), fakeSearch("Map2"), fakeSearch("Map3")) }()

	for i := 0; i < 4; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Printf("timed out wait:%d\n", i)
			return
		}
	}
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()

	//results := GoogleV1("Golang")
	//results := GoogleV2("Golang")
	//results := GoogleV2_1("Golang")
	//results := GoogleFirst("Golang")
	results := GoogleV3("Golang")

	fmt.Printf("RESULTS:\n %d result %v\n", len(results), results)
	fmt.Printf("Search use time:%+v\n", time.Since(start))
}
