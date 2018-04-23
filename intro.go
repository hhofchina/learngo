package main

//go简单入门介绍代码
import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math"
	"math/rand"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"sort"
	"sync"
	"time"
	"golang.org/x/net/context"

	_ "github.com/go-sql-driver/mysql"

	"database/sql"
	"runtime"
	"reflect"
	"strings"
	"io/ioutil"
)

func test1() {
	m := make(map[string]string)
	m["a"] = "av"
	m["b"] = "bv"
	if m["c"] == "" {
		fmt.Println("not exist m[c]")
	} else {
		fmt.Println("m[c]=", m["c"])
	}
	if cv, ok := m["c"]; ok {
		fmt.Println(cv, ok)
	} else {
		fmt.Println("m[c] not ok")
	}
	fmt.Println(m)
}

func add(args ...int) (tot int) {
	for _, v := range args {
		tot += v
	}
	return tot
}

func makeSeq() func() uint {
	seq := uint(0)
	return func() (r uint) {
		r = seq
		seq += 1
		return
	}
}

func testmakeseq() {
	nextval := makeSeq()
	fmt.Println(nextval())
	fmt.Println(nextval())
	fmt.Println(nextval())
	fmt.Println(nextval())
	fmt.Println(nextval())
}

func fac(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * fac(x-1)
}

func cleanup1() {
	fmt.Println("call cleanup1")
}
func cleanup2() {
	fmt.Println("call cleanup2")
}
func panictest() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a[10])
	//panic("panic...occurred")
}
func testPointer(vPtr *int) {
	*vPtr += 1
}

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Person struct {
	age  uint
	name string
}
type Student struct {
	Person Person
	id     uint
}

type Teacher struct {
	Person
	tid uint
}
type Shape interface {
	area() float64
}

func sumarea(shapes ...Shape) (totalarea float64) {
	for _, s := range shapes {
		totalarea += s.area()
	}
	return
}

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println("sub process:", n, ":", i)
		t := time.Duration(rand.Intn(1000))
		time.Sleep(time.Millisecond * t)
	}
}

func testgoroutine() {
	defer func() {
		fmt.Println("all process exit")
	}()

	fmt.Println("Press any key to start/stop subthread")
	var input string
	//阻止主进程退出
	fmt.Scanln(&input)
	go f(0)
	//同步运行子进程
	fmt.Scanln(&input)
}

func procuder(c chan<- int) {
	for i := 0; ; i++ {
		fmt.Println("write ", i)
		c <- i //通道写入数据
		time.Sleep(time.Second * 1)
	}

}
func consumer(c <-chan int) {
	for {
		i := <-c
		fmt.Println("recive:", i)
	}
}

func testSelect() {
	c1 := make(chan string, 5)
	c2 := make(chan string, 10)
	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()
	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println("C1", msg1)
			case msg2 := <-c2:
				fmt.Println("C2", msg2)
			case <-time.After(time.Second * 1):
				fmt.Println("time out")
				//default:
				//	fmt.Println("nothing ready") //读不到数据立即返回
			}
		}
	}()
	var input string
	fmt.Scanln(&input)
}

func testchannel() {
	defer func() {
		fmt.Println("all process exit")
	}()

	fmt.Println("Press any key to start/stop subthread")
	var input string
	//阻止当前进程退出
	fmt.Scanln(&input)

	var c chan int = make(chan int) //
	go procuder(c)
	go procuder(c)
	go consumer(c)
	fmt.Scanln(&input)
}

type ByName []Person

func (this ByName) Len() int {
	return len(this)
}
func (this ByName) Less(i, j int) bool {
	return this[i].name < this[j].name
}
func (this ByName) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
func testSort() {
	a := []Person{
		{10, "name10"},
		{17, "name17"},
		{12, "name12"},
	}
	fmt.Println("before sort:", a)
	sort.Sort(ByName(a))
	fmt.Println("after sort:", a)

}

func testMutexlock() {
	m := new(sync.Mutex)
	for i := 0; i < 10; i++ {
		go func(i int) {
			m.Lock()
			fmt.Println(i, "start")
			time.Sleep(time.Second)
			fmt.Println(i, "end")
			m.Unlock()
		}(i)
	}
	var input string
	fmt.Scanln(&input)
}

