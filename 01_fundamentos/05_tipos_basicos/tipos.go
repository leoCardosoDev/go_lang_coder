package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// tipos inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))
	// sem sinal (só positivos)... uint8 uint16 uint32 uint64 uint
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal... int8 int16 int32 int64 int
	i1 := math.MaxInt64
	fmt.Println("O valor máximo é", i1)
	fmt.Println("O tipo é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento de tabela Unicode (int32)
	fmt.Println("O valor rune", i2)
	fmt.Println("O tipo é", reflect.TypeOf(i2))

	// números reais (float32, float64)
	var x float32 = 49.99
	y := float32(49.99)
	fmt.Println("O valor de y é", reflect.TypeOf(y))
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O valor do literal de 49.99 é", reflect.TypeOf(49.99))

	// boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo) // negando

	// string
	s1 := "Olá meu nome é Leo"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	// string com multlipas linhas
	 s2 := `
  Olá 
	meu nome 
	é 
	Leo!`
	fmt.Println(s2)
	fmt.Println("O tamanho da string é", len(s2))

	// char ???
	// no - var char = 'b 
	char := 'a'
	fmt.Println(char)
	fmt.Println("O tipo é char é", reflect.TypeOf(char))
}
