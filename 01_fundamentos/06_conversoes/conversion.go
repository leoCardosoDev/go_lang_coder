package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))
	// fmt.Println(x / y) // not working
	fmt.Println(int(x)/ y)
	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado
	// converta int para string antes de concatenar
	// fmt.Println("Teste " + string(123)) // resltado vai ser o codigo da tabela asci

	// Maneiras correta para converter para string
	fmt.Println("Teste " + strconv.Itoa(123))
	fmt.Println("Teste", 123)

	// string para int
	num, _ := strconv.Atoi("123") // retirna 2 valore: numero(num) ou error(_) _ ignora o retorno
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("1") // apenas a string "true" e "1", o resto é false
	fmt.Println(b)
	fmt.Println(b, reflect.TypeOf(b))
}
