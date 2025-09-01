package main // this should be always main

/*
multiline
comment
*/

import (

	// contains functions for "ForMatting Text".
	// this is a default package so we don't have to add to go.mod file.
	"fmt"

	// both below are same:
	"golang-demo-module/examples-directory"
	// alias is optional:
	examples_directory_alias "golang-demo-module/examples-directory"
)

func main() {
	fmt.Println("hello main")

	// All below functions are independent.

	examples.Function_Return_Values_Demo()

	// calling functions via:
	// - it's real package name
	// - it's alias
	examples.Initialize_Variables_Demo()
	examples_directory_alias.Initialize_Variables_Demo()

	examples.Parenthesis_Demo()

	examples.Pointer_Demo()

	examples.Private_Function_Demo()

	examples.Receiver_Function_Demo()

	examples.Semi_Colon_Demo()

	examples.Types_Demo()
}
