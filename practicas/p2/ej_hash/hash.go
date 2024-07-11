package main

import ( 
	"fmt"
	d "tdas/diccionario"
	l "tdas/lista"
)


/*Implementar una función que reciba un hash de claves genéricas y como dato una lista cuyos datos sean genéricos, y
devuelva un nuevo hash cuyas claves sean las mismas del recibido por parámetro, pero el valor sea el dato del 
medio de la lista correspondiente al dato original. 
Se sabe que cada lista tiene a lo sumo M elementos (un valor no constante).
El diccionario original (ni sus listas) deben verse modificados por esta función. 
Por ejemplo, si en el diccionario original hay una clave boquita y como dato una lista [1977, 1978, 2000, 2001, 2003, 2007], en el nuevo diccionario debe
estar boquita como clave y como dato 2000. 
Indicar y justificar la complejidad del algoritmo implementado. La firma
de la función debe ser:func DictMedio[K comparable, T any](dict Diccionario[K, Lista[T]]) Diccionario[K, T] */

func DictMedio[K comparable, T any](dict d.Diccionario[K, l.Lista[T]]) d.Diccionario[K, T] {
	nuevo := d.CrearHash[K, T]()
	for i := dict.Iterador(); i.HaySiguiente(); i.Siguiente() {
		clave, lista := i.VerActual()
		medio := lista.Largo() / 2
		iterLista := lista.Iterador()
		for j := 0; j < medio-1; j++ {
			iterLista.Siguiente()
		}
		nuevo.Guardar(clave, iterLista.VerActual())
	}
	return nuevo
}

/*Implementar una función que reciba una lista de números y un valor k y devuelva si hay dos elementos
dentro de dicha lista que sumen exactamente k. Indicar y justificar la complejidad de la función
implementada.*/

func suma_k(lista l.Lista[int], k int) bool {
	hash := d.CrearHash[int, int]()
	for i := lista.Iterador(); i.HaySiguiente(); i.Siguiente() {
		elem := i.VerActual()
		if hash.Pertenece(elem) {
			valor := hash.Obtener(elem)
			hash.Guardar(elem, valor+1)
		} else {
			hash.Guardar(elem, 1)
		}
	}

	for i := lista.Iterador(); i.HaySiguiente(); i.Siguiente() {
		n := i.VerActual()
		if k > n && ( (hash.Pertenece(k-n) && k-n != n) || (hash.Pertenece(k-n) && k-n == n && hash.Obtener(k-n) > 1) ) {
			return true
		}
		if n > k && ( (hash.Pertenece(n-k) && n-k != n) || (hash.Pertenece(n-k) && n-k == n && hash.Obtener(n-k) > 1) ) {
			return true
		}
	
	}
	return false
}


/*Se cuenta con un arreglo desordenado donde todos los elementos se encuentran duplicados, salvo 1.
Implementar un algoritmo que determine cuál es dicho elemento que aparece una única vez. Indicar
y justificar la complejidad del algoritmo implementado.*/

func encontrar_single(arr []int) int {
	hash := d.CrearHash[int, int]()

	for _, e := range arr {
		if hash.Pertenece(e) {
			v := hash.Obtener(e)
			hash.Guardar(e, v+1)
		} else {
			hash.Guardar(e, 1)
		}
	}

	for i := hash.Iterador(); i.HaySiguiente(); i.Siguiente() {
		clave, dato := i.VerActual()
		if dato == 1 {
			return clave
		}
	}
	return -1
}



func main(){
	lista := l.CrearListaEnlazada[int]()
	años := []int{1, 3, 4, 5, 6}
	for _, a := range años {
		lista.InsertarUltimo(a)
	}
	booleano := suma_k(lista, 25)
	fmt.Println(booleano)
}
