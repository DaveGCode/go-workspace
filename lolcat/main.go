package main

import (
	"fmt"
	"math"
	"strings"

	"syreclabs.com/go/faker"
)

// RGB color gen
func rgb(i int) (int, int, int) {
	var f = 0.1
	return int(math.Sin(f*float64(i)+0)*127 + 128),
		int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128),
		int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128)
}

func main() {
	var phrases []string

	// loop fake text info into phrases string array
	for i := 1; i < 3; i++ {
		phrases = append(phrases, faker.Hacker().Phrases()...)
	}

	// conc
	output := strings.Join(phrases[:], "; ")

	// for each char in output, set rgb color
	for j := 0; j < len(output); j++ {
		r, g, b := rgb(j)
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, output[j])
	}
	fmt.Println()
}
