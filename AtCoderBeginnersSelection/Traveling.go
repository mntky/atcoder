package main

import (
	"fmt"
)

func main() {
	var n,t,x,y int
	f := true
	fmt.Scanf("%d", &n)

	for i:=0; i<n; i++ {
		fmt.Scanf("%d %d %d", &t, &x, &y)
		if(f==true){
			if(t < x+y){
				f = false
			}else if(t%2==1 && (x+y)%2==1){
				continue
			}else if(t%2==0 && (x+y)%2==0){
				continue
			}else{
				f = false
			}
		}
	}
	if(f == true) {
		fmt.Println("Yes")
	}else {
		fmt.Println("No")
	}
}
