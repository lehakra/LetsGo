package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

var isOne bool // variable declared at package level

var isTwo, num, str = true, 10, "Romil" // type can be obmitted when values provided

var (
	toBe   bool       = false
	maxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func needInt(x int) int { return x*10 + 1 }

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y)) // there are no implitit conersions in go
	const True = true                             // no short assignment allowed for constants
	fmt.Println(rand.Intn(10))
	fmt.Println(add(10, 20))
	fmt.Println(split(20))
	fmt.Println(isTwo, num, str)
	fmt.Println(f, True)

	const big = 1<<100 + 1

	fmt.Println(needInt(big))
}

func split(value int) (x, y int) { // naming return values cal help with naming
	x = 10
	y = 20
	return //naked return, bad for long functions ok for small functions like this
}

func add(x int, y int) int {
	// sum := x + y // := is called short assignment and only exists in a function
	return x + y
}
func init() {
	rand.Seed(time.Now().UnixNano())
}
