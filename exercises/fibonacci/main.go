package main

import (
	"fmt"
	_"strings"
)
var b []byte
func main() {
	var (
		n,sum int
		// str []string
	)
	fmt.Println("put a number")
	fmt.Scanf("%d",&n)
	sum = run(n)
	s := fbn(n)
	fmt.Println()
	for _,v := range s {
		fmt.Printf("%d, ",v)
	}
	fmt.Println()
	for i := 1;i <= n;i++ {
		a := run(i)
		fmt.Printf("%d, ",a)
	}
	fmt.Println()
	fmt.Println("sum =",sum)
	// s := make([]string,len(b))
	// fmt.Println("the string is",b)

}

func fbn(n int) ([]uint64) {
	s := make([]uint64,n)
	s[0] = 1
	s[1] = 1
	for i :=2;i < n;i++ {
		s[i] = s[i-1] + s[i-2]
	}

	return s
}

func run(n int) int{
	var a int
	if n < 3 {
		a = 1
	}else {
		a = run(n-1) + run(n-2)
	}
	return a
}