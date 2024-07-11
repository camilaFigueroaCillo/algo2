package main

import (
	"fmt"
	"strings"
	d "tdas/diccionario"
)

/*Implementar un algoritmo que obtenga la parte entera de la raíz de un número n entero en O (log n). Justificar la
complejidad de la primitiva implementada.*/

func raizEnteraRec(n, ini, fin int) int {
	if fin == ini {
		return 1
	}

	m := (ini+fin)/2

	if m*m == n || m*m == n-1 {
		return m
	}
	if m*m > n {
		return raizEnteraRec(n, ini, m-1)
	}
	return raizEnteraRec(n, m+1, fin)
}

func raizEntera(n int) int {
	return raizEnteraRec(n, 0, n)
}

//La complejidad del algoritmo se puede ver con el teorema maestro debido a que es una funcion recursiva con ecuación
//de recurrencia de la forma: T(n) = 1.T(n/2)+O(n°), entonces como log(1) == 0 la complejidad del algoritmo es 
//O(log n)


/*Implementar una primitiva para una Cola implementada como una estructura en arreglo (como la vista en clase),
Filtrar[T](func condicion(T) bool) Cola[T] que devuelva una nueva cola para la cual los elementos de la cola
original dan true en la función condicion pasada por parámetro. La cola original debe quedar intacta, y 
los elementos de la final deben tener el orden relativo que tenían en la original. 
Indicar y justificar la complejidad del algoritmo implementado.*/

type cola[T any] struct {
	primero int
	ultimo int
	datos []T
	cantidad int
}

func (cola *cola[T]) Filter(f func(T) bool) Cola[T] {
	otra := &cola[T]{0, 0, make([]T, cola.cantidad)} //con esto me aseguro de no redimensionar ni tener que llenar el array desde el inicio hasta otra pos
	for i := cola.primero; i < len(cola.datos); i++ {
		if !f(cola.datos[i]) {
			continue
		}
		otra.datos[otra.cantidad] = cola.datos[i]
		otra.ultimo = otra.cantidad
		otra.cantidad++
	}
	if cola.ultimo < len(cola.datos) {
		for i := 0; i <= cola.ultimo; i++ {
			if !f(cola.datos[i]) {
				continue
			}
		
			otra.datos[otra.cantidad] = cola.datos[i]
			otra.ultimo = otra.cantidad
			otra.cantidad++
		}
	}
	return otra
}

/*Implementar un algoritmo que dado un texto, devuelva cuál es la palabra más frecuente del mismo. 
Indicar y justificar la complejidad del algoritmo implementado. Nota: recordar que existe la 
función split(cadena, separador), que funciona en O(m), siendo m el largo de la cadena. */

func frecuencias(texto string) string {
	hash := d.CrearHash[string, int]()
	t := strings.Split(texto, " ")
	for _, palabra := range t {
		if hash.Pertenece(palabra) {
			v := hash.Obtener(palabra)
			hash.Guardar(palabra, v+1)
		} else {
			hash.Guardar(palabra, 1)
		}
	}
	var mayor int
	var clave string
	for i := hash.Iterador(); i.HaySiguiente(); i.Siguiente() {
		k, v := i.VerActual()
		if v > mayor {
			mayor = v
			clave = k
		}
	}
	return clave
}



func main() {
	fmt.Println(raizEntera(25))
	fmt.Println(raizEntera(10))

}