//RPC服务
type RpcServer struct{}

//RPC服务方法
func (this *RpcServer) Addsome(i int64, reply *int64) error {
	*reply = i + 100
	return nil //ok，no error
}
func (this *RpcServer) Ping(ignore int64, reply *int64) error {
	return nil //ok，no error
}
func rpcserver() {
	rpc.Register(new(RpcServer)) //注册RPC服务器
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := ln.Accept() //建立连接
		if err != nil {
			continue //忽略错误
		}
		go rpc.ServeConn(c) //响应服务
	}
}
func rpcclient() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()
	var wg sync.WaitGroup
	var result int64
	b := time.Now()
	err = c.Call("RpcServer.Ping", 1, &result)
	if err == nil {
		fmt.Println("Ping ok")
		for i := 0; i < 10000; i++ {
			wg.Add(1)
			go func(i int, result int64) {
				defer wg.Done()
				err = c.Call("RpcServer.Addsome", int64(i), &result)
				if err != nil {
					fmt.Println("fail", i, err)
				} else {
					fmt.Printf("\rRpcServer.Addsome(%d)=%d", i, result)
				}
			}(i, result)
		}
	} else {
		fmt.Println(err)
	}
	wg.Wait()
	fmt.Println("\nuse time:", time.Since(b))
}

func testRpc() {
	go rpcserver()
	time.Sleep(time.Second * 3)
	go rpcclient()
	var input string
	fmt.Scanln(&input)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path=%q\n", r.URL.Path)
	log.Println(r.URL.Path)

}

func headerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	fmt.Fprintf(w, "Host=%q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr=%q\n", r.RemoteAddr)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q]=%q\n", k, v)
	}
}

type Filter struct {
	http.Handler
}

//http.Hanlder接口实现，也可作为Web入口
func (this *Filter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	headerHandler(w, r)
}
func testWebserver() {

	//注册handler函数
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/header", headerHandler)
	////注册handler
	//http.Handle("/",new(Filter))
	//http.Handle("/header",http.HandlerFunc(headerHandler))//http.HandlerFunc把函数转换成handler
	err := http.ListenAndServe(":8000", nil) //采用：http.HandleFunc设定的handler
	//可以设置服务器端超时时间,第一有效，为何后面无效？
	//http.HandleFunc("/header",http.TimeoutHandler(http.HandlerFunc(headerHandler),
	//	time.Millisecond * 100, "Time out").ServeHTTP)
	//给定handler，响应所有的web请求，可以作为入口handler
	//自定义handler入口
	//err := http.ListenAndServe(":8000",
	//	http.TimeoutHandler(new(Filter), time.Millisecond * 100, "Time out"))
	log.Fatal(err)
}

func testString() {
	s := "hello，中文"
	for i, r := range s {
		fmt.Printf("%d %q %c\n", i, r, r)
	}
	n := 0
	for range s {
		n++
	}
	fmt.Println(n)
	for i, r := range []rune(s) {
		fmt.Printf("%d %q %x\n", i, r, r)
	}
}

type MyInt int

func mustInt(v int) {
	fmt.Printf("recieve:%d", v)
}

func mustMyInt(v MyInt) {
	fmt.Printf("recieve:%d", v)
}

func testConst() {
	const (
		A     = 1
		B int = 9
		C int = 3
	)
	kva := [...]string{A: "aa", B: "bb", C: "CC"}
	fmt.Println(A, B, C, kva[A])
	t := [...]int{6: 1, 2, 3}
	fmt.Println(t)
	const (
		_     MyInt = iota //忽略0
		One
		_      //忽略2
		Three
	)
	//mustInt(Three)  //常量类型不是int，虽然type MyInt int，也会报错！这点哟意思
	//v := 2
	mustMyInt(2)
}

func testSlice() {
	s1 := []int{1, 2, 3}
	fmt.Println(s1, len(s1), cap(s1)) //[1 2 3] 3 3
	s1 = append(s1, 4)
	fmt.Println(s1, len(s1), cap(s1)) //[1 2 3 4] 4 6
}

type Movie struct {
	//go语法规则，首字母大写的成员、函数才能被外部访问，所以，默认json字段名和结构体字段一致，可以通过json声明所需的字段
	Id   int32  `json:"id"`
	Name string `json:"name"`
	Man  bool   `json:"man,omitempty"` //“零”值，空串、0、false不输出
	Sex  int32  `json:"-"` // - 忽略
}

