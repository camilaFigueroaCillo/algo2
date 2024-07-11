package cola_prioridad

const (
	tamañoInicial     = 10
	factorRedimension = 2
	escalaComparativa = 4
	posInvalida       = -1
)

type funcCmp[T any] func(T, T) int

type colaPrioridad[T any] struct {
	datos    []T
	cantidad int
	cmp      funcCmp[T]
}

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	arr := make([]T, tamañoInicial)
	return &colaPrioridad[T]{arr, 0, funcion_cmp}
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	if len(arreglo) == 0 {
		return CrearHeap(funcion_cmp)
	}
	datos := make([]T, len(arreglo))
	heap := &colaPrioridad[T]{datos, 0, funcion_cmp}
	copy(heap.datos, arreglo)
	heap.cantidad = len(arreglo)
	heap.cmp = funcion_cmp
	heapify(heap.datos, heap.cantidad, heap.cmp)
	return heap
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	heapify(elementos, len(elementos), funcion_cmp)
	heapsort(elementos, funcion_cmp, len(elementos)-1)
}

func (heap *colaPrioridad[T]) EstaVacia() bool {
	return heap.cantidad == 0
}

func (heap *colaPrioridad[T]) Encolar(dato T) {

	if heap.cantidad == cap(heap.datos) {
		nuevaCapacidad := cap(heap.datos) * factorRedimension
		heap.redimension(nuevaCapacidad)
	}

	heap.datos[heap.cantidad] = dato

	padre := obtenerPadre(heap.cantidad)

	upheap(heap.datos, heap.cantidad, padre, heap.cmp)

	heap.cantidad++
}

func (heap *colaPrioridad[T]) Desencolar() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}

	heap.cantidad--

	if heap.EstaVacia() {
		return heap.datos[0]
	}

	if heap.cantidad*escalaComparativa <= cap(heap.datos) {
		nuevaCapacidad := cap(heap.datos) / factorRedimension
		if nuevaCapacidad > tamañoInicial {
			heap.redimension(nuevaCapacidad)
		}

	}

	dato := heap.datos[0]
	swap(heap.datos, 0, heap.cantidad)

	mayor := obtenerMayorHijo(heap.datos, heap.cantidad, heap.cmp, 0)

	downheap(heap.datos, mayor, 0, heap.cmp, heap.cantidad)

	return dato
}

func (heap *colaPrioridad[T]) Cantidad() int {
	return heap.cantidad
}

func (heap *colaPrioridad[T]) VerMax() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	return heap.datos[0]
}

func upheap[T any](datos []T, hijo, padre int, cmp func(a, b T) int) {

	if padre == posInvalida || hijo == 0 || cmp(datos[padre], datos[hijo]) >= 0 {
		return
	}

	swap(datos, hijo, padre)

	nuevoPadre := obtenerPadre(padre)

	upheap(datos, padre, nuevoPadre, cmp)
}

func downheap[T any](datos []T, hijo, padre int, cmp func(a, b T) int, cantidad int) {

	if hijo == posInvalida || cmp(datos[padre], datos[hijo]) >= 0 {
		return
	}

	swap(datos, hijo, padre)

	mayor := obtenerMayorHijo(datos, cantidad, cmp, hijo)

	downheap(datos, mayor, hijo, cmp, cantidad)
}

func maximo[T any](hijo1, hijo2 int, datos []T, cmp func(a, b T) int) int {
	if hijo2 == posInvalida || hijo1 != posInvalida && cmp(datos[hijo1], datos[hijo2]) > 0 {
		return hijo1
	} else {
		return hijo2
	}
}

func (heap *colaPrioridad[T]) redimension(capacidad int) {
	nuevosDatos := make([]T, capacidad)
	copy(nuevosDatos, heap.datos)
	heap.datos = nuevosDatos
}

func swap[T any](arr []T, hijo, padre int) {
	arr[hijo], arr[padre] = arr[padre], arr[hijo]
}

func heapify[T any](arr []T, cant int, comparacion func(a, b T) int) {
	for i := cant - 1; i >= 0; i-- {
		mayor := obtenerMayorHijo(arr, cant, comparacion, i)
		downheap(arr, mayor, i, comparacion, len(arr))
	}
}

func obtenerPadre(posHijo int) int {
	if posHijo == 0 {
		return posInvalida
	}
	return (posHijo - 1) / 2
}

func obtenerHijos[T any](posPadre int, cantidad int) (int, int) {
	posDer := 2*posPadre + 2
	posIzq := 2*posPadre + 1
	if posIzq >= cantidad {
		posIzq = posInvalida
	}
	if posDer >= cantidad {
		posDer = posInvalida
	}
	return posIzq, posDer
}

func heapsort[T any](elementos []T, cmp func(a, b T) int, ultimo int) {
	if len(elementos) == 0 || ultimo == 0 {
		return
	}
	swap(elementos, 0, ultimo)
	mayor := obtenerMayorHijo(elementos, ultimo, cmp, 0)
	downheap(elementos, mayor, 0, cmp, ultimo)
	heapsort(elementos, cmp, ultimo-1)
}

func obtenerMayorHijo[T any](arr []T, cant int, comparacion func(a, b T) int, padre int) int {
	izq, der := obtenerHijos[T](padre, cant)
	return maximo(izq, der, arr, comparacion)
}
