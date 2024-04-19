package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"strings"
)

type CheckOut func(int, int) int

// GetTotal 闭包函数
func GetTotal(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

type GenerateRandom func() int

func RandomSum() GenerateRandom {
	a, b := rand.Intn(10), rand.Intn(20)
	return func() int {
		a, b = b, a+b
		return a
	}
}
func (g GenerateRandom) Read(p []byte) (n int, err error) {
	next := g()
	if next > 21 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func PrintResult(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	var checkout CheckOut = func(a, b int) int {
		return a + b
	}
	println(checkout(1, 2))

	total := GetTotal(68)
	sum := total(100)
	println(sum)

	r := RandomSum()
	PrintResult(r)
}
