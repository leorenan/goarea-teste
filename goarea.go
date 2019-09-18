package goarea

import "math"

//Pi é a proporção numerica
const Pi = 3.1416

//Circ retorna a area
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//Rect retorna a area do retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
