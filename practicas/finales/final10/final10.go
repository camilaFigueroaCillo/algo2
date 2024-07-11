package main

import (
	p "tdas/pila"
)
/*Un algoritmo iterativo sencillo para obtener la potencia de un número (b'n) tiene complejidad O(n). 
Tal como vieron en el secundario (esperamos), sabemos que b'n = b'n/2'2. 
Utilizar esta propiedad para implementar un algoritmo que calcule b'n, en tiempo O(log n). 
Justificar la complejidad del algoritmo implementado. Recordar tener cuidado con el caso que
n sea un valor impar.*/

func potencia(b, n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return b
	}

	if n % 2 == 0 {
		return potencia(b, n/2)*potencia(b, n/2)
	}
	return potencia(b, (n-1)/2)*potencia(b, (n-1)/2)*b
} 

/*Implementar una función que dada una pila, determine si la misma se encuentra ordenada (es decir, se ingresaron los
elementos de menor a mayor). La pila debe quedar en el mismo estado al original al terminar la ejecución de la función.
Indicar y justificar la complejidad de la función.*/

func estaOrdenada(pila p.Pila[int]) bool {
	aux := p.CrearPilaDinamica[int]()
	for !pila.EstaVacia() {
		aux.Apilar(pila.Desapilar())
	}
	var condicion bool
	for !aux.EstaVacia() {
		var b int
		a := aux.Desapilar()
		if !aux.EstaVacia() {
			b = pila.VerTope()
		}
		if a > b {
			condicion = false
		}
		pila.Apilar(a)
	}
	return condicion
}

//La complejidad de la funcion es O(n) porque veo todos los elementos de la pila yr ealizo operaciones constantes.
