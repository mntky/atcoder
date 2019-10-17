package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var n,x,d,max,cnt int
	var a []int
	fmt.Scanf("%d %d", &n, &x)
	sc.Split(bufio.ScanWords)
	for i:=0; i<n; i++ {
		sc.Scan()
		d, _ = strconv.Atoi(sc.Text())
		if(max<d){
			max=d
		}
		a = append(a, d)
	}
	for _,c := range a {
		if(c+x >= max) {
			cnt++
		}
	}
	fmt.Println(cnt)
}
