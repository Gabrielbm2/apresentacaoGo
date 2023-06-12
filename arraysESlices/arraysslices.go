package arraysslices

// Função que retorna um slice a partir de um array
func CriarSlice(arr []int) []int {
	return arr[1:4]
}

// Função que adiciona um elemento ao final do slice
func AdicionarElemento(slice []int, elemento int) []int {
	return append(slice, elemento)
}

// Função que retorna a soma dos elementos de um array
func SomarElementos(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}
