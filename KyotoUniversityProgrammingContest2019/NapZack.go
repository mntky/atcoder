package main

import (
	"fmt"
)

func main() {
	var n,m,w,a int
	var sina,j []int
	fmt.Scanf("%d %d %d", &n, &m, &w)
	for i:=0; i<(n*2); i++ {
		fmt.Scanf("%d", &a)
		sina = append(sina, a)
	}
	for k:=0; k<(m*2); k++{
		fmt.Scanf("%d", &a)
		j = append(j, a)
	}
	fmt.Println(sina)
	fmt.Println(j)
}
