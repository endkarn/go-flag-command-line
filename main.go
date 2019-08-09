package main

import (
	"flag"
	"fmt"
)

func main() {
	profile := flag.String("profile", "default", "your profile")
	weight := flag.Int("weight", 0, "your weight")

	flag.Parse()

	fmt.Printf("Hello, you are starting GO with %s profile. \n", *profile)

	if *weight != 0 {
		fmt.Printf("Your weight is %d kilogram, Let's do some fittest \n", *weight)
	} else {
		fmt.Printf("You may forget about setting weight before start.\n")
	}

	fmt.Println("tail:", flag.Args())
}
