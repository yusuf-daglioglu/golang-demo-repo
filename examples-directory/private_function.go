package examples

func privateFunc() {
	// this is private function. because it does not start with capital case.
	// can not be called from other packages.
}

func Private_Function_Demo() {
	// this is public function.

	privateFunc()

	// optional semi-colon can be used:
	// privateFunc(); privateFunc()
}
