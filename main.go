package main

import "fmt"

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)
}

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {

	fmt.Println("pointer value:", iptr)
	fmt.Println("pointer dereferenced:", *iptr)
	fmt.Println("pointer address:", &iptr)

	*iptr = 0
}
