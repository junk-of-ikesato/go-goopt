package main

import (
	"fmt"
	goopt "github.com/droundy/goopt"
)

func main() {
	var color = goopt.Alternatives([]string{"--color", "--colour"},
		[]string{"default", "red", "green", "blue"},
		"determine the color of the output")
	var repetitions = goopt.Int([]string{"-n", "--repeat"}, 1, "number of repetitions")
	var username = goopt.String([]string{"-u", "--user"}, "User", "name of user")
	var children = goopt.Strings([]string{"--child"}, "name of child", "specify child of user")
	goopt.Description = func() string {
		return "Example program for using the goopt flag library."
	}
	goopt.Version = "1.0"
	goopt.Summary = "goopt demonstration program"
	goopt.Parse(nil)
	defer fmt.Print("\033[0m") // defer resetting the terminal to default colors
	switch *color {
	case "default":
	case "red":
		fmt.Print("\033[31m")
	case "green":
		fmt.Print("\033[32m")
	case "blue":
		fmt.Print("\033[34m")
	default:
		panic("Unrecognized color!") // this should never happen!
	}
	fmt.Println("I have now set the color to", *color, ".")
	for i := 0; i < *repetitions; i++ {
		fmt.Println("Greetings,", *username)
		fmt.Println("You have", *repetitions, "children.")
		for _, child := range *children {
			fmt.Println("I also greet your child, whose name is", child)
		}
	}
}
