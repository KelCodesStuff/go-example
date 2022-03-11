package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	name := "Kel Reid"
	course := "Getting started with Kubernetes"
	module := "4"
	clip := 2
	//courseComplete := false

	fmt.Println("Name and course are set to", name, "and", course, ".")
	fmt.Println("Module and clip are set to", module, "and", clip, ".")
	fmt.Println("Name is of type", reflect.TypeOf(name))
	fmt.Println("Module is of type", reflect.TypeOf(module))

	iModule, err := strconv.Atoi(module)
	if err == nil {
		total := iModule + clip
		fmt.Println("Module plus clip equals", total)
	}

	//Pointers = variables that point to other variables
	fmt.Println("Memory address of *course* variable is", &course)

	var ptr *string = &course
	fmt.Println("Pointing to course variable at address,", ptr, "which holds this value,", *ptr)
}
