package cola_test

import (
	"github.com/stretchr/testify/require"
	TDACola "tdas/cola"
	"testing"
)

func TestColaVacia(t *testing.T) {
	// TDA Cola recien creada esta vacía y se comporta como tal.

	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestEncolarEnteros(t *testing.T) {

	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(100)
	cola.Encolar(200)
	require.False(t, cola.EstaVacia())
	require.Equal(t, 100, cola.VerPrimero())
}

func TestEncolarStrings(t *testing.T) {

	cola := TDACola.CrearColaEnlazada[string]()
	cola.Encolar("100")
	cola.Encolar("200")
	require.False(t, cola.EstaVacia())
	require.Equal(t, "100", cola.VerPrimero())
}

func TestEncolarBools(t *testing.T) {

	cola := TDACola.CrearColaEnlazada[bool]()
	cola.Encolar(true)
	cola.Encolar(false)
	require.False(t, cola.EstaVacia())
	require.Equal(t, true, cola.VerPrimero())

}

func TestVolumenEncolar(t *testing.T) {

	cola := TDACola.CrearColaEnlazada[int]()
	for i := range 1000 {
		cola.Encolar(i)
	}
	require.False(t, cola.EstaVacia())
	require.Equal(t, 0, cola.VerPrimero())
}

func TestDesencolar(t *testing.T) {

	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(145)
	cola.Encolar(175)
	require.Equal(t, 145, cola.Desencolar())
	require.Equal(t, 175, cola.Desencolar())

}

func TestEstaVaciaDesencolar(t *testing.T) {

	cola := TDACola.CrearColaEnlazada[int]()

	for i := 0; i <= 10; i++ {
		cola.Encolar(i)
	}

	for i := 0; i <= 10; i++ {
		cola.Desencolar()
	}
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}
