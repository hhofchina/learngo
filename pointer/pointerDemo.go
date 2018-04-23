package main

import (
	"fmt"
)

// ByteSlice 对比指针方法和值方法的异同.
type ByteSlice []byte

// 值方法.
func (slice ByteSlice) AppendV(data []byte) []byte {
	l := len(slice)
	if l + len(data) > cap(slice) {  // reallocate
		// Allocate double what's needed, for future growth.
		newSlice := make([]byte, (l+len(data))*2)
		// The copy function is predeclared and works for any slice type.
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:l+len(data)]
	for i, c := range data {
		slice[l+i] = c
	}
	return slice
}

// 指针方法
func (p *ByteSlice) AppendP(data []byte) {
	slice := *p
	l := len(slice)
	if l + len(data) > cap(slice) {  // reallocate
		newSlice := make([]byte, (l+len(data))*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:l+len(data)]
	for i, c := range data {
		slice[l+i] = c
	}
	*p = slice
}

// 采用Write接口风格，推荐这么做.
func (p *ByteSlice) Write(data []byte) (n int, err error) {
	slice := *p
	l := len(slice)
	if l + len(data) > cap(slice) {  // reallocate
		newSlice := make([]byte, (l+len(data))*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:l+len(data)]
	for i, c := range data {
		slice[l+i] = c
	}
	*p = slice
	return len(data), nil
}

func main()  {
	var b ByteSlice
	b.AppendV([]byte("vvv "))// 值方法调用，b无法保存变化
	(&b).AppendV([]byte("&vvv ")) //通过指针变量的值方法调用，b无法保存变化

	b.AppendP([]byte("ppp ")) // 值调用指针方法，相当于指针方法调用，b保存变化，因ByteSlice类型可以寻址。编译器自动编译为(&b).AppendP
	(&b).AppendP([]byte("&ppp ")) // 指针方法调用，b保存变化

	b.Write([]byte("www ")) // 值调用指针方法，相当于指针方法调用，编译器自动编译为下行
	(&b).Write([]byte("&www ")) //指针方法调用

	fmt.Printf("%+v\n",string(b)) // ppp &ppp www &www

	var b2 ByteSlice
	fmt.Fprintf(&b2, "This hour has %d days", 7)
	fmt.Printf("%+v\n",string(b2)) //This hour has 7 days
}
