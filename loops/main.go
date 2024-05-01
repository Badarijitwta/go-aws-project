package main

import (
	"fmt"
	"slices"
)

func main() {
	//Here the arrays are more efficient as we are assigning the memory that we need
	animals := [3]string{}
	animals[0] = "cat"
	animals[1] = "dog"
	animals[2] = "cat"
	//another way to initialize aarray which is called Slicing.. where memory isnt declared initially.
	//Instead of array it is called Slice
	trees := []string{
		"Mango",
		"Custard Apple",
		"Pomegranate",
	}
	// fmt.Println("Line One", trees)
	//With the use of Slicing, We can avail different inbuilt methods on the array
	//Append method to add element to an array
	trees = append(trees, "Jackfruit")
	//To delete make use of helper module "slices"
	trees = slices.Delete(trees, 0, 1)

	//Cant append new things into array
	//For loop to print all the details of array
	for i := 0; i < len(animals); i++ {
		fmt.Printf("The animal is %s\n", animals[i])
	}
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	//Range instead of boundaries
	for value := range 10 {
		fmt.Println(value, value%2 == 0)
	}
	for index, value := range trees {
		fmt.Printf("This is index %d and this is the %s tree\n", index, value)
	}
	//There is no while loop in Go
	i := 1
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// fmt.Println(animals)
	// fmt.Println(trees)
}
