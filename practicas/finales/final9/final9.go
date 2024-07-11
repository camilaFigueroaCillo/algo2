package main

import (
	"fmt"
	"math"
	d "tdas/diccionario"
	e "tdas/cola_prioridad"
)

/*Implementar un algoritmo que, dado un árbol binario, determine si el mismo es completo (es decir,
 que todos los niveles que tenga estén completos). 
 Indicar y justificar la complejidad del algoritmo implementado.*/

func (ab *Ab) es_completo() bool {
	if ab == nil {
		return true
	}
	altura := ab.Altura()
	nodos := ab.ContarNodos()
	return math.Pow(2, altura-1) == nodos
}

func (ab *Ab) Altura() int {
	if ab == nil {
		return 0
	}
	izq := ab.izq.Altura()
	der := ab.der.Altura()
	return maximo(izq, der) +1
}

func maximo(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (ab *Ab) ContarNodos() int {
	if ab == nil {
		return 0
	}
	izq := ab.izq.ContarNodos()
	der := ab.der.ContarNodos()
	return izq+der+1
}

//La complejidad del algoritmo es lineal porque veo todos los nodos del arbol y realizo operaciones constantes

/*Implementar un algoritmo que reciba dos arreglos desordenados y determine si ambos arreglos tienen los mismos
elementos (y en mismas cantidades). Indicar y justificar la complejidad del algoritmo implementado."""*/

func sonArreglosIguales[T any](arr1, arr2 []T) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	hash := d.CrearHash[any, int]()
	for _, e := range arr1 {
		if hash.Pertenece(e) {
			v := hash.Obtener(e)
			hash.Guardar(e, v+1)
		} else if !hash.Pertenece(e) {
			hash.Guardar(e, 1)
		}
	}

	for _, e := range arr2 {
		if !hash.Pertenece(e) {
			return false
		}
		v := hash.Obtener(e)
		hash.Guardar(e, v+1)
	}

	for i := hash.Iterador(); i.HaySiguiente(); i.Siguiente() {
		_, d := i.VerActual()
		if d%2 != 0 {
			return false
		}
	}
	return true
}

//La complejidad del algoritmo es lineal porque paso por todos los elementos del arr1 y del arr2 y realizo operaciones constantes


/*Trabajamos para una escuela primaria muy estructurada. En dicha escuela hay k cursos, cada uno con m alumnos 
(es decir, hay un total de n = k · m alumnos). 
Todas las mañanas hay que armar filas para cantar Aurora en el patio del
colegio. 
Primero los alumnos se ubican en una fila correspondiente a su curso, de menor a mayor altura para cantar.

Una vez terminado, proceden a entrar a la escuela de a un alumno por vez, pero deben hacerlo de menor a mayor 
altura.

Es decir, se debe ordenar a todos los alumnos de menor a mayor. 
Nosotros sabemos que esto es ineficiente (suelen usar mergesort, así que es O(n log n)), y desaprovechamos que 
los alumnos ya estaban ordenados por cursos. 
Implementar un algoritmo que reciba k filas (arreglos) de alumnos, cada una previamente ordenada de menor a mayor
altura, y nos devuelva un único arreglo con todos los alumnos ordeados por altura en tiempo menor a O(n log n). 
Indicar y justificar la complejidad del algoritmo implementado.*/

type Alumno struct {
	nombre string
	curso int
	altura int
}

type combinado struct {
	vector int
	pos int
	dato Alumno
}

func ordenarALumnos(alumnos [][]Alumno) []Alumno {
	res := make([]Alumno, len(alumnos)*len(alumnos[0]))
	heap := h.CrearHeap(func(a, b combinado) int {return b.dato.altura-a.dato.altura})
	for i := range alumnos {
		nodo := combinado{i, 0, alumnos[i][0]}
		heap.Encolar(nodo)

	}

	for i := 0; i < len(alumnos)*len(alumnos[0]); i++ {
		min := heap.Desencolar()
		res[i] = min.dato
		if min.pos >= len(min.vector) {
			continue
		}
		nodo := combinado{min.vector, min.pos+1, alumnos[min.vector][min.pos+1]}
		heap.Encolar(nodo)
	}
	return res
}

//La complejidad del algoritmo es O(nlogk) porque termino encolando y desencolando los n elementos del array y siempre mantengo
//El heap con k elementos por lo que las operaciones de encolar y desencolar son logk

