package puppy

import "github.com/Asac2142/dog"

func Bark() string {
	return "Woof"
}

func Barks() string {
	return "Woof woof woof"
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}
