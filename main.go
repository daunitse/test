package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanln(&a)
	for i := 0; i < a; i++ {
		fmt.Scanln(&b)
		if b%8 == 0 && b/10 > 0 && b/10 < 10 {
			c = c + b
		}
	}
	fmt.Println(c)

}
