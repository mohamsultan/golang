package main

import (
	"fmt"
)

var (
	a = 10
	b = "abc"
)
var mk int =90
func main() {
	// Declare Variable
	var v float32; // static
	var sul_3;
	
	// Initialize
	v = 10; // or (var v int = 10)

	// Another Way
	var i int = 10.0;

	fmt.Println(i, v);

	// format specifier
	fmt.Printf("%v %T", b, a);
	// %v -> value
	// %T -> type
}
