package main

import "fmt"

type person struct {
	age int
}

func (p person) NewPerson(initialAge int) person {
	if initialAge < 0 {
		fmt.Println("Age is not valid, setting age to 0.")
		p.age = 0
	}

	return p
}

func (p person) amIOld() {
	switch {
	case p.age < 13:
		fmt.Println("You are young.")
	case p.age < 18:
		fmt.Println("You are a teenager.")
	default:
		fmt.Println("You are old.")
	}
}

func (p person) yearPasses() person { // why is this returning a person?
	p.age++

	return p
}

func main() {
	var T, age int

	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&age)
		p := person{age: age} // this two lines are weird: just delete this line
		p = p.NewPerson(age)  // and make NewPerson a function instead of a method.
		p.amIOld()
		for j := 0; j < 3; j++ {
			p = p.yearPasses()
		}
		p.amIOld()
		fmt.Println()
	}
}
