package examples

import (
	"fmt"
	"reflect"
)

// structure
type Person struct {
	name string
	age  int
}

func Initialize_Struct_Demo() {
	var pers1 Person

	// Pers1 specification
	pers1.name = "Jack"
	pers1.age = 20

	fmt.Println("Name: ", pers1.name)

	pers2 := Person{"Bob", 20}
	fmt.Println(pers2)

	pers3 := Person{name: "Alice", age: 30}
	fmt.Println(pers3)

	pers4 := Person{name: "Alice"} // age is 0.
	fmt.Println(pers4)

	// "new" is a built-in function of go.
	pers5 := new(Person)
	// it is same with:
	pers6 := Person{}

	// type of pers5 and 6 are also same.

	fmt.Println(pers5, pers6)
}

func Initialize_Variables_Demo() {
	// constants are read-only and un-changeable.
	const ok2 = true

	/////////////////////////
	/// NEW VARIABLES
	/////////////////////////
	// if you don't use any initialized variable, go compiler will give error.
	var b1 bool = true // typed declaration with initial value
	b2 := true         // untyped declaration with initial value. but type is automatic assign on runtime.
	var b3 bool        // typed declaration without initial value
	b4 := true         // untyped declaration with initial value. iki nokta sadece "var" keyword'ünü yazmamamızı sağlıyor.

	// we print them so compiler will not give any error.
	fmt.Println("new variables: ", b1, b2, b3, b4)

	fmt.Println("type of: ", reflect.TypeOf(b2))

	// we define 2 values on same line.
	var firstName, lastName string = "name1", "surname1"

	fmt.Println(firstName, lastName)

	/////////////////////////
	/// NEW VARIABLES SYNTAX (another example)
	/////////////////////////

	// all below are equal:
	var (
		speed1    int
		velocity1 int
	)

	// or:
	var speed2 int
	var velocity2 int

	// or:
	var speed3, velocity3 int

	fmt.Println(speed1, speed2, speed3, velocity1, velocity2, velocity3)
}
