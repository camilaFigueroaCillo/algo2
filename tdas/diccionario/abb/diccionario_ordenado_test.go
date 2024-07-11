package diccionario_test

import (
	"strings"
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

var f = func(a, b string) int { return strings.Compare(a, b) }
var f2 = func(a, b int) int { return compare(a, b) }

func compare(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}

func TestDiccVacio(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, int](f)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("A") })
}

func TestUnElemento(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, int](f)
	dic.Guardar("A", 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.EqualValues(t, 10, dic.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("B") })
}

func TestGuardar(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, int](f)
	require.False(t, dic.Pertenece("A"))

	dic.Guardar("D", 40)
	require.EqualValues(t, 1, dic.Cantidad())
	dic.Guardar("A", 10)
	require.EqualValues(t, 2, dic.Cantidad())
	dic.Guardar("B", 20)
	require.EqualValues(t, 3, dic.Cantidad())
	dic.Guardar("F", 60)
	require.EqualValues(t, 4, dic.Cantidad())
	dic.Guardar("E", 50)
	require.EqualValues(t, 5, dic.Cantidad())
	dic.Guardar("H", 60)
	require.EqualValues(t, 6, dic.Cantidad())

	require.True(t, dic.Pertenece("A"))
	require.True(t, dic.Pertenece("B"))
	require.False(t, dic.Pertenece("C"))
	require.EqualValues(t, 10, dic.Obtener("A"))
	require.EqualValues(t, 20, dic.Obtener("B"))
	require.EqualValues(t, 50, dic.Obtener("E"))
	require.EqualValues(t, 60, dic.Obtener("F"))
	require.EqualValues(t, 60, dic.Obtener("H"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("C") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("L") })
}

func TestReemplazarValor(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, int](f)
	dic.Guardar("A", 10)
	require.EqualValues(t, 10, dic.Obtener("A"))
	dic.Guardar("A", 100)
	require.EqualValues(t, 100, dic.Obtener("A"))
	require.EqualValues(t, 1, dic.Cantidad())
}

func TestBorrar(t *testing.T) {
	//este borrar deberia cubrir los casos con 0 o 1 o 2 hijos
	dic := TDADiccionario.CrearABB[string, int](f)
	dic.Guardar("F", 10)
	dic.Guardar("D", 20)
	dic.Guardar("G", 30)
	dic.Guardar("X", 40)
	dic.Guardar("A", 50)
	require.EqualValues(t, 5, dic.Cantidad())

	dato := dic.Borrar("A")
	require.EqualValues(t, 50, dato)
	require.False(t, dic.Pertenece("A"))
	require.EqualValues(t, 4, dic.Cantidad())
	require.True(t, dic.Pertenece("F"))
	require.True(t, dic.Pertenece("D"))
	require.True(t, dic.Pertenece("G"))
	require.True(t, dic.Pertenece("X"))

	dato2 := dic.Borrar("F")
	require.EqualValues(t, 10, dato2)
	require.False(t, dic.Pertenece("F"))
	require.EqualValues(t, 3, dic.Cantidad())
	require.True(t, dic.Pertenece("D"))
	require.True(t, dic.Pertenece("G"))
	require.True(t, dic.Pertenece("X"))

	dato3 := dic.Borrar("G")
	require.EqualValues(t, 30, dato3)
	require.False(t, dic.Pertenece("G"))
	require.EqualValues(t, 2, dic.Cantidad())
	require.True(t, dic.Pertenece("D"))
	require.True(t, dic.Pertenece("X"))

	dato4 := dic.Borrar("X")
	require.EqualValues(t, 40, dato4)
	require.False(t, dic.Pertenece("X"))
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("D"))

	dato5 := dic.Borrar("D")
	require.EqualValues(t, 20, dato5)
	require.False(t, dic.Pertenece("D"))
	require.EqualValues(t, 0, dic.Cantidad())
}

func TestVolumen(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, float64](f2)
	for i := 1000; i < 2000; i++ {
		if i%2 == 0 {
			dic.Guardar(i/2, float64(i)+0.5)
			require.EqualValues(t, float64(i)+0.5, dic.Obtener(i/2))
		} else {
			dic.Guardar(i*2, float64(i)+0.5)
			require.EqualValues(t, float64(i)+0.5, dic.Obtener(i*2))
		}
	}
	require.EqualValues(t, 1000, dic.Cantidad())
	for i := 1000; i < 2000; i++ {
		if i%2 == 0 {
			require.EqualValues(t, float64(i)+0.5, dic.Borrar(i/2))
		} else {
			require.EqualValues(t, float64(i)+0.5, dic.Borrar(i*2))
		}
	}
	require.EqualValues(t, 0, dic.Cantidad())
}

func TestIteradorInterno(t *testing.T) {
	claves := []string{"L", "E", "P", "A", "W", "B", "R", "F", "O"}
	dic := TDADiccionario.CrearABB[string, int](f)
	for i, e := range claves {
		dic.Guardar(e, i)
		require.EqualValues(t, i, dic.Obtener(e))
	}
	suma := 0
	dic.Iterar(func(clave string, dato int) bool {
		suma += dato
		return true
	})
	require.EqualValues(t, 36, suma)
}

