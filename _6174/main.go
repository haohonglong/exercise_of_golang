package main

import (
	"fmt"
	_"os"
)
func main() {
	var num int 
	fmt.Println("please input a number that is four numbers")
	fmt.Scanf("%d",&num);
	n := len(string(num))
	fmt.Println("length",n)

}

func run(str int) int {
	if 6174 == str {
		return 6174
	}else{

	}
	return 0
	
}

func order(n []int,literal string) {
	len := len(n)
	for i := 0;i < len;i++ {
		for j := i+1;i < len;j++ {
			if "ASC" == literal {
					if n[i] < n[j] {
						swap(&n[i],&n[j])
					}
			}else if "DESC" == literal {
					if n[i] > n[j] {
						swap(&n[i],&n[j])
					}
			}
		}
	}
}

func swap(a,b *int) {
	var c int
	c = *a
	*a = *b
	*b = c
}