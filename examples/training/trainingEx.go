package main

import "fmt"

func main() {
	x := 10
	y := &x
	fmt.Println(x)
	fmt.Println(*y)

}

/*type person struct {
	name string
	age  int
}

func main() {

	var p1 person
	var p2 person

	p1.name = "jd"
	p1.age = 30

	p2.name = "pk"
	p2.age = 36

	printperson(p1)
	printperson(p2)

}

func printperson(pers person) {
	fmt.Println("name:", pers.name)
	fmt.Println("age:", pers.age)
}*/
