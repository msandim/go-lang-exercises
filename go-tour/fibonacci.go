package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	i, n1, n2 := 0, 0, 1 // i is the counter
	
	return func() int {
		
		i++
		
		if i - 1 == 0 {
			return n1
		} else if i -1 == 1 {
			return n2
		}
		
		sum := n1 + n2
		n1 = n2
		n2 = sum
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}
