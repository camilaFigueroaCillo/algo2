package operaciones

import (
	"fmt"
	h "tdas/cola_prioridad"
	diccionario "tdas/diccionario"
)

type sitio struct {
	nombre     string
	frecuencia int
}

func VerMasVisitados(frecuenciasSitios diccionario.Diccionario[string, int], k int) {
	var largo int

	if k > frecuenciasSitios.Cantidad() {
		largo = frecuenciasSitios.Cantidad()
	} else {
		largo = k
	}

	arr := make([]sitio, largo)

	iter := frecuenciasSitios.Iterador()

	for i := 0; iter.HaySiguiente() && i < k; i++ {
		c, d := iter.VerActual()
		arr[i] = sitio{c, d}
		iter.Siguiente()
	}

	heap := h.CrearHeapArr[sitio](arr, func(s1, s2 sitio) int { return s1.frecuencia - s2.frecuencia }) //O(K)

	for iter.HaySiguiente() {

		sitioActual, frecActual := iter.VerActual()

		if frecActual > heap.VerMax().frecuencia {
			heap.Desencolar()
			heap.Encolar(sitio{sitioActual, frecActual})
		}
		iter.Siguiente()
	}

	imprimirMasVisitados(heap)
}

func imprimirMasVisitados(heap h.ColaPrioridad[sitio]) {

	fmt.Printf("Sitios más visitados:\n")

	for !heap.EstaVacia() {
		sitio := heap.Desencolar()
		fmt.Printf("\t%s - %d\n", sitio.nombre, sitio.frecuencia)
	}

	fmt.Printf("OK\n")
}
