package lista

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type iterListaEnlazada[T any] struct {
	lista    *listaEnlazada[T]
	anterior *nodoLista[T]
	actual   *nodoLista[T]
}

func CrearListaEnlazada[T any]() Lista[T] {
	return new(listaEnlazada[T])
}

func (listaEnlazada *listaEnlazada[T]) EstaVacia() bool {
	return listaEnlazada.largo == 0 && listaEnlazada.primero == nil && listaEnlazada.ultimo == nil
}

func (listaEnlazada *listaEnlazada[T]) InsertarPrimero(dato T) {
	nuevoNodo := nodoCrear(dato, listaEnlazada.primero)
	if listaEnlazada.EstaVacia() {
		listaEnlazada.ultimo = nuevoNodo
	}
	listaEnlazada.primero = nuevoNodo
	listaEnlazada.largo++
}

func (listaEnlazada *listaEnlazada[T]) InsertarUltimo(dato T) {
	nuevoNodo := nodoCrear(dato, nil)
	if listaEnlazada.EstaVacia() {
		listaEnlazada.primero = nuevoNodo
	} else {
		listaEnlazada.ultimo.siguiente = nuevoNodo
	}
	listaEnlazada.ultimo = nuevoNodo
	listaEnlazada.largo++
}

func (listaEnlazada *listaEnlazada[T]) BorrarPrimero() T {
	if listaEnlazada.EstaVacia() {
		panic("La lista esta vacia")
	}
	dato := listaEnlazada.primero.dato
	if listaEnlazada.largo == 1 {
		listaEnlazada.ultimo = listaEnlazada.ultimo.siguiente
	}
	listaEnlazada.primero = listaEnlazada.primero.siguiente
	listaEnlazada.largo--
	return dato
}

func (listaEnlazada *listaEnlazada[T]) VerPrimero() T {
	if listaEnlazada.EstaVacia() {
		panic("La lista esta vacia")
	}
	return listaEnlazada.primero.dato
}

func (listaEnlazada *listaEnlazada[T]) VerUltimo() T {
	if listaEnlazada.EstaVacia() {
		panic("La lista esta vacia")
	}
	return listaEnlazada.ultimo.dato
}

func (listaEnlazada *listaEnlazada[T]) Largo() int {
	return listaEnlazada.largo
}

func (listaEnlazada *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	act := listaEnlazada.primero
	for act != nil {
		if visitar(act.dato) {
			act = act.siguiente
		} else {
			break
		}
	}
}

func (listaEnlazada *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterListaEnlazada[T]{listaEnlazada, nil, listaEnlazada.primero}
}

func (iterador *iterListaEnlazada[T]) HaySiguiente() bool {
	return iterador.actual != nil
}

func (iterador *iterListaEnlazada[T]) VerActual() T {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iterador.actual.dato
}

func (iterador *iterListaEnlazada[T]) Siguiente() {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iterador.anterior = iterador.actual
	iterador.actual = iterador.actual.siguiente
}

func (iterador *iterListaEnlazada[T]) Insertar(dato T) {
	nuevoNodo := nodoCrear(dato, iterador.actual)
	if iterador.actual == iterador.lista.primero {
		iterador.lista.primero = nuevoNodo
	} else {
		iterador.anterior.siguiente = nuevoNodo
	}
	if !iterador.HaySiguiente() {
		iterador.lista.ultimo = nuevoNodo
	}
	iterador.actual = nuevoNodo
	iterador.lista.largo++
}

func (iterador *iterListaEnlazada[T]) Borrar() T {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	dato := iterador.actual.dato
	if iterador.lista.primero == iterador.actual {
		iterador.lista.primero = iterador.actual.siguiente
	} else {
		iterador.anterior.siguiente = iterador.actual.siguiente
	}
	if iterador.lista.ultimo == iterador.actual {
		iterador.lista.ultimo = iterador.anterior
	}
	iterador.actual = iterador.actual.siguiente
	iterador.lista.largo--
	return dato
}

func nodoCrear[T any](dato T, sig *nodoLista[T]) *nodoLista[T] {
	return &nodoLista[T]{dato, sig}
}
