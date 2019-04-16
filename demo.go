package main

import "fmt"

var a, i, j, k int

func main() {
	for a = 100; a < 1000; a++ {
		i = a / 100
		j = (a %100)/10
		k = a % 10

		if a == i*i*i+j*j*j+k*k*k {
			fmt.Println(a)
		}
	}
}
