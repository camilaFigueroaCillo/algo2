package main

import (
	"fmt"
	h "tdas/cola_prioridad"
	l "tdas/lista"
)

/* Implementar una función que reciba un arreglo A de n enteros y un número k y devuelva un nuevo arreglo en 
el que para cada posición i de dicho arreglo, contenga el resultado de la máxima suma obtenible entre k elementos entre los
índices [0;i] del arreglo A (incluyendo a i). 
Las primeras k − 1 posiciones del arreglo a devolver deben tener como valor -1. 
Por ejemplo, para el arreglo [1, 5, 3, 4, 2, 8] y k = 3, el resultado debe ser [-1, -1, 9, 12, 12, 17]. 
La complejidad del algoritmo debe ser mejor que O(n · k). 
Indicar y justificar la complejidad del algoritmo implementado.*/

//creo un array del largo n y a las primeras k-1 posiciones les pongo un -1 -> O(n)
//creo un heap de maximos hasta la posicicion k de elementos -> O(k)
//desencolo los 2 maximos, los sumo, appendeo el resultado al arrB O(3*logk)
//appendeo al heap la suma maxima O(1)

func sumasMax(arr []int, k int) []int {
	heap := h.CrearHeapArr[int](arr[:k], func(a, b int) int {return a-b})
	res := make([]int, len(arr))
	for i := 0; i < k-1; i++ { res[i] = -1 }
	
	for i := k-1; i < len(arr); i++ {
		var suma int
		suma = 0
		for !heap.EstaVacia() {
			suma = suma + heap.Desencolar()
		}

	}	
	return res
}

/*Implementar una función que reciba K listas ordenadas de enteros, cada una con una misma longitud
n, y combine estas listas en una única lista ordenada evitando repetidos (ninguna lista original tiene
repetidos, pero sí puede suceder que dos listas diferentes tengan un mismo valor). Indicar y justificar
la complejidad temporal adecuadamente.*/


func k_merge(arr []l.Lista[int]) l.Lista[int] {
	cdad_listas := len(arr)
	res := l.CrearListaEnlazada[int]()
	heap := h.CrearHeap(func(a, b l.IteradorLista[int]) int {return b.VerActual() -a.VerActual()})

	for i := 0; i < cdad_listas ; i++ {
		heap.Encolar(arr[i].Iterador())
	}

	for !heap.EstaVacia() {
		min := heap.Desencolar()
		if res.EstaVacia() || res.VerUltimo() != min.VerActual() {
			res.InsertarUltimo(min.VerActual())
		}
		if min.HaySiguiente() {
			min.Siguiente()
			heap.Encolar(min)
		}
	}
	return res
}

func main() {
	arr := []int{1, 5, 3, 4, 2, 8}
	res := sumasMax(arr, 3)
	fmt.Println(res)
}
