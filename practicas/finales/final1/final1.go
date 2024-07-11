package final

import (
	"strconv"
	h "tdas/cola_prioridad"
	d "tdas/diccionario"
)

/*Implementar un algoritmo que dado un arreglo de dígitos (0-9) determine cuál es el número más grande que se puede
formar con dichos dígitos."*/

func numMayor(arr []int) int {
	heap := h.CrearHeapArr[int](arr, func(a, b int) int {return a-b})
	cadena := ""
	for !heap.EstaVacia() {
		cadena += strconv.Itoa(heap.Desencolar())
	}
	n, _ := strconv.Atoi(cadena)
	return n
}

//La complejidad de este algoritmo es O(nlogn)


/*Implementar un algoritmo que reciba un arreglo desordenado de enteros, su largo (n) y un número K y determinar en
O(n) si existe un par de elementos en el arreglo que sumen exactamente K.*/

func sumaK(arr []int, largo, k int) bool{
	hash := d.CrearHash[int, int]()
	for _, n := range arr {
		if hash.Pertenece(n-k) || hash.Pertenece(k-n) {
			return true
		}
		hash.Guardar(n, 1)
	}
	return false
}