func TestIteradorInternoCondicion(t *testing.T) {
	claves := []string{"L", "E", "P", "A", "W", "B", "R", "F", "O"}
	dic := TDADiccionario.CrearABB[string, int](f)
	for i, e := range claves {
		dic.Guardar(e, i)
		require.EqualValues(t, i, dic.Obtener(e))
	}
	suma := 0
	dic.Iterar(func(clave string, dato int) bool {
		if dato%2 == 0 {
			suma += dato
		}
		return dato != 2
	})
	require.EqualValues(t, 10, suma)
}

func TestIteradorInternoCondicionAIzquierda(t *testing.T) {
	claves := []string{"D", "A", "E", "C", "F", "B"}
	dic := TDADiccionario.CrearABB[string, int](f)
	for i, e := range claves {
		dic.Guardar(e, i)
		require.EqualValues(t, i, dic.Obtener(e))
	}
	suma := 0
	dic.Iterar(func(clave string, dato int) bool {
		if dato%2 == 1 {
			suma += dato
		}
		return dato != 3
	})
	require.EqualValues(t, 9, suma)
}

func TestIteradorInternoVacio(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, int](f)
	suma := 0
	dic.Iterar(func(clave string, dato int) bool {
		suma += dato
		return true
	})
	require.EqualValues(t, 0, suma)
}

func TestIteradorInternoVacioCondicion(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, int](f)
	suma := 0
	dic.Iterar(func(clave string, dato int) bool {
		suma += dato
		return dato != 7
	})
	require.EqualValues(t, 0, suma)
}

func TestIteradorInternoRango(t *testing.T) {
	claves := []int{78, 32, 22, 41, 20, 50}
	dic := TDADiccionario.CrearABB[int, int](f2)
	for i, e := range claves {
		dic.Guardar(e, i)
		require.EqualValues(t, i, dic.Obtener(e))
	}
	suma := 0
	k1 := 20
	k2 := 50
	dic.IterarRango(&k1, &k2, func(clave int, dato int) bool {
		suma += dato
		return true
	})
	require.EqualValues(t, 15, suma)
}

func TestIteradorInternoRangoCondicion(t *testing.T) {
	claves := []int{78, 32, 22, 41, 20, 50}
	dic := TDADiccionario.CrearABB[int, int](f2)
	for i, e := range claves {
		dic.Guardar(e, i)
		require.EqualValues(t, i, dic.Obtener(e))
	}
	suma := 0
	k1 := 20
	k2 := 50
	dic.IterarRango(&k1, &k2, func(clave int, dato int) bool {
		if clave%2 == 0 {
			suma += dato
		}
		return true
	})
	require.EqualValues(t, 12, suma)
}

func TestIteradorInternoDesdeNil(t *testing.T) {
	claves := []int{78, 32, 22, 41, 20, 50}
	dic := TDADiccionario.CrearABB[int, int](f2)
	for i, e := range claves {
		dic.Guardar(e, i)
		require.EqualValues(t, i, dic.Obtener(e))
	}
	suma := 0
	k2 := 50
	dic.IterarRango(nil, &k2, func(clave int, dato int) bool {
		suma += dato
		return true
	})
	require.EqualValues(t, 15, suma)
}

func TestIteradorInternoHastaNil(t *testing.T) {
	claves := []int{78, 32, 22, 41, 20, 50}
	dic := TDADiccionario.CrearABB[int, int](f2)
	for i, e := range claves {
		dic.Guardar(e, i)
		require.EqualValues(t, i, dic.Obtener(e))
	}
	suma := 0
	k1 := 30
	dic.IterarRango(&k1, nil, func(clave int, dato int) bool {
		suma += dato
		return true
	})
	require.EqualValues(t, 9, suma)
}

func TestItInternoFueraDeRango(t *testing.T) {
	claves := []int{78, 32, 22, 41, 20, 50}
	dic := TDADiccionario.CrearABB[int, int](f2)
	for i, e := range claves {
		dic.Guardar(e, i)
		require.EqualValues(t, i, dic.Obtener(e))
	}
	suma := 0
	k1 := 80
	k2 := 100
	dic.IterarRango(&k1, &k2, func(clave int, dato int) bool {
		suma += dato
		return true
	})
	require.EqualValues(t, 0, suma)
}

func TestItInternoDmayorH(t *testing.T) {
	claves := []int{78, 32, 22, 41, 20, 50}
	dic := TDADiccionario.CrearABB[int, int](f2)
	for i, e := range claves {
		dic.Guardar(e, i)
		require.EqualValues(t, i, dic.Obtener(e))
	}
	suma := 0
	k1 := 100
	k2 := 50
	dic.IterarRango(&k1, &k2, func(clave int, dato int) bool {
		suma += dato
		return true
	})
	require.EqualValues(t, 0, suma)
}

