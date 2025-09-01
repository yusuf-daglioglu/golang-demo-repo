package examples

import "fmt"

// we pass integer by value, so if we change the value,
// the parent/caller function will not see the update.
func changeIt(deger int) {
	deger = 0
}

// the sender/caller sends the integer as address,
// but changeItViaPointer function use it as pointer.
// if we change the value inside this function,
// the parent/caller function will see the change directly.
func changeItViaPointer(deger *int) {
	*deger = 0
}

func Pointer_Demo() {
	// meaning of pointers and address (&, *) are same as C.

	// we require pointers for sending parameters as reference (not their values).

	i := 1
	fmt.Println(i) // prints: 1

	changeIt(i)
	fmt.Println(i) // prints: 1

	changeItViaPointer(&i)
	fmt.Println(i) // prints: 0

	fmt.Println(&i) // prints: an address

	// on C we can define directly a pointer.
	// but on go we need the value first:

	var val1 int = 99
	var po1 *int
	po1 = &val1

	fmt.Println("value", val1)                                      // prints: 99
	fmt.Println("pointer's value", *po1)                            // prints: 99
	fmt.Println("pointer itself (which is an address)", po1, &val1) // prints: same 2 addresses
	fmt.Println("&po1", &po1)                                       // prints: address different than above

	// pointers can be nil:
	po1 = nil
	// but we can not assign nil to a value:
	// val1 = nil

	////////////
	// multiple star example (pointer of pointer)
	////////////
	var num int = 42

	var ptr1 *int = &num

	// create pointer of pointer
	var ptr2 **int = &ptr1

	// create pointer of pointer
	var ptr3 ***int = &ptr2

	// Değerleri yazdırıyoruz
	fmt.Println("num:", num)      // 42
	fmt.Println("ptr1:", *ptr1)   // 42
	fmt.Println("ptr2:", **ptr2)  // 42
	fmt.Println("ptr3:", ***ptr3) // 42

	***ptr3 = 100
	fmt.Println("num after change:", num) // prints: 100
}
