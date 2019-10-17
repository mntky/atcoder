package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var a,b,c,sum,ans int
	sc.Split(bufio.ScanWords)
	sc.Scan()
	a, _ = strconv.Atoi(sc.Text())
	sc.Scan()
	b, _ = strconv.Atoi(sc.Text())
	sc.Scan()
	c, _ = strconv.Atoi(sc.Text())
	sc.Scan()
	sum, _ = strconv.Atoi(sc.Text())

	for i:=0; i<=a; i++ {
		for j:=0; j<=b; j++ {
			for k:=0; k<=c; k++ {
				if(i*500+j*100+k*50 == sum) {
					ans++
				}
			}
		}
	}
	fmt.Println(ans)
}