func TestItInternoRgoUno(t *testing.T) {
	claves := []int{78, 32, 22, 41, 20, 50}
	dic := TDADiccionario.CrearABB[int, int](f2)
	for i, e := range claves {
		dic.Guardar(e, i)
		require.EqualValues(t, i, dic.Obtener(e))
	}
	suma := 0
	k1 := 30
	k2 := 40
	dic.IterarRango(&k1, &k2, func(clave int, dato int) bool {
		suma += dato
		return true
	})
	require.EqualValues(t, 1, suma)
}

func TestIteradorExternoVacio(t *testing.T) {
	dic := TDADiccionario.CrearABB[string, int](f)
	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIterarExterno(t *testing.T) {

	claves := []string{"river", "boca", "racing"}
	valores := []int{4, 6, 1}
	dic := TDADiccionario.CrearABB[string, int](f)
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])

	iter := dic.Iterador()

	require.True(t, iter.HaySiguiente())
	primerClave, _ := iter.VerActual()

	iter.Siguiente()
	segundaClave, _ := iter.VerActual()
	require.NotEqualValues(t, primerClave, segundaClave)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	tercerClave, _ := iter.VerActual()
	require.NotEqualValues(t, primerClave, tercerClave)
	require.NotEqualValues(t, segundaClave, tercerClave)
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorExternoNoLlegaAlFinal(t *testing.T) {

	dic := TDADiccionario.CrearABB[string, string](f)
	claves := []string{"A", "B", "C"}
	dic.Guardar(claves[0], "")
	dic.Guardar(claves[1], "")
	dic.Guardar(claves[2], "")

	dic.Iterador()
	iter2 := dic.Iterador()
	iter2.Siguiente()
	iter3 := dic.Iterador()
	primero, _ := iter3.VerActual()
	iter3.Siguiente()
	segundo, _ := iter3.VerActual()
	iter3.Siguiente()
	tercero, _ := iter3.VerActual()
	iter3.Siguiente()
	require.False(t, iter3.HaySiguiente())
	require.EqualValues(t, "A", primero)
	require.EqualValues(t, "B", segundo)
	require.EqualValues(t, "C", tercero)
}

func TestIterarExternoRango(t *testing.T) {

	claves := []int{4, 6, 1, 7, 0}
	valores := []string{"river", "boca", "racing", "independiente", "chacarita"}
	dic := TDADiccionario.CrearABB[int, string](f2)
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])
	dic.Guardar(claves[3], valores[3])
	dic.Guardar(claves[4], valores[4])

	k1 := 1
	k2 := 5
	iter := dic.IteradorRango(&k1, &k2)

	require.True(t, iter.HaySiguiente())
	primerClave, _ := iter.VerActual()
	iter.Siguiente()

	segundaClave, _ := iter.VerActual()
	require.EqualValues(t, 1, primerClave)
	require.EqualValues(t, 4, segundaClave)
	require.True(t, iter.HaySiguiente())
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorExternoRangoAbbCompleto(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](f2)
	dic.Guardar(100, 200)
	dic.Guardar(50, 100)
	dic.Guardar(150, 300)
	dic.Guardar(100, 200)
	dic.Guardar(25, 50)
	dic.Guardar(75, 150)
	dic.Guardar(125, 250)
	dic.Guardar(200, 400)
	dic.Guardar(10, 20)
	dic.Guardar(30, 40)
	dic.Guardar(60, 120)
	dic.Guardar(80, 160)
	dic.Guardar(110, 220)
	dic.Guardar(130, 260)
	dic.Guardar(190, 380)
	dic.Guardar(300, 600)

	recorridas := []int{60, 75, 80, 100, 110}
	k1 := 60
	k2 := 120
	iter := dic.IteradorRango(&k1, &k2)
	for i := 0; i < 5; i++ {
		clave, _ := iter.VerActual()
		require.EqualValues(t, recorridas[i], clave)
		require.True(t, iter.HaySiguiente())
		iter.Siguiente()
	}
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorExternoSinDesde(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](f2)
	dic.Guardar(8, 16)
	dic.Guardar(6, 12)
	dic.Guardar(10, 20)
	dic.Guardar(4, 8)
	dic.Guardar(7, 14)
	dic.Guardar(9, 18)
	dic.Guardar(11, 22)

	recorridas := []int{4, 6, 7, 8, 9}
	k1 := 9
	iter := dic.IteradorRango(nil, &k1)
	for i := 0; i < 5; i++ {
		clave, _ := iter.VerActual()
		require.EqualValues(t, recorridas[i], clave)
		require.True(t, iter.HaySiguiente())
		iter.Siguiente()
	}
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorExternoSinHasta(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](f2)
	dic.Guardar(8, 16)
	dic.Guardar(6, 12)
	dic.Guardar(10, 20)
	dic.Guardar(4, 8)
	dic.Guardar(7, 14)
	dic.Guardar(9, 18)
	dic.Guardar(11, 22)

	recorridas := []int{7, 8, 9, 10, 11}
	k1 := 7
	iter := dic.IteradorRango(&k1, nil)
	for i := 0; i < 5; i++ {
		clave, _ := iter.VerActual()
		require.EqualValues(t, recorridas[i], clave)
		require.True(t, iter.HaySiguiente())
		iter.Siguiente()
	}
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}
