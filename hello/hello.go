package main

import "fmt"

func change(x *int, y *int) {
	*x, *y = *y, *x

}
func main() {
	x := 1
	y := 2
	fmt.Println(x, y) // x всё еще равен 5
	change(&x, &y)
	fmt.Println(x, y)
	var buf []int
	buf = append(buf, 9, 10)
	buf[0] = 25
	newbuf := buf[:]
	newbuf[0] = 1
	fmt.Println(buf[0])

}
