package main
import "fmt"

type person struct {
  name string
  age int
}

func Older(p1 person, p2 person) (person, int) {
  if p1.age > p2.age {
    return p1, p1.age-p2.age
  } else {
    return p2, p2.age-p1.age
  }
}

func main() {
  var tom person
  tom.name, tom.age = "tom", 20

  bob := person{age:14, name:"bob"}
  paul := person{"paul", 23}

  round_one_winner, diff := Older(tom, bob)
  fmt.Printf("%s vs %s. winner is %s. diff is %d.\n", tom.name, bob.name, round_one_winner.name, diff)

  round_two_winner, diff := Older(bob, paul)
  fmt.Printf("%s vs %s. winner is %s. diff is %d.\n", bob.name, paul.name, round_two_winner.name, diff)

}
