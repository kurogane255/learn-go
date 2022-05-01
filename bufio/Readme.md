パッケージ bufio

目的
・goで標準入力から数値やテキストを読み込む
・そのためにbufioを使う

https://xn--go-hh0g6u.com/pkg/bufio/

# 定数
MaxScanTokenSize 64kbyte
デフォルトのトークンを保存する最大バッファサイズ

# 

package main

import (ｯｦ使う
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
