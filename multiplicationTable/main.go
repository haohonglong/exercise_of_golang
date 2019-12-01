package main

import "fmt"

func main() {
	run(9,9)
}

func run(row int,col int) {
	for i := 1;i <= row;i++ {
		for j :=1;j <=i; j++ {
			fmt.Printf("%d * %d = %d\t",j,i,j*i)
		}
		fmt.Println();
	}
}