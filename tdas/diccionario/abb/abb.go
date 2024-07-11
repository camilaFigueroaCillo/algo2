package diccionario

import (
	"tdas/pila"
)

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

type iterDiccionarioOrd[K comparable, V any] struct {
	actual      *nodoAbb[K, V]
	desde       *K
	hasta       *K
	diccionario *abb[K, V]
	pila        pila.Pila[*nodoAbb[K, V]]
}

func (abb *abb[K, V]) Guardar(clave K, dato V) {

	ant, act := abb.buscarNodo(abb.raiz, clave, nil)

	if act != nil {
		act.dato = dato
		return
	}

	nuevoNodo := crearNodo(clave, dato, nil, nil)
	abb.cantidad++

	if abb.raiz == nil {
		abb.raiz = nuevoNodo
	} else {
		if abb.cmp(ant.clave, nuevoNodo.clave) > 0 {
			ant.izquierdo = nuevoNodo
		} else {
			ant.derecho = nuevoNodo
		}
	}
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	_, act := abb.buscarNodo(abb.raiz, clave, nil)
	return act != nil
}

func (abb *abb[K, V]) Obtener(clave K) V {
	_, act := abb.buscarNodo(abb.raiz, clave, nil)
	if act == nil {
		panic("La clave no pertenece al diccionario")
	}
	return act.dato
}

func (abb *abb[K, V]) Borrar(clave K) V {
	ant, act := abb.buscarNodo(abb.raiz, clave, nil)
	if act == nil {
		panic("La clave no pertenece al diccionario")
	}
	abb.cantidad--
	var mayorIzquierda *nodoAbb[K, V]
	if act.derecho != nil && act.izquierdo != nil {
		ant, mayorIzquierda = buscarDerecha(act, act.izquierdo)
		swap(act, mayorIzquierda)
		act = mayorIzquierda
	}
	if act.derecho != nil {
		return abb.suprimir(ant, act, act.derecho)
	} else {
		return abb.suprimir(ant, act, act.izquierdo)
	}
}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	abb.recorrer(abb.raiz, nil, nil, visitar)
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	abb.recorrer(abb.raiz, desde, hasta, visitar)
}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return abb.IteradorRango(nil, nil)
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	pila := pila.CrearPilaDinamica[*nodoAbb[K, V]]()
	apilarAIzquierda(abb, pila, abb.raiz, desde, hasta)
	return &iterDiccionarioOrd[K, V]{abb.raiz, desde, hasta, abb, pila}
}

func (iter *iterDiccionarioOrd[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	ant := iter.pila.Desapilar()
	apilarAIzquierda(iter.diccionario, iter.pila, ant.derecho, iter.desde, iter.hasta)
}

func (iter *iterDiccionarioOrd[K, V]) HaySiguiente() bool {
	return !iter.pila.EstaVacia()
}

func (iter *iterDiccionarioOrd[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := iter.pila.VerTope()
	return nodo.clave, nodo.dato
}

func (abb *abb[K, V]) buscarNodo(nodo *nodoAbb[K, V], clave K, ant *nodoAbb[K, V]) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	if nodo == nil {
		return ant, nodo
	}
	if abb.cmp(nodo.clave, clave) > 0 {
		return abb.buscarNodo(nodo.izquierdo, clave, nodo)
	} else if abb.cmp(nodo.clave, clave) < 0 {
		return abb.buscarNodo(nodo.derecho, clave, nodo)
	} else {
		return ant, nodo
	}
}

func (abb *abb[K, V]) recorrer(nodo *nodoAbb[K, V], desde *K, hasta *K, f func(K, V) bool) bool {

	estado := true

	if nodo == nil {
		return estado
	}

	if desde == nil || abb.cmp(nodo.clave, *desde) > 0 {
		estado = abb.recorrer(nodo.izquierdo, desde, hasta, f)
	}

	if !estado {
		return estado
	}

	if estaEnRango(abb, nodo.clave, desde, hasta) {
		if !f(nodo.clave, nodo.dato) {
			return false
		}
	}

	if hasta == nil || abb.cmp(nodo.clave, *hasta) < 0 {
		estado = abb.recorrer(nodo.derecho, desde, hasta, f)
	}

	return estado
}

func (abb *abb[K, V]) suprimir(ant, act, hijo *nodoAbb[K, V]) V {
	dato := act.dato
	if act == abb.raiz {
		abb.raiz = hijo
	} else if ant.derecho == act {
		ant.derecho = hijo
	} else {
		ant.izquierdo = hijo
	}
	return dato
}

func estaEnRango[K comparable, V any](abb *abb[K, V], clave K, desde *K, hasta *K) bool {
	return (desde == nil || abb.cmp(clave, *desde) >= 0) && (hasta == nil || abb.cmp(clave, *hasta) <= 0)
}

func swap[K comparable, V any](nodo1, nodo2 *nodoAbb[K, V]) {
	nodo1.clave, nodo1.dato, nodo2.clave, nodo2.dato = nodo2.clave, nodo2.dato, nodo1.clave, nodo1.dato
}

func buscarDerecha[K comparable, V any](ant, act *nodoAbb[K, V]) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	if act.derecho == nil {
		return ant, act
	}
	return buscarDerecha(act, act.derecho)
}

func apilarAIzquierda[K comparable, V any](abb *abb[K, V], pila pila.Pila[*nodoAbb[K, V]], act *nodoAbb[K, V], desde *K, hasta *K) {
	if act == nil {
		return
	}
	if estaEnRango(abb, act.clave, desde, hasta) {
		pila.Apilar(act)
		apilarAIzquierda(abb, pila, act.izquierdo, desde, hasta)
	} else if desde != nil && abb.cmp(act.clave, *desde) < 0 {
		apilarAIzquierda(abb, pila, act.derecho, desde, hasta)
	} else if hasta != nil && abb.cmp(act.clave, *hasta) > 0 {
		apilarAIzquierda(abb, pila, act.izquierdo, desde, hasta)
	}
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{nil, 0, funcion_cmp}
}

func crearNodo[K comparable, V any](clave K, dato V, izq *nodoAbb[K, V], der *nodoAbb[K, V]) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{izq, der, clave, dato}
}
