package funcoesemetodos

// Função que retorna a soma de dois números inteiros
func Somar(a, b int) int {
	return a + b
}

// Definição de uma estrutura chamada Retangulo
type Retangulo struct {
	Largura, Altura int
}

// Método CalcularArea associado à estrutura Retangulo
func (r Retangulo) CalcularArea() int {
	return r.Largura * r.Altura
}
