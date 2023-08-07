package main

import (
	"fmt"
)

func main() {
	a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	for k, v := range a {
		fmt.Printf("%v : %v, ", k, v)
	}
}

/*package main

import (
	"fmt"
)

func main() {

	fmt.Println(hello(2, 3))
	a := []int{1, 2, 3, 4, 5, 6}

	for _, y := range a {
		fmt.Println(y)
	}
	fmt.Println(a)

	a = append(a, 10)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
}

func hello(a int, b string) (result int, txt string) {
	result = a + a

}*/
