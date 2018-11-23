package main

import "fmt"

var g_groceries []string

func add_grocery(a ...string) {
	fmt.Println("capacity", cap(g_groceries))
	for _, d := range a {
		g_groceries = append(g_groceries, d)
	}
}

func list_groceries() {
	fmt.Println("grocery list is as follows:")
	/*for i := 0; i < len(g_groceries); i++ {
		fmt.Println(g_groceries[i])
	}*/

	for _, d := range g_groceries {
		fmt.Println(d)
	}
}

func main() {
	add_grocery("bread", "fruit cake", "cucumber", "potato chips", "ice cream")

	list_groceries()
}
