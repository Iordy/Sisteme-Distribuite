package main

import (
	"fmt"
	"math"
)

func hello() {

	fmt.Println("Hello Go")
}

func name(name string) {

	fmt.Println("Hello,", name)
}

func variables() {

	var name string = "Andrei"
	var age int32 = 21

	gender := "male"

	fmt.Println(name, age, gender)
}

func swap(x int, y int) (int, int) {

	aux := x
	x = y
	y = aux

	return x, y
}

func increment(x int) int {

	y := x + 1

	return y
}

func concatanate(x string, y string) string {

	var ret string = x + y

	return ret
}

func circle(radius int) float64 {

	x := math.Pi

	return x * x * float64(radius)
}

func parse(n int) int {

	sum := 0

	for i := 0; i <= n; i++ {
		sum += i
	}

	return sum

}

func decide(n int) bool {

	if n%2 == 0 {

		return true

	} else {

		return false
	}

	if n > 50 {
		return true
	} else {

		return false

	}
}

func swi(day int) string {

	switch day {
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	case 3:
		return "Wednesday"
	default:
		return "Other"

	}

}

func arr() (int, int) {

	var arr [100]int

	arr2 := [100]int{1, 2, 3, 4, 5}

	arr = arr2

	sum := 0
	max := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if arr[i] > max {
			max = arr[i]
		}
	}

	fmt.Println(arr)

	return sum, max

}

func recursion(n int) int {

	if n == 1 {
		return 1
	} else {
		return recursion(n-1) * n
	}

}

func fib(n int) int {

	if n == 1 {
		return 1
	} else if n == 0 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}

}

func slice() []string {

	s := make([]string, 5)

	s = append(s, "n")
	
	//Cum creste capacitatea cand depasim capacitatea anterioara?

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	s[3] = "d"
	s[4] = "e"

	s = append(s[:2] , s[3:]...)

	fmt.Println((s))

	return s
}
