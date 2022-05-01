package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	s, t := nextLine(), nextLine()

	fmt.Println(s)
	fmt.Println(t)
}
