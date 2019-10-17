package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var n,a,b,ans,sum int
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ = strconv.Atoi(sc.Text())
	sc.Scan()
	a, _ = strconv.Atoi(sc.Text())
	sc.Scan()
	b, _ = strconv.Atoi(sc.Text())

	for i:=1; i<=n; i++ {
		sum = i%10
		sum += (i/10)%10
		sum += (i/100)%10
		sum += (i/1000)%10
		sum += (i/10000)%10

		if(sum>=a && sum<=b) {
			ans+=i
		}
	}
	fmt.Println(ans)
}
