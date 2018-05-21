package main

import(
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	WalkAux(t, ch)
	close(ch)
}

func WalkAux(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		WalkAux(t.Left, ch)
	}
	
	ch <- t.Value
	
	if t.Right != nil {	
		WalkAux(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	// Create channels to use the Walk function:
	ch1, ch2 := make(chan int), make(chan int)
	
	// Walk the trees:
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	
	for {
		t1Value, t1Closed := <- ch1
		t2Value, t2Closed := <- ch2
		
		// If both channels got terminated at the same time, then the sequence has the same size:
		if (t1Closed && t2Closed) {
			return t1Value == t2Value
		}
		
		onlyOneChannelClosed := t1Closed != t2Closed
		
		// If one of the channels got closed, or the values are different:
		if onlyOneChannelClosed || (t1Value != t2Value) {
			return false
		}
	}
}

func main() {
	// Create channel to test Walk:
	ch := make(chan int)
	
	// Walk the tree:
	go Walk(tree.New(1), ch)
	
	for value := range ch {
		fmt.Println(value)
	}
	
	if Same(tree.New(1), tree.New(1)) {
		fmt.Println("TEST 1 PASSED")
	} else {
		fmt.Println("TEST 1 FAILED")
	}
	
	if Same(tree.New(1), tree.New(2)) {
		fmt.Println("TEST 2 FAILED")
	} else {
		fmt.Println("TEST 2 PASSED")
	}
}
