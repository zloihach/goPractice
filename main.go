package main

import "fmt"

//type Point struct {
//	X int
//	Y int
//	S string
//}

//func (p *Point) zalupa() string {
//	return fmt.Sprintf("X: %v, Y: %v, S: %v", p.X, p.Y, p.S)
//}

//callback
//func doSomething(callback func(int, int) int, s string) int {
//	result := callback(5, 2)
//	fmt.Println(s)
//	return result
//}

//zamykaniya
//func totalPrice(initPrice int) func(int) int {
//	sum := initPrice
//	return func(x int) int {
//		sum += x
//		return sum
//	}
//}

type Point struct {
	X, Y int
}

func (p Point) movePoint(dx, dy int) Point {
	p.X += dx
	p.Y += dy
	return p
}

func (p *Point) movePointPtr(dx, dy int) {
	p.X += dx
	p.Y += dy
}

func main() {
	p := Point{1, 2}
	fmt.Println(p)
	fmt.Println(p.movePoint(3, 4))
	fmt.Println(p)
	p.movePointPtr(3, 4)
	fmt.Println(p)
	p.movePointPtr(4, 10)
	fmt.Println(p)

	v := &p
	p.movePointPtr(4, 10)
	fmt.Println(*v)
	//orderPrice := totalPrice(1)
	//fmt.Println(orderPrice(1))
	//fmt.Println(orderPrice(1))
	//fmt.Println(orderPrice(1))
	//fmt.Println(orderPrice(1))
	//fmt.Println(orderPrice(1))
	//fmt.Println(orderPrice(1))
	//callback
	//sumCallback := func(n1, n2 int) int {
	//	return n1 + n2
	//}
	//
	//result := doSomething(sumCallback, "hello")
	//fmt.Println(result)
	//
	//multiplyCallback := func(n1, n2 int) int {
	//	return n1 * n2
	//}
	//result = doSomething(multiplyCallback, "world")
	//fmt.Println(result)

	//slice3
	//var snil []int
	//fmt.Println(snil, len(snil), cap(snil))
	//if snil == nil {
	//	fmt.Println("nil")
	//}
	//slice 3
	//s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//t := s[:6]
	//fmt.Println(t)
	//tt := s[6:]
	//fmt.Println(tt)
	//
	//ttt := s[0:]
	//fmt.Println(ttt)
	//
	//tttt := s[:]
	//fmt.Println(tttt)
	//slice2
	//animalsArray := [4]string{"dog", "cat", "bird", "fish"}
	//fmt.Println(animalsArray)
	//
	//var a []string = animalsArray[0:3]
	//fmt.Println(a)
	//
	//a[0] = "zalupa"
	//fmt.Println(a)

	//slice
	//letters := []string{"a", "b", "c", "d", "e"}
	//letters[0] = "z"
	//letters = append(letters, "f")
	//fmt.Println(letters)
	//
	//createSlice := make([]string, 3)
	//fmt.Println(fmt.Sprintf("len: %v, cap: %v", len(createSlice), cap(createSlice)))
	//createSlice[0] = "a"
	//createSlice[1] = "b"
	//createSlice[2] = "c"
	//
	//fmt.Println(len(createSlice))
	//fmt.Println(cap(createSlice))
	//createSlice = append(createSlice, "d")
	//fmt.Println(fmt.Sprintf("len: %v, cap: %v", len(createSlice), cap(createSlice)))
	//fmt.Println(createSlice)
	//createSlice = append(createSlice, "d")
	//fmt.Println(fmt.Sprintf("len: %v, cap: %v", len(createSlice), cap(createSlice)))
	//arrays
	//var a [2]string
	//a[0] = "Hello"
	//a[1] = "World"
	//fmt.Println(a[0], a[1])
	//fmt.Println(a)
	//
	//numbers := [...]int{1,2,3}
	//numbers[2] = 4
	//

	////LIFO-Last in first out: LIFO
	////defer fmt.Println("1")
	////pointers()
	//structs()
	//
	//p1 := Point{1, 2, "hello"}
	//fmt.Println(p1)
	//p2 := &p1
	//fmt.Println(p2.zalupa())
}

//func structs() {
//	p1 := Point{1, 2, "hello"}
//	fmt.Println(p1)
//	fmt.Println(p1.X)
//	fmt.Println(p1.Y)
//	fmt.Println(p1.S)
//	p2 := Point{X: 123}
//	fmt.Println(p2)
//
//	g := &p1
//	fmt.Println(g)
//	fmt.Println(g.X)
//	fmt.Println(*g)
//
//	fmt.Println(p1.zalupa())
//}

//func pointers() {
//	a := "hello world"
//	b := 42
//	fmt.Println(a, b)
//
//	p := &a
//	fmt.Println(p)
//	fmt.Println(*p)
//	*p = "oh my"
//	fmt.Println(a)
//
//	g := &b
//	*g = *g / 2
//	fmt.Println(b)
//
//}
