Interfaces in Go are types that specify a set of method signatrues that have no implementation..
it defines what methods a type must have, but not how they are implemented....
Any type that implements those emthods implicitly satisfies the interface, no explicit declaration needed
eg....
// any type with a speak() string method satisifies the Speaker interface.
type Speaker interface{
Speak() string
}
Advantages of interfaces
-- polymorphism --> writing functions that can accept any type that satisfies interface
-- Flexibility and decoupling
-- code is more testable and easier to extend.

In Go, no explicit declaration of 'implementing' an interface is needed. its implicit if a type has all required methods.

Lets take Java as an example, when implementing an interface, you must decalre class Dog implments Speaker, in Go this is not needed, as long as class Dog, has all methods that were defined in the interface, then all is well,, as long as class Dog, implements all the methods that were in the interface.
-- There is no implements keyword in Go.

--> Marshalling -- Converting a Go struct into JSON or data suitable for transmission. (Same as serialization)
--> Unmarshalling (deserialization)-- converting transmitted data back into in memory data structure (Go Struct)
