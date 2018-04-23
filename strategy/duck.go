package strategy

import (
	"fmt"
)

type QuackBehavior interface {
	Quack()
}

type ArtificialQuack struct{}

func (quack ArtificialQuack) Quack() {
	fmt.Println("this is an artificial quack!")
}

type NaturalQuack struct{}

func (quack NaturalQuack) Quack() {
	fmt.Println("this is a natural quack!")
}

type Duck struct {
	quackBehavior QuackBehavior
}

func (duck *Duck) SetQuackBehavior(behavior QuackBehavior) {
	duck.quackBehavior = behavior
}

func (duck Duck) MakeQuack() {
	duck.quackBehavior.Quack()
}
