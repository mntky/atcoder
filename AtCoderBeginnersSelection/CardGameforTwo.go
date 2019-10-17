package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"sort"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var n,a,b,c int
	var srt[]int
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ = strconv.Atoi(sc.Text())

	for i:=1; i<=n; i++ {
		sc.Scan()
		d, _ := strconv.Atoi(sc.Text())
		srt = append(srt, d)
	}
	sort.Ints(srt)
	for i:=n-1; i>=0; i-- {
		if(c%2==0) {
			a += srt[i]
		}else{
			b += srt[i]
		}
		c++
	}
	fmt.Println(a-b)
}
