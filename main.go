package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("Hello World")
	somme := add(3, 4)
	fmt.Printf("La somme de 3 et 4 est : %v\n", somme)
}
