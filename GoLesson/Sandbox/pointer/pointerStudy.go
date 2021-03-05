package main

import "fmt"

type bb struct {
	b uint
}

type cc *struct {
	c uint
}

func changeValue(a *uint, b *uint, bs *bb) {

	// 포인터 변수의 포인팅 값을 바꾸기 위해서는 *를 사용해야 한다.
	*a = 12

	// b = 12 // !! compile error

	// 포인터 구조체는 컴파일러가 알아서 포인팅을 해준다.
	bs.b = 12
}

func changePointerStruct(a cc) {

	// *a.c = 12 // compile error
}

func main() {

	// Initialize pointer variable
	var a = new(uint)
	b := new(uint)
	bs := new(bb)

	*a = 4
	*b = 4
	bs.b = 4

	fmt.Println("*bs is ", *bs, bs)
	changeValue(a, b, bs)
	fmt.Println("bs is ", bs, bs)

	var cs cc
	changePointerStruct(cs)
	fmt.Println("cs.c is ", cs.c)
}
