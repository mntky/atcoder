package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var ans int
	var mochi[]int
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	for i:=1; i<=n; i++ {
		sc.Scan()
		d, _ := strconv.Atoi(sc.Text())
		mochi = append(mochi, d)
	}

	for i:=1; i<=100; i++ {
		for _, d := range mochi {
			if d == i {
				ans+=1
				break
			}
		}
	}
	fmt.Println(ans)
}
