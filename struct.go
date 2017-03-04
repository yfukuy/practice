package main

import "fmt"

// Person exported struct
type Person struct {
	name string
	age  int
}

// Older return older person and diff
func Older(p1 Person, p2 Person) (Person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func main() {
	var tom Person
	tom.name, tom.age = "tom", 20

	bob := Person{age: 14, name: "bob"}
	paul := Person{"paul", 23}

	roundOneWinner, diff := Older(tom, bob)
	fmt.Printf("%s vs %s. winner is %s. diff is %d.\n", tom.name, bob.name, roundOneWinner.name, diff)

	roundTwoWinner, diff := Older(bob, paul)
	fmt.Printf("%s vs %s. winner is %s. diff is %d.\n", bob.name, paul.name, roundTwoWinner.name, diff)

	roundThreeWinner, diff := Older(paul, tom)
	fmt.Printf("%s vs %s. winner is %s. diff is %d\n", paul.name, tom.name, roundThreeWinner.name, diff)

}
