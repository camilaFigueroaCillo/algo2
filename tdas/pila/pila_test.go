package pila_test

import (
	"github.com/stretchr/testify/require"
	TDAPila "tdas/pila"
	"testing"
)

func TestPilaVacia(t *testing.T) {
	//Pila vacía y recién creada se comporta como tal

	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestApilarEnteros(t *testing.T) {
	//Prueba que puedan apilarse datos de tipo integer

	pilaEnteros := TDAPila.CrearPilaDinamica[int]()
	pilaEnteros.Apilar(9)
	require.Equal(t, 9, pilaEnteros.VerTope(), "El tope de la pila debe ser 9")
	pilaEnteros.Apilar(27)
	require.Equal(t, 27, pilaEnteros.VerTope(), "El tope de la pila c +1 elem debe ser 27")

}

func TestApilarStrings(t *testing.T) {
	//Prueba que puedan apilarse strings correctamente

	pilaStrings := TDAPila.CrearPilaDinamica[string]()
	pilaStrings.Apilar("nueve")
	require.Equal(t, "nueve", pilaStrings.VerTope(), "El tope de la pila debe ser 'nueve'")
	pilaStrings.Apilar("diez")
	require.Equal(t, "diez", pilaStrings.VerTope(), "El tope de la pila c/ +1 elem debe ser 'diez'")

}

func TestApilarBools(t *testing.T) {
	//Prueba que puedan apilarse booleanos correctamente

	pilaBool := TDAPila.CrearPilaDinamica[bool]()
	pilaBool.Apilar(true)
	require.Equal(t, true, pilaBool.VerTope(), "El tope de la pila debe ser true")
	pilaBool.Apilar(false)
	require.Equal(t, false, pilaBool.VerTope(), "El tope de la pila c/ +1 elem debe ser false")

}

func TestVolumenApilar(t *testing.T) {
	//Se pueden apilar grandes cantidades de elementos

	pilaEnteros := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i <= 1000; i++ {
		pilaEnteros.Apilar(i)
	}
	require.Equal(t, 1000, pilaEnteros.VerTope(), "El tope debe ser 1000 despues de haberse apilado 1000 elementos")
}

func TestDesapilar(t *testing.T) {
	//Desapilar enteros

	pilaEnteros := TDAPila.CrearPilaDinamica[int]()
	pilaEnteros.Apilar(9)
	require.Equal(t, 9, pilaEnteros.Desapilar(), "prueba desapilar cuando hay un elemento en la pila")
	pilaEnteros.Apilar(27)
	pilaEnteros.Apilar(48)
	require.Equal(t, 48, pilaEnteros.Desapilar(), "prueba desapilar cuando hay mas de un elemento en la pila")

}

func TestEstaVaciaDesapilar(t *testing.T) {
	//Prueba que se comporte como una pila vacía luego de apilar y desapilar elementos

	pilaEnteros := TDAPila.CrearPilaDinamica[int]()

	for i := 0; i <= 10; i++ {
		pilaEnteros.Apilar(i)
	}

	for i := 0; i <= 10; i++ {
		pilaEnteros.Desapilar()
	}
	require.True(t, pilaEnteros.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaEnteros.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaEnteros.Desapilar() })
}

func TestVolumenDesapilar(t *testing.T) {
	//Prueba que se desapilen grandes cantidades de elementos

	pilaEnteros := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i <= 1000; i++ {
		pilaEnteros.Apilar(i)
	}
	for i := 0; i <= 1000; i++ {
		pilaEnteros.Desapilar()
	}
	require.True(t, pilaEnteros.EstaVacia(), "Prueba que esta vacía luego de desapilar todos los elementos")
}
