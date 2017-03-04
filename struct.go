package main

import "fmt"

// person exported struct
type person struct {
	name string
	age  int
}

// Older return older person and diff
func older(p1 person, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func main() {
	var tom person
	tom.name, tom.age = "tom", 20

	bob := person{age: 14, name: "bob"}
	paul := person{"paul", 23}

	roundOneWinner, diff := older(tom, bob)
	fmt.Printf("%s vs %s. winner is %s. diff is %d.\n", tom.name, bob.name, roundOneWinner.name, diff)

	roundTwoWinner, diff := older(bob, paul)
	fmt.Printf("%s vs %s. winner is %s. diff is %d.\n", bob.name, paul.name, roundTwoWinner.name, diff)

	roundThreeWinner, diff := older(paul, tom)
	fmt.Printf("%s vs %s. winner is %s. diff is %d\n", paul.name, tom.name, roundThreeWinner.name, diff)

}
