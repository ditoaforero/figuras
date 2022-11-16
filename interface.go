package figuras

import "fmt"

type Figura interface {
	calcularArea() float64
	calcularPerimetro() float64
}

func Medidas(figura Figura) {
	fmt.Println("Medidas:", figura)
	fmt.Println("Area: ", figura.calcularArea())
	fmt.Println("Perimetro: ", figura.calcularPerimetro())
}
