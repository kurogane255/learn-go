package main

import (
	"fmt"

	"github.com/kurogane255/go-test/mypkg"
)

func main() {
	a := mypkg.Add(1, 10)
	fmt.Println(a)
}
