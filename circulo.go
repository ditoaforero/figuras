package figuras

import "math"

type Circulo struct {
	Radio float64
}

func (circulo *Circulo) calcularArea() float64 {
	return float64(circulo.Radio*circulo.Radio) * math.Pi
}

func (circulo *Circulo) calcularPerimetro() float64 {
	return float64(2*circulo.Radio) * math.Pi
}
