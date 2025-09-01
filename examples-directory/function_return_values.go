package examples

func Function_Return_Values_Demo() {
	sumV1(1, 2)

	sumV2(3, 4)

	multiReturnV1(5, "6")

	multiReturnV2(7, "8")
}

func sumV1(x int, y int) int {
	return x + y
}

func sumV2(x int, y int) (result int) {
	result = x + y
	return // or: return result
}

func multiReturnV1(x int, y string) (result1 int, result2 string) {
	result1 = x + 1
	result2 = y + "text"
	return
}

func multiReturnV2(x int, y string) (int, string) {
	result1 := x + 1
	result2 := y + "text"
	return result1, result2
}
