package main

import (
	"math"
	l "tdas/lista"
	"fmt"
	"strings"
)

/*3. Implementar en Go una primitiva para el árbol binario func (ab *Arbol[T]) EsCompleto() que determine si el mismo
es un árbol completo. Indicar y justificar la complejidad de la primitiva. A fines del ejercicio, considerar que el árbol
está definido como:*/

type Arbol[T any] struct {
	clave T
	izq *Arbol[T]
	der *Arbol[T]
}


func (ab *Arbol[T]) EsCompleto() bool {
	altura := altura(ab)
	nodos := nodos(ab)
	return int(math.Pow(2, float64(altura))) - 1 == nodos	
}

func altura[T any](ab *Arbol[T]) int {
	if ab == nil {
		return 0
	}
	alturaIzq := altura(ab.izq) +1
	alturaDer := altura(ab.der) +1
	return max(alturaIzq, alturaDer)
}

func nodos[T any](ab *Arbol[T]) int {
	if ab == nil {
		return 0
	}
	izq := nodos(ab.izq)
	der := nodos(ab.der)
	return 1 + izq + der
}


/*Implementar una primitiva del ABB que dado un valor entero M, una clave inicial inicio y una clave final fin, se
devuelva una lista con todos los datos cuyas claves estén entre inicio y fin, que estén dentro de los primeros M
niveles del árbol (considerando a la raíz en nivel 1). Indicar y justificar la complejidad temporal.*/


type funcCmp[K comparable] func(K, K) int

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      funcCmp[K]
}
func (abb abb[K,V]) ClavesMedio(m int, ini, fin K) l.Lista[V] {
	lista := l.CrearListaEnlazada[V]()
	abb.clavesMedioRec(m, abb.raiz, lista, 1, ini, fin)
	return lista
}

func (abb abb[K, V])clavesMedioRec(nivel int, nodo *nodoAbb[K,V], lista l.Lista[V], nivel2 int, ini, fin K) {
	if nodo == nil {
		return
	}
	if nivel2 > nivel {
		return
	}
	abb.clavesMedioRec(nivel, nodo.izquierdo, lista, nivel2+1, ini, fin)
	abb.clavesMedioRec(nivel, nodo.derecho, lista, nivel2+1, ini, fin)
	if abb.cmp(nodo.clave, ini) >= 0 && abb.cmp(nodo.clave, fin) <= 0 && nivel2 <= nivel {
		lista.InsertarUltimo(nodo.dato)
	} 
}

/*	 10
	/ \
	5 15 		Un resultado final serían los datos de las
   / \ / \ 		claves 10, 5, 8, 15, 12 (en cualquier orden).
   3 8 12 20
     / \ 
    7 14*/





/*	Implementar una primitiva del árbol binario que determine si el mismo cumple con la propiedad de
	AVL. Indicar y justificar la complejidad del algoritmo implementado.
	A fines del ejercicio, considerar que el árbol tiene la siguiente estructura:
	type Arbol struct {
		izq *Arbol
		der *Arbol
		clave string
	}
*/
// debo verificar que sea abb y tambien que cumpla la propiedad de avl

func (ab *Arbol[int]) es_AVL() bool {
	_, condicion := ab.es_avl_rec()
	return condicion
}

func (ab *Arbol[int]) es_avl_rec() (int, bool) {
	if ab == nil {
		return 0, true
	}
	h_izq, cond_izq := ab.izq.es_avl_rec()
	h_der, cond_der := ab.der.es_avl_rec()

	altura := max(h_izq, h_der) + 1
	
	es_abb := (ab.izq != nil && ab.izq.clave < ab.clave) && (ab.der != nil && ab.der.clave > ab.clave)

	es_avl := math.Abs(float64(h_izq) - float64(h_der)) <= 1 

	if cond_izq && cond_der && es_abb && es_avl {
		return altura, true
	}
	return altura, false
	
}










func main(){
	raiz := nodoAbb[int, int]{nil, nil, 10, 10}
	izq1 := nodoAbb[int, int]{nil, nil, 5, 5}
	der1 := nodoAbb[int, int]{nil, nil, 15, 15}
	izq2 := nodoAbb[int, int]{nil, nil, 3, 3}
	izqder := nodoAbb[int, int]{nil, nil, 8, 8}
	derizq := nodoAbb[int, int]{nil, nil, 12, 12}
	der2 := nodoAbb[int, int]{nil, nil, 20, 20}
	izq3 := nodoAbb[int, int]{nil, nil, 7, 7}
	izqder2 := nodoAbb[int, int]{nil, nil, 14, 14}

	raiz.izquierdo = &izq1
	raiz.derecho = &der1
	izq1.izquierdo = &izq2
	izq1.derecho = &izqder
	izqder.izquierdo = &izq3
	der1.izquierdo = &derizq
	der1.derecho = &der2
	derizq.derecho = &izqder2 

	arbol := abb[int, int]{&raiz, 9, func(a, b int) int {return a-b}}

	lista := arbol.ClavesMedio(3, 5, 15)

	for i := lista.Iterador(); i.HaySiguiente(); i.Siguiente() {
		fmt.Println(i.VerActual())
	}

}


