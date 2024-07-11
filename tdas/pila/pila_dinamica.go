package pila

/* Definición del struct pila proporcionado por la cátedra. */

const (
	tamañoInicial = 10

	factorRedimension = 2

	escalaComparativa = 4
)

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	datos := make([]T, tamañoInicial)
	pila.datos = datos
	return pila
}

func (pila *pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func (pila *pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	} else {
		return pila.datos[pila.cantidad-1]
	}
}

func (pila *pilaDinamica[T]) Apilar(elem T) {
	if pila.cantidad == cap(pila.datos) {
		nuevaCapacidad := (cap(pila.datos)) * factorRedimension
		pila.redimension(nuevaCapacidad)
	}
	pila.datos[pila.cantidad] = elem
	pila.cantidad++
}

func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	dato := pila.datos[pila.cantidad-1]
	pila.cantidad--
	if pila.cantidad*escalaComparativa <= cap(pila.datos) {
		nuevaCapacidad := cap(pila.datos) / factorRedimension
		if nuevaCapacidad >= tamañoInicial {
			pila.redimension(nuevaCapacidad)
		} else {
			nuevaCapacidad = tamañoInicial
			pila.redimension(nuevaCapacidad)
		}
	}
	return dato
}

func (pila *pilaDinamica[T]) redimension(capacidad int) {
	nuevosDatos := make([]T, capacidad)
	copy(nuevosDatos, pila.datos)
	pila.datos = nuevosDatos
}
