package main

import "fmt"

func main() {
	var d *dog = &dog{
		feet: "four feet",
		Animal:&Animal{
			Name: "dog",
			sex: "xiong",
		},
	}
	d.Eat()
	d.Talk()
}

type Animal struct {
	Name string
	sex  string
}

func (a *Animal) Talk() {
	fmt.Printf("i talk i,m %s\n", a.Name)

}

type dog struct {
	feet string
	*Animal
}

func (d *dog) Eat() {
	fmt.Printf("dog is eat")
}