func testJson() {
	var d = []Movie{
		{Id: 1, Name: "women", Man: false},
		{Id: 2, Name: "man", Man: true},
		{Id: 3, Name: "child", Man: false},
	}
	v, err := json.Marshal(d)
	if err != nil {
		log.Fatalf("JSON marshal:%s", err)
	}
	fmt.Printf("%s\n", v)
	v, err = json.MarshalIndent(d, "", "	")
	if err != nil {
		log.Fatalf("JSON marshal:%s", err)
	}
	fmt.Printf("%s\n", v)
	var ud []struct {
		//可以只解读部分字段
		Id  int32
		Man bool
	}
	if err := json.Unmarshal(v, &ud); err != nil {
		log.Fatalf("JSON unmarshal:%s", err)
	}
	fmt.Println(ud)
}
func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
func testHtmlTemplate() {
	const tpl = `<p>A:{{.A}}</p><p>B:{{.B}}</p>`
	t := template.Must(template.New("test").Parse(tpl))
	var data struct {
		A string        //不信任字符，会被HTML转码
		B template.HTML //信任的HTML，不转码
	}
	data.A = "<b>A</b>"
	data.B = "<b>B</b>"
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err) //<p>A:&lt;b&gt;A&lt;/b&gt;</p><p>B:<b>B</b></p>
	}
}

func perr(err error) {
	if err != nil {
		//panic(err.Error())
		panic(err)
	}
}

func testMysql() {
	//datasource 格式：[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	db, err := sql.Open("mysql", "user:pass@tcp(localhost:3306)/laraveldb?charset=utf8")
	perr(err)
	defer db.Close()
	err = db.Ping()
	perr(err)
	fmt.Println("ping ok")
	stmtOut, err := db.Prepare("SELECT count(*) FROM users WHERE id>?")
	perr(err)
	defer stmtOut.Close()
	var cnt int
	err = stmtOut.QueryRow(0).Scan(&cnt)
	perr(err)
	fmt.Println("count:", cnt)
	rows, err := db.Query("SELECT * FROM users")
	perr(err)
	defer rows.Close()
	colnames, err := rows.Columns() //取字段名
	perr(err)
	values := make([]sql.RawBytes, len(colnames))
	//需要interface接口作为rowScan参数
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		perr(err)
		//逐个字段处理
		var v string
		for i, col := range values {
			if col == nil {
				//值为空
				v = "NULL"
			} else {
				v = string(col)
			}
			fmt.Println(colnames[i], "=", v)
		}
	}
}

func testRefSlice(stack []int, i *int) {
	stack = append(stack, *i)
	fmt.Println(*i, stack)
	*i++
	p := *i
	if *i < 10 {
		testRefSlice(stack, i)
		fmt.Println("\t", p, "after", stack)
	}
}

func testRef() {
	i := 1
	var stack []int = nil //slice nil零值相当于：make([]int,0,0)
	var stack2 []int = make([]int, 0, 0)
	fmt.Println("nil==make([]int,0,0):", stack == nil, stack2 != nil, len(stack) == len(stack2), cap(stack) == cap(stack2))
	testRefSlice(stack, &i)
	fmt.Println("stack", stack)
}

func testClosure() {
	for _, dir := range os.TempDir() {
		fmt.Println(dir)
	}
}

func testUsetime() {
	defer usetime("A")   //A on entry: 2016-04-12 16:52:32.5774344 +0800 CST
	defer usetime("B")() //B on entry: 2016-04-12 16:52:27.5605483 +0800 CST
	//B on exit,use time: 5.0168861s
	time.Sleep(5 * time.Second)
}

