package main

import (
	"bufio"
	"fmt"
	"strconv"
)

var sc = bufio.newScanner(os.stdin)

func nextInt() int {
	sc.scan()
	i, e := strconv.Atoi(sc.text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	sc.split(bufio.ScanWords)
	n := nextInt()
	x := 0
	for i := 0; i < n; i++ {
		x += nextInt()
	}
	fmt.Println(x)
}
