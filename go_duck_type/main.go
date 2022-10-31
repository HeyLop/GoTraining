package main

import "fmt"

type quackAnimal interface {
	quack()
}
type duck struct{}

func (d duck) quack() {
	fmt.Println("duck quack ga-ga")
}

type cat struct{}

func (c cat) quack() {
	fmt.Println("cat quack mio-mio")
}

type dog struct{}

func (d dog) quack() {
	fmt.Println("dog quack wang-wang")
}

func AnimalQuack(a quackAnimal) {
	a.quack()
}

func main() {
	animals := []quackAnimal{new(duck), new(cat), new(dog)}
	for _, i := range animals {
		i.quack()
		AnimalQuack(i)
	}

}