func usetime(id string) func() {
	start := time.Now()
	fmt.Println(id, "on entry:", start)
	return func() {
		fmt.Println("\n", id, "on exit,use time:", time.Since(start))
	}
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func testRecover() {
	defer func() {
		if p := recover(); p != nil {
			log.Fatal(p)
		}
	}()
	for i := 10; i >= 0; i-- {
		fmt.Println(99 / i)
	}
}

type student struct {
	id   int32
	name string
}

func (s *student) Name() string {
	return s.name
}

func (s *student) SetName(name string) {
	s.name = name
}

func (s student) Showme(p *student) {
	//(s student)值复制传递
	fmt.Println(s.id, s.name, &s, s == *p, &s == p) //1 student1 &{1 student1} true false
}

func (s *student) Pshowme(p *student) {
	//(s *student) 指针地址专题，无需值复制,性能更高
	fmt.Println(s.id, s.name, s, *s == *p, s == p) //1 student1 &{1 student1} true true
}

func testMethodVp() {
	a := student{id: 1, name: "student1"}
	p := &a
	a.Showme(&a)
	p.Showme(&a) //隐式(*p).Pshowname
	p.Pshowme(&a)
	a.Pshowme(&a) //隐式(&a).Pshowname
	a.SetName("aa2")
	fmt.Println(a, a.Name())
}

type IntSet struct{}

func (*IntSet) String() string {
	return ""
}

func testTypeHaveMethod() {

	//var _ =  IntSet{}.String()//编译错误，需要*IntSet的String方法:cannot take the address of IntSet literal
	var s IntSet
	var _ = s.String() //语法糖：s是变量，&s 有String()，正确,
	//String()方法，fmt.Stringer接口有此方法
	var _ fmt.Stringer = &s //正确，&s有String()方法，满足接口
	//var _ fmt.Stringer = s //编译错误，s变量缺少String()方法:IntSet does not implement fmt.Stringer (String method has pointer receiver)

	fmt.Println("Im fine")
}

//基于反射，
func printType(x interface{}) {
	v := reflect.ValueOf(x)
	t := v.Type()
	fmt.Printf("type %s\n", t)

	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("func (%s) %s%s\n", t, t.Method(i).Name, methodType.String())
	}
}

func testPrintType() {
	printType(time.Hour)
	printType(new(strings.Replacer))
	var r string = "abcde"
	var pr = &r //
	printType(r)
	printType(pr)
}

func testHack() {
	//设置环境变量，对比效果set GOMAXPROCS=1 ,2
	//GOMAXPROCS定义goroutine最多允许使用的OS线程数，通常和CPU数一致
	for {
		go fmt.Print(0)
		fmt.Print(1)
	}
}

func Spinner() {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func testSscanf() {
	var u string
	var n float64
	fmt.Sscanf("123.45/cm", "%f/%s", &n, &u) //不支持%.2f
	fmt.Printf("result is:%.2f(%s)", n, u)   //result is:123.45(cm)
}

func testTimeTick() {
	tick := time.Tick(1 * time.Second) //Tick是个定时节拍器，定时向chan发送时间戳，可用于定时输出
	for c := 10; c > 0; c-- {
		fmt.Println(c)
		<-tick //忽略时间戳
	}
}

func testMemSync() {
	var x, y int
	var n sync.WaitGroup
	n.Add(1)
	go func() {
		defer n.Done()
		x = 1
		fmt.Print("y:", y, " ")
	}()
	n.Add(1)
	go func() {
		y = 1
		fmt.Print("x:", x, " ")
		n.Done()
	}()
	n.Wait()
	//var input string
	//fmt.Scan(&input)
}

//函数类型，用于传入抓取函数
type Func func(key string) (interface{}, error)

//结果
type result struct {
	value interface{} //结果值
	err   error
}

//抓取项（任务，相当于FetchEntry）
type entry struct {
	res   result        // 结果
	ready chan struct{} //结果是否就绪chan（抓取完毕）,用于阻塞等待抓取结果
}

//请求
type request struct {
	key      string        //key：url
	response chan<- result //结果chan，存放处理结果
}

//缓存任务项
type Memo struct {
	requests chan request //通道：请求
}

//启动监控实例，响应处理请求任务
func New(f Func) *Memo {
	//memo := &Memo{requests:make(chan request)}//请求通道，无缓存
	memo := &Memo{requests: make(chan request, 50)} //请求通道，有缓存
	go memo.server(f)                               //处理请求
	return memo
}

//向任务项的请求通道发送请求，读取响应
func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response} //向通道发送请求+结果返回通道，等待服务响应
	time.Sleep(time.Millisecond * 1000)     //强制拖延时间
	res := <-response                       //等待结果
	return res.value, res.err
}

