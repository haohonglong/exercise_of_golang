package main

import (
	"fmt"
	"lib"
)

type Test struct {
	Arr [3][3]int
}

func (t *Test) Echo() {
	
    arr := t.Arr
	for i,len := 0,len(arr); i < len; i++ {
		for j := 0; j < len; j++ {
			fmt.Printf("%d ",arr[i][j])
		}
		fmt.Println()
	}
	fmt.Println()

	for i,len := 0,len(arr); i < len; i++ {
		for j := 0; j < i; j++ {
			if i != j {
				lib.Swap(&arr[i][j],&arr[j][i])
			}
		}
	}
	t.Show()

	

	
}

func (t *Test) Show() {
	arr := t.Arr
	for i,len := 0,len(arr); i < len; i++ {
		for j := 0; j < len; j++ {
			fmt.Printf("%d ",arr[i][j])
		}
		fmt.Println()
	}
}

func main() {
	var test Test
	arr := [3][3]int {
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	test.Arr = arr
	test.Echo()
}