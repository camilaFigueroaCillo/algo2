package cola

type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

func nodoCrear[T any](dato T, sig *nodoCola[T]) *nodoCola[T] {
	nodo := new(nodoCola[T])
	nodo.dato = dato
	nodo.prox = sig
	return nodo
}

func CrearColaEnlazada[T any]() Cola[T] {
	cola := new(colaEnlazada[T])
	return cola
}

func (cola *colaEnlazada[T]) EstaVacia() bool {
	return cola.primero == nil
}

func (cola *colaEnlazada[T]) VerPrimero() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	return cola.primero.dato
}

func (cola *colaEnlazada[T]) Encolar(elem T) {
	nuevoNodo := nodoCrear(elem, nil)
	if cola.primero == nil {
		cola.primero = nuevoNodo
	} else {
		cola.ultimo.prox = nuevoNodo
	}
	cola.ultimo = nuevoNodo
}

func (cola *colaEnlazada[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	primero := cola.primero.dato
	if cola.primero == cola.ultimo {
		cola.ultimo = nil
	}
	cola.primero = cola.primero.prox
	return primero
}
