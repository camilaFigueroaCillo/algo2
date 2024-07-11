package operaciones

import (
	"math"
)

func Suma(a, b int64) int64 {
	return int64(a + b)
}

func Resta(a, b int64) int64 {
	return int64(a - b)
}

func Division(a, b int64) int64 {
	return int64(a / b)
}

func Producto(a, b int64) int64 {
	return int64(a * b)
}

func RaizCuadrada(a int64) int64 {
	return int64(math.Sqrt(float64(a)))
}

func Potencia(base, exponente int64) int64 {
	return int64(math.Pow(float64(base), float64(exponente)))
}

func Logaritmo(a, base int64) int64 {
	log := math.Log(float64(a)) / math.Log(float64(base))
	return int64(log)

}

func Ternario(a, b, c int64) int64 {
	if a != 0 {
		return b
	}
	return c
}
