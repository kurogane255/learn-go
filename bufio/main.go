package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sc.Text())
}
