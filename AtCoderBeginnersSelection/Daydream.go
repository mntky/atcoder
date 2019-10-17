package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	var s,r string
	var t = []string{"maerd","remaerd","esare","resare"}
	var ans[]string
	fmt.Scanf("%s", &s)
	ans = strings.Split(s, "")
	for i:=1; i<=len(ans); i++{
		r += ans[len(ans)-i]
		if(len(r) >= 5 && len(r) <= 7) {
			for _, j := range t {
				if(j == r){
					r = ""
					break
				}
			}
		}else if(len(r) > 7) {
			fmt.Println("NO")
			os.Exit(0)
		}
	}
	if(r==""){
		fmt.Println("YES")
	}else{
		fmt.Println("NO")
		os.Exit(0)
	}
}
