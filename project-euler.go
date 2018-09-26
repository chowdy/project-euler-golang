package main

import (
	"fmt"
	"os"
	"plugin"
)

func main() {

	// First command line argument is the problem number
	var problem = os.Args[1]

	// Prepend the problem number with zeroes if necessary
	for i := len(problem); i < 3; i++ {
		problem = "0" + problem
	}

	// Open the built plugin corresopnding to the command line arg
	solution, err := plugin.Open("solution" + problem + ".so")
	if err != nil {
		panic(err)
	}

	// Get the solution from the plugin we opened
	solutionFunc, err := solution.Lookup("Solution" + problem)
	if err != nil {
		panic(err)
	}

	// Run the solution and print the result
	answer := solutionFunc.(func() string)()
	fmt.Println(answer)
}
