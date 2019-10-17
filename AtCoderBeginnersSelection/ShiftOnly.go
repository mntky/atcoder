package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var num[]int
	var num2[]int
	var ans int
	sc.Split(bufio.ScanWords)
	sc.Scan()

	n, _ := strconv.Atoi(sc.Text())
	for i:=0; i<n; i++ {
		sc.Scan()
		arg, _ := strconv.Atoi(sc.Text())
		num = append(num, arg)
	}
	for{
		for _, arg := range num {
			if(arg%2 == 0) {
				num2 = append(num2, arg/2)
			}else{
				fmt.Println(ans)
				os.Exit(0)
			}
		}
		_ = copy(num, num2)
		num2 = nil
		ans++
	}
}
