// string formatting test/display

package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	p := point{1, 2}
	fmt.Printf("%v\n", p) // {1,2}

	fmt.Printf("%+v\n", p) // if val struct, display struct field names = {x:1 y:2}

	fmt.Printf("%#v\n", p) // src code location snippet =  main.point{x:1, y:2}

	fmt.Printf("%T\n", p) // print var's type = main.point

	fmt.Printf("%t\n", true) // boolean = true

	fmt.Printf("%d\n", 123) // int formatting = 123

	fmt.Printf("%b\n", 14) // binary = 1110

	fmt.Printf("%c\n", 33) // prints char corresponding int = !

	fmt.Printf("%x\n", 456) // hex encoding = 1c8

	fmt.Printf("%f\n", 78.9) // float decimal formating = 78.900000

	fmt.Printf("%e\n", 123400000.0) // float in scientific notation
	fmt.Printf("%E\n", 123400000.0)

	fmt.Printf("%s\n", "\"string\"") // basic string = "string"

	fmt.Printf("%q\n", "\"string\"") // double quote / escape char strings = "\"string\""

	fmt.Printf("%x\n", "hex this") // hexing = 6865782074686973

	fmt.Printf("%p\n", &p) // pointer notation = 0xc000012090

	fmt.Printf("|%6d|%6d|\n", 12, 345) // width position default right padding
	/*
		|    12|   345|
		|  1.20|  3.45|
		|1.20  |3.45  |
		|   foo|     b|
		|foo   |b     |
	*/

	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")

}
