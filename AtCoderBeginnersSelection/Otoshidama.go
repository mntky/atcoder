package main

import (
	"fmt"
	"os"
)

func main() {
	var n,y int
	fmt.Scanf("%d %d", &n, &y)

	for i:=0; i<=n; i++ {
		for j:=0; j<=n-i; j++ {
			k := (y - (i*10000+j*5000))/1000
			if((i+j+k) == n){
				fmt.Printf("%d %d %d",i,j,k)
				os.Exit(0)
			}
		}
	}
	fmt.Println("-1 -1 -1")
}