//关闭任务
func (memo *Memo) Close() {
	close(memo.requests) //关闭请求chan
}

//响应服务
func (memo *Memo) server(f Func) {
	fmt.Println("begin serve")
	cache := make(map[string]*entry) //抓取项缓存
	for req := range memo.requests {
		//fmt.Printf("\rrequests queue size:%d",len(memo.requests))
		//chan被关闭则循环接受
		//读取通道里的请求，无请求阻塞等待
		e := cache[req.key]
		if e == nil {
			//无缓存
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			fmt.Println("fetching:", req.key)
			go e.call(f, req.key) //异步执行抓取任务
		}
		//异步返回结果，可能会阻塞等待结果
		go e.deliver(req.response)
	}
	fmt.Printf("\rserve chan is closed ") //
}

//调用传入的函数（抓取）
func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready) //通知抓取完成
}

//传送结果
func (e *entry) deliver(response chan<- result) {
	<-e.ready         //阻塞，等待结果，可不必阻塞
	response <- e.res //传递结果
}

//执行http Get请求抓取网页，返回为interface{}，让返回更自由
func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

//带监控实例的并发抓取，监控实例缓存抓取项，
// 如已缓存直接返回结果不抓取，如不存在，缓存此抓取项并启动抓取实例5
//请求、响应通过chan进行传递
func testFetchUrl() {
	urls := []string{"http://www.baidu.com"}
	wg := sync.WaitGroup{}
	m := New(httpGetBody)
	defer m.Close()
	for i := 0; i < 50000; i++ {
		for _, url := range urls {
			wg.Add(1)
			//并发发送抓取请求
			go func(url string, i int) {
				defer wg.Done()
				start := time.Now()
				value, err := m.Get(url)
				if err != nil {
					log.Print(err)
				} else {
					//os.Stdout.Write(value.([]byte))
					fmt.Printf("\r%10d\turl:%s\t%d bytes\ttime:%s", i, url, len(value.([]byte)), time.Since(start))
				}
			}(url, i)
		}
	}
	wg.Wait()
	fmt.Printf("\nFetch end")
}

//测试并发可以承受的goroutine数目
func testGoroutines() {
	var wg sync.WaitGroup
	var mu sync.RWMutex
	shareStop := false
	b := time.Now()
	for i := 1; i <= 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				mu.RLock()
				stop := shareStop
				mu.RUnlock()
				if stop {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
		}()
		fmt.Printf("\r\t started goroutines:%10d", i)
	}
	fmt.Printf("\nStart goroutines use time:%s", time.Since(b))
	//var input string
	//fmt.Printf("\nPress q to  stop all goroutes:\n")
	//fmt.Scan(&input)
	//fmt.Printf("\nWait all goroutes to end")
	c := time.Now()
	mu.Lock()
	shareStop = true
	mu.Unlock()
	fmt.Printf("\nWait lock use time:%s", time.Since(c))
	wg.Wait()
	fmt.Printf("\nWait stop use time:%s", time.Since(c))
}

func testChannel() {
	nchan := make(chan int)      //unbuffered
	sqchan := make(chan int, 10) //buffered
	go func() {
		defer close(nchan) //好习惯，显示告诉读方，再也没有数据了，避免死锁:fatal error: all goroutines are asleep - deadlock!
		for x := 0; x < 20; x++ {
			nchan <- x
		}

	}()

	go func() {
		defer close(sqchan) //好习惯：写方显示告诉读方，再也没有数据了，避免死锁
		for {
			x, ok := <-nchan
			if !ok {
				break
			}
			sqchan <- x * x
		}
	}()

	for x := range sqchan {
		//range方式读取chan
		fmt.Printf("\n%d\t%d", x, len(sqchan)) //显示sqlchan缓存的可读数据量
	}
}

func testContextTimeOut() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}
}

func testMapInit() {
	type tt struct {
		id string
	}
	v := make(map[string]tt)
	//v := map[string]tt{} // 相当于前面一行
	v["234"] = tt{"234"}
	if _, ok := v["123"]; !ok {
		fmt.Print("not found")
	}
}

func main() {
	defer printStack()
	defer usetime("Main")()
	//testFetchUrl()
	//testChannel()
	//testRpc()
	testContextTimeOut()
}
