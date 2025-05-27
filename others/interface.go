package interfaces

import "fmt"

// Define an interface called Speak
// any type that wants to satisfy this interface must implement a method called Speak that takes no arguments and returns a string
type Speaker interface {
	Speak() string
}

// Define a struct
type Dog struct {
	Name string
}

// Dog implements Speak()
// the (d Dog) is called the receiver, its how Go associates this function qith the Dog type.
func (d Dog) Speak() string {
	return d.Name + " says Woof!"
}

// Define another struct
type Cat struct {
	Name string
}

// Cat implements Speak()
func (c Cat) Speak() string {
	return c.Name + " says Meow!"
}

// A function that accepts anything that implements Speaker
func makeSpeak(s Speaker) {
	fmt.Println(s.Speak())
}

func interfaces() {
	dog := Dog{Name: "Buddy"}
	cat := Cat{Name: "Misty"}

	makeSpeak(dog) // Buddy says Woof!
	makeSpeak(cat) // Misty says Meow!
}
