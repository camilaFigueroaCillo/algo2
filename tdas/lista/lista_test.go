package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListaVacia(t *testing.T) {
	//lista se crea y esta vacia
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
}

func TestInsertarPrimeroVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())
}

func TestInsertarPrimeroNoVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(3)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 3, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	require.Equal(t, 3, lista.Largo())

}

func TestInsertarUltimoVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())
}

func TestInsertarUltimoNoVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(5)

	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.VerPrimero())
	require.Equal(t, 5, lista.VerUltimo())
	require.Equal(t, 3, lista.Largo())
}

func TestBorrarPrimeroVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
}

func TestBorrarPrimeroLenUno(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	require.Equal(t, 1, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
}

func TestBorrarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	require.Equal(t, 1, lista.BorrarPrimero())
	require.Equal(t, 2, lista.VerPrimero())
	require.False(t, lista.EstaVacia())
}

func TestListaEnlazada(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	for i := 0; i <= 10; i++ {
		lista.InsertarPrimero(i)
	}
	require.False(t, lista.EstaVacia())
	require.Equal(t, 11, lista.Largo())
	require.Equal(t, 10, lista.VerPrimero())
	require.Equal(t, 0, lista.VerUltimo())

	for i := 10; i >= 0; i-- {
		borrado := lista.BorrarPrimero()
		require.Equal(t, i, borrado)
	}
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.Equal(t, 0, lista.Largo())
	require.True(t, lista.EstaVacia())

	lista.InsertarUltimo(22)
	require.Equal(t, 22, lista.VerPrimero())

}

func TestIterarTodos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, e := range arr {
		lista.InsertarUltimo(e)
	}
	suma := 0
	lista.Iterar(func(v int) bool {
		if v%2 == 0 {
			suma += v
		}
		return true
	})

	require.Equal(t, 20, suma)
}

func TestIterarConCondicion(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, e := range arr {
		lista.InsertarUltimo(e)
	}
	suma := 0
	lista.Iterar(func(v int) bool {
		if v%2 == 0 {
			suma += v
		}
		if v == 7 {
			return false
		}
		return true
	})

	require.Equal(t, 12, suma)
}

func TestIterarListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	suma := 0
	lista.Iterar(func(v int) bool {
		return true
	})
	require.Equal(t, 0, suma)
}

func TestIterarListaVaciaCondicion(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	suma := 0
	lista.Iterar(func(v int) bool {
		return v%2 != 0
	})
	require.Equal(t, 0, suma)
}

func TestInsertarConIteradorListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[float64]()
	iter := lista.Iterador()
	iter.Insertar(3.14)
	iter.Siguiente()
	iter.Insertar(2.19)
	require.Equal(t, 3.14, lista.VerPrimero())
	require.Equal(t, 2.19, lista.VerUltimo())
	require.Equal(t, 2, lista.Largo())
}

func TestBorrarConIteradorListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	iter := lista.Iterador()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })
	iter.Insertar("A")
	iter.Insertar("B")
	borrado := iter.Borrar()
	require.Equal(t, borrado, "B")
	borrado2 := iter.Borrar()
	require.Equal(t, borrado2, "A")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })
}

func TestBorrarConIteradorAlPrincipio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(8)
	iter := lista.Iterador()
	dato := iter.Borrar()
	require.Equal(t, 8, dato)
	require.True(t, lista.EstaVacia())
}

func TestInsertarConIteradorAlPrincipio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarUltimo("A")
	iter := lista.Iterador()
	require.Equal(t, "A", iter.VerActual())
	iter.Insertar("B")
	require.Equal(t, "B", lista.VerPrimero())
	require.Equal(t, "A", lista.VerUltimo())
}

func TestVerPanicListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[[]int]()
	iter := lista.Iterador()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestInsertarAlFinalIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 10; i++ {
		lista.InsertarUltimo(i)
	}
	iter := lista.Iterador()
	for i := 0; i < lista.Largo(); i++ {
		iter.Siguiente()
	}
	iter.Insertar(10)
	require.Equal(t, 11, lista.Largo())
	require.Equal(t, 10, lista.VerUltimo())
}

func TestInsertarAlMedioIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		lista.InsertarUltimo(i)
	}
	iter := lista.Iterador()
	for i := 0; i < 10; i++ {
		if i == 5 {
			iter.Insertar(5)
		}
		iter.Siguiente()
	}
	require.Equal(t, 10, lista.Largo())
	for i := 0; i < 5; i++ {
		borrado := lista.BorrarPrimero()
		require.Equal(t, i, borrado)
	}
	require.Equal(t, 5, lista.VerPrimero())
}

func TestBorrarConIteradorAlFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(2)
	iter := lista.Iterador()
	iter.Siguiente()
	iter.Siguiente()
	borrado := iter.Borrar()
	require.Equal(t, 2, borrado)
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 3, lista.VerUltimo())
}

func TestBorrarConIteradorAlMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 10; i++ {
		lista.InsertarUltimo(i)
	}
	iter := lista.Iterador()
	for i := 0; i < 9; i++ {
		if i == 5 {
			borrado := iter.Borrar()
			require.Equal(t, 5, borrado)
		}
		iter.Siguiente()
	}
	require.Equal(t, 9, lista.Largo())
	for i := 0; i < 5; i++ {
		require.Equal(t, i, lista.BorrarPrimero())
	}
	require.Equal(t, 6, lista.VerPrimero())
}

func TestHaySiguiente(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarPrimero("A")
	lista.InsertarPrimero("B")
	lista.InsertarPrimero("C")
	lista.InsertarPrimero("D")
	lista.InsertarPrimero("E")
	iter := lista.Iterador()
	for i := range 5 {
		require.True(t, iter.HaySiguiente(), i)
		iter.Siguiente()
	}
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.False(t, iter.HaySiguiente())
}

func TestVerActual(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := range 10 {
		lista.InsertarUltimo(i)
	}
	iter := lista.Iterador()
	for i := range 10 {
		require.Equal(t, i, iter.VerActual())
		iter.Siguiente()
	}
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
}
