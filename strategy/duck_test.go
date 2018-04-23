package strategy

import (
	"testing"
)

func Test(t *testing.T) {
	duck := Duck{}

	artificial := ArtificialQuack{}
	natural := NaturalQuack{}

	duck.SetQuackBehavior(artificial)
	duck.MakeQuack()

	duck.SetQuackBehavior(natural)
	duck.MakeQuack()
}
