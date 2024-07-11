package main

import (
	d "tdas/diccionario"
)


/*Implementar en Go una primitiva para la lista enlazada Filter(func(T) bool) Lista[T] que reciba una función
por parámetro, y devuelva una lista nueva que contenga los elementos de la original para los cuales la función 
pasada por parámetro devuelve true. 
El orden relativo de los elementos debe ser el mismo que en la lista original. La lista
original debe quedar en el mismo estado que el original. 
Indicar y justificar la complejidad de la primitiva.*/

func (lista *listaEnlazada[T]) Filter(f func(T) bool) Lista[T] {
	otra := &listaEnlazada{nil, nil, 0}
	actual = lista.primero
	for actual != nil {
		if f(actual.dato) {
			nuevoNodo := nodoCrear(actual.dato)
			if otra.primero == nil {
				otra.primero = nuevoNodo
			} else {
				otra.ultimo.siguiente = nuevoNodo
			}
			otra.ultimo = nuevoNodo
			otra.largo++
		}
		actual.siguiente
	}
	return otra
}
// La complejidad del algoritmo es O(N), osea lineal en la cdad de elementos de la lista original
//Porque recorro toda la lista, paso por cada nodo 1 sola vez y realizo operacioes constantes.

/*Implementar un algoritmo que cuente la mínima cantidad de elementos a eliminar de un arreglo para que todos los
elementos sean iguales. Indicar y justificar la complejidad de la función.*/

func elementosIguales[K comparable](arr []K) int {
	hash := d.CrearHash[K, int]()
	for _, elemento := range arr {
		if hash.Pertenece(elemento) {
			v := hash.Obtener(elemento)
			hash.Guardar(elemento, v+1)
		} else {
			hash.Guardar(elemento, 1)
		}
	}
	var suma int
	var mayorFreq int
	for i:= hash.Iterador(); i.HaySiguiente(); i.Siguiente() {
		_, valor := i.VerActual()
		if valor > mayorFreq {
			mayorFreq = valor
		}
		suma = suma + valor
	}

	return suma - mayorFreq
}

//La complejidad del algoritmo es O(N) porque recorro el arreglo 1 vez realizando operaciones constantes, Luego 
//Recorro todo el diccionario realizando operaciones constantes y recorrer el diccionario es a lo sumo O(N) entonces
//La complejidad final del algoritmo es O(N)

