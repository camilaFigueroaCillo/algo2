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


/*Para implementar un TDA Cola de prioridad, nos proponen la siguiente solución: usar un arreglo
desordenado (arr) para insertar los datos, y una variable (maximo) para poder obtener el máximo en
O(1). Se mantiene actualizada la variable maximo cada vez que se encola o desencola. ¿Es una buena
solución en el caso general? Justificar (recomendación: comparar contra la implementación de colas
de prioridad vista en clase).

A mi parecer no es una buena implementacion porque si bien yo voy a tener el valor del maximo para tenerlo a mano siempre
no tengo su indice en el array, por lo que al desencolar deberia buscar en el array ese elemento, sacarlo cuidar de no romper el array
porque puede estar en el medio y deberia correr los elementos de lugar, por lo que seria no solo mas complicado si no tambien
ineficiente al momento de desencolar, pues en el peor caso el maximo quedo en el ultimo lugar y voy a tener que hacer una 
busqueda lineal del maximo para desencolarlo
Tambien digamos, como calculo los hijos si el array esta desordenado? se romperia la propiedad del heap si yo al calcular un hijo es mayor a mi 
pero luego resulta que en realidad ese no es mi hijo y etc.
*/



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
