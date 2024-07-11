package cola_prioridad_test

import (
	"cmp"
	TDAPQ "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

var max = func(a, b int) int { return a - b }
var min = func(a, b int) int { return b - a }
var max_str = func(a, b string) int { return cmp.Compare[string](a, b) }
var min_str = func(a, b string) int { return cmp.Compare[string](b, a) }

func TestHeapVacio(t *testing.T) {

	//Prueba que el heap recién creado esté vacío y se comporte copmo tal.

	heap := TDAPQ.CrearHeap[int](max)
	require.True(t, heap.EstaVacia())
	require.Equal(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestColaDeMaximos(t *testing.T) {

	//Prueba que se encolen los elementos con prioridad de manera correcta

	heap := TDAPQ.CrearHeap[int](max)

	for i := 0; i <= 10; i++ {
		heap.Encolar(i)
	}
	require.False(t, heap.EstaVacia())
	require.Equal(t, 11, heap.Cantidad())
	require.Equal(t, 10, heap.VerMax())
}

func TestColaDeMinimos(t *testing.T) {

	//Prueba que se encolen los elementos con prioridad de manera correcta

	heap := TDAPQ.CrearHeap[int](min)

	for i := 10; i >= 0; i-- {
		heap.Encolar(i)
	}

	require.False(t, heap.EstaVacia())
	require.Equal(t, 11, heap.Cantidad())
	require.Equal(t, 0, heap.VerMax())
}

func TestVolumenEncolar(t *testing.T) {

	//Se pueden encolar grandes cantidades de elementos

	heap := TDAPQ.CrearHeap[int](max)

	for i := 0; i <= 1000; i++ {
		heap.Encolar(i)
	}

	require.False(t, heap.EstaVacia())
	require.Equal(t, 1001, heap.Cantidad())
	require.Equal(t, 1000, heap.VerMax())
}

func TestDesencolarMaximos(t *testing.T) {

	//Prueba que al desencolar, siempre desencole el maximo

	arr := []int{7, 1, 4, 11, 32, 58, 44, 10, 1, 6}

	ordenado := []int{58, 44, 32, 11, 10, 7, 6, 4, 1, 1}

	heap := TDAPQ.CrearHeap[int](max)

	for _, e := range arr {
		heap.Encolar(e)
	}

	for i := 0; i < len(ordenado); i++ {
		require.Equal(t, ordenado[i], heap.VerMax())
		require.Equal(t, ordenado[i], heap.Desencolar())
	}

}

func TestDesencolarMinimos(t *testing.T) {

	//Prueba que al desencolar, siempre desencole el minimo

	arr := []int{7, 1, 4, 11, 32, 58, 44, 10, 1, 6}

	ordenado := []int{1, 1, 4, 6, 7, 10, 11, 32, 44, 58}

	heap := TDAPQ.CrearHeap[int](min)

	for _, e := range arr {
		heap.Encolar(e)
	}

	for i := 0; i < len(ordenado); i++ {
		require.Equal(t, ordenado[i], heap.VerMax())
		require.Equal(t, ordenado[i], heap.Desencolar())
	}

}

func TestEstaVaciaDesencolar(t *testing.T) {
	//Prueba que se comporte como un heap vacío luego de encolar y desencolar elementos

	heap := TDAPQ.CrearHeap[int](max)

	for i := 0; i <= 10; i++ {
		heap.Encolar(i)
	}

	for i := 0; i <= 10; i++ {
		heap.Desencolar()
	}
	require.True(t, heap.EstaVacia())
	require.Equal(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestOrdenadoInverso(t *testing.T) {

	// Prueba encolar y desencolar en un heap donde los elementos se ingresan de menor a mayor

	arr := []string{"a", "b", "c", "d", "d", "e"}
	heap := TDAPQ.CrearHeap(max_str)

	for i := 0; i < 6; i++ {
		heap.Encolar(arr[i])
		require.Equal(t, arr[i], heap.VerMax())
	}

	for i := 0; i < 6; i++ {
		require.Equal(t, arr[5-i], heap.Desencolar())
	}

	require.True(t, heap.EstaVacia())
	require.Equal(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })

}

func TestHeapify(t *testing.T) {

	arr := []int{2, 8, 4, 1, 90, 32, 0}
	heap := TDAPQ.CrearHeapArr(arr, max)

	require.True(t, heap.Cantidad() == 7)
	require.Equal(t, 90, heap.VerMax())
	require.Equal(t, 90, heap.Desencolar())
	require.Equal(t, 32, heap.VerMax())
	require.Equal(t, 32, heap.Desencolar())
	require.Equal(t, 8, heap.VerMax())
	require.Equal(t, 8, heap.Desencolar())
	require.Equal(t, 4, heap.VerMax())
	require.Equal(t, 4, heap.Desencolar())
	require.Equal(t, 2, heap.VerMax())
	require.Equal(t, 2, heap.Desencolar())
	require.Equal(t, 1, heap.VerMax())
	require.Equal(t, 1, heap.Desencolar())
	require.Equal(t, 0, heap.VerMax())
	require.Equal(t, 0, heap.Desencolar())
	require.True(t, heap.EstaVacia())
	require.Equal(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })

}

func TestHeapifyVacio(t *testing.T) {

	// heapify de un arreglo vacio deja el arreglo vacio

	arr := []string{}
	heap := TDAPQ.CrearHeapArr(arr, min_str)

	require.Equal(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapifyVacioInt(t *testing.T) {

	arr := []int{}
	heap := TDAPQ.CrearHeapArr(arr, max)

	require.Equal(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapifyVolumen(t *testing.T) {

	arr := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		arr[i] = i
	}
	heap := TDAPQ.CrearHeapArr(arr, max)

	for i := 0; i < 9999; i++ {
		require.Equal(t, 9999-i, heap.Desencolar())
	}

}

func TestHeapifyVolumenMin(t *testing.T) {

	arr := []int{}
	for i := 0; i > 10000; i++ {
		arr = append(arr, 10000-i)
	}
	heap := TDAPQ.CrearHeapArr(arr, min)

	for i, _ := range arr {
		require.Equal(t, i, heap.Desencolar())
	}

}

func TestHeapifyEsHeap(t *testing.T) {

	// crearheaparr crea un heap correctamente

	arr := []string{"f", "d", "e", "a", "c", "d"}
	heap := TDAPQ.CrearHeapArr(arr, max_str)

	require.Equal(t, "f", heap.Desencolar())
	require.Equal(t, "e", heap.Desencolar())
	require.Equal(t, "d", heap.Desencolar())
	require.Equal(t, "d", heap.Desencolar())
	require.Equal(t, "c", heap.Desencolar())
	require.Equal(t, "a", heap.Desencolar())
	require.Equal(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapsort(t *testing.T) {
	arr := []int{4, 2, 93, 2, 1, 3, 23, 12, 15, 45}
	TDAPQ.HeapSort(arr, max)
	require.Equal(t, 1, arr[0])
	require.Equal(t, 2, arr[1])
	require.Equal(t, 2, arr[2])
	require.Equal(t, 3, arr[3])
	require.Equal(t, 4, arr[4])
	require.Equal(t, 12, arr[5])
	require.Equal(t, 15, arr[6])
	require.Equal(t, 23, arr[7])
	require.Equal(t, 45, arr[8])
	require.Equal(t, 93, arr[9])
}

func TestHeapsortVacio(t *testing.T) {

	// heapsort de arreglo vacio no modifica el arreglo

	arr := []string{}
	TDAPQ.HeapSort(arr, max_str)

	require.Equal(t, 0, len(arr))
}

func TestHeapsortVolumen(t *testing.T) {
	arr := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		arr[i] = i
	}
	TDAPQ.HeapSort(arr, min)
	for i := 0; i < 10000; i++ {
		require.Equal(t, 9999-i, arr[i])
	}
}

func TestHeapsortVolumenMin(t *testing.T) {
	arr := []int{}
	for i := 10000; i >= 0; i-- {
		arr = append(arr, i)
	}
	TDAPQ.HeapSort(arr, max)
	for i := 0; i < 10000; i++ {
		require.Equal(t, i, arr[i])
	}
}

func TestHeapSortOrdenado(t *testing.T) {

	// heapsort de arreglo ordenado no modifica el arreglo

	arr := []string{"x", "s", "m", "m", "d", "c", "a", "a"}
	TDAPQ.HeapSort(arr, min_str)

	require.Equal(t, "x", arr[0])
	require.Equal(t, "s", arr[1])
	require.Equal(t, "m", arr[2])
	require.Equal(t, "m", arr[3])
	require.Equal(t, "d", arr[4])
	require.Equal(t, "c", arr[5])
	require.Equal(t, "a", arr[6])
	require.Equal(t, "a", arr[7])
}

func TestHeapSortOrdenadoAlReves(t *testing.T) {
	arr := []string{"x", "s", "m", "m", "d", "c", "a", "a"}
	TDAPQ.HeapSort(arr, max_str)

	require.Equal(t, "a", arr[0])
	require.Equal(t, "a", arr[1])
	require.Equal(t, "c", arr[2])
	require.Equal(t, "d", arr[3])
	require.Equal(t, "m", arr[4])
	require.Equal(t, "m", arr[5])
	require.Equal(t, "s", arr[6])
	require.Equal(t, "x", arr[7])
}
