package main

import (
	"apresentacaoGo/arraysESlices"
	"apresentacaoGo/funcoesEMetodos"
	"fmt"
)

func main() {
	// Exemplo de uso das funções e métodos
	resultado := funcoesemetodos.Somar(3, 5)
	fmt.Println("Resultado da função:", resultado)

	ret := funcoesemetodos.Retangulo{Largura: 5, Altura: 3}
	area := ret.CalcularArea()
	fmt.Println("Área do retângulo:", area)

	// Exemplo de uso de arrays e slices
	meusNumeros := []int{1, 2, 3, 4, 5}
	fmt.Println("Array:", meusNumeros)

	meuSlice := arraysslices.CriarSlice(meusNumeros)
	fmt.Println("Slice:", meuSlice)

	meuSlice = arraysslices.AdicionarElemento(meuSlice, 6)
	fmt.Println("Slice após adicionar elemento:", meuSlice)

	meuSlice[0] = 10
	fmt.Println("Slice após modificar elemento:", meuSlice)

	soma := arraysslices.SomarElementos(meusNumeros)
	fmt.Println("Soma dos elementos do array:", soma)
}
