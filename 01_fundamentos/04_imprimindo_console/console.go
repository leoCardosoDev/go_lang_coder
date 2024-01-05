package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha. \n")


	fmt.Println("Nova")
	fmt.Println("linha")

	x := 3.141516
	fmt.Println("O valor de x é", x)
	fmt.Printf("O valor de x é %.2f", x)
	fmt.Printf("\nO valor de x é %f\n", x)
	a := 1
	b := 1.9999
	c := false
	d := "opa"
	// %f de float
	// %s de string
	// %d de inteiro
	// %t de boolean
	// %v todos os tipos
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v\n", a, b, c, d)
}
