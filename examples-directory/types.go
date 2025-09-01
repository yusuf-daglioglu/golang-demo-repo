package examples

import "fmt"

func Types_Demo() {
	// integer types
	var i int // platform depended (32 or 64 bit machine has different max and min size)
	// all below integers have fixed size on all platforms.
	var i8 int8   // 8 bit (max min: -128 to 127)
	var i16 int16 // 16 bit (max min: -32768 to 32767)
	var i32 int32
	var i64 int64

	// Unsigned Integers
	// all above integers may have "u" as prefix.
	var var99 uint

	// float types
	var f32 float32
	var f64 float64

	// bool type
	var b bool

	// string types
	var s string
	var r rune  // type of numeric type
	var by byte // type of numeric type

	fmt.Println(
		i, i8, i16, i32, i64,
		var99,
		f32, f64,
		b, s, r, by,
	)

	/////////////////////////
	/// ARRAYS
	/////////////////////////
	// arrays are always fixed size.

	arr1 := [3]int{1, 2, 3}

	// in below line 3 (the length) is automatically set by compiler:
	arr2 := [...]int{1, 2, 3}

	// we can set all or some variables:
	arr3 := [5]int{1, 2}          // partially set
	arr4 := [5]int{1, 2, 3, 4, 5} // fully set

	// set only second and third variables:
	arr5 := [5]int{1: 10, 2: 40}

	// set third element:
	arr5[2] = 50

	fmt.Println(arr1, arr2, arr3, arr4, arr5)

	/////////////////////////
	/// SLICE
	/////////////////////////

	// slice are similar with arrays but they are resizable
	// and they don't need length when initialize on compile time.
	slice1 := []int{1, 2, 3}

	fmt.Println(slice1)
}
