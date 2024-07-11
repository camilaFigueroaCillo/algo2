package main

import ("fmt"
	"math"
	d "tdas/diccionario"
)

/*Implementar en Go una primitiva Invertir() para el Heap, que haga que el heap se comporte con la función de
comparación contraria a la que venía utilizando hasta ese momento. 
El heap debe quedar en estado correcto para que las operaciones siguientes sean válidas considerando la función 
de comparación, ahora invertida. No se puede modificar la estructura del heap para implementar esta primitiva. 
Es decir, La implementación actual que tienen del heap debería poder trabajar con esta primitiva, sin tener que 
modificar ninguna primitiva ni función auxiliar. Indicar y justificar la complejidad del algoritmo implementado. 
Bajo estas condiciones, Si se llamara k veces a esta nueva primitiva, ¿podría esto afectar a la complejidad de 
otras primitivas, como Encolar y Desencolar? Justificar.

type funcCmp[T any] func(T, T) int

type colaPrioridad[T any] struct {
	datos    []T
	cantidad int
	cmp      funcCmp[T]
}

func (heap *colaPrioridad[T]) Invertir() {
	f := func (a, b T) int {return -heap.cmp(a, b)}
	heap.cmp = f
	heap.heapify(heap.datos, heap.cantidad, heap.cmp)
}*/

//Debido a que crear una funcion que compara es O(1), La complejidad de la primitiva es la complejidad de heapify, que es O(n)
//SI se llama k vees a invertir y despues de esas k llamadas se realizan las operaciones de encolar y desencolar
//Entonces la complejidad seria O(k*n + c*logn) siendo c la cdad de llamados a encolar y desencolar 
//Entocnes la complejidad final no se veria afectada porque c*log(n) > k*n por logn ser una funcion creciente
//Por lo que la complejidad final de las primitivas no se veria modificada y serias O(c*logn). Esto es porque encolar y desencolar
//No dependen de la operatoria 'Heapify' si no que dependen de upheap y downheap.

/*Se tiene un arreglo de n elementos, ordenado, cuyos valores van de 0 a log(n). Implementar un algoritmo que permita
obtener la frecuencia de todos los números entre 0 y log(n) en tiempo menor a O(n). Justificar el orden del algoritmo.*/

func frequency(arr []int) d.Diccionario[int, int] {
	var fin int
	log_n := int(math.Log(float64(len(arr))) / math.Log(float64(2)))
	hash := d.CrearHash[int, int]()
	for i := 0; i < log_n; i++ {
		ini := busquedaBinaria(arr, i, 0, len(arr))
		if ini == -1 || arr[ini] != i {
			hash.Guardar(i, 0)
			continue
		}
		fin = busquedaBinaria(arr, i+1, 0, len(arr))
		hash.Guardar(i, fin-ini)
	}
	hash.Guardar(log_n, len(arr)-fin)
	return hash
}

func busquedaBinaria(arr []int, elem, ini, fin int) int {
	if ini > fin {
		return -1
	}
	if ini == fin && arr[ini] == elem{
		return ini
	}
	medio := (ini+fin)/2
	if arr[medio] > elem && arr[medio-1] < elem {
		return medio
	}
	if arr[medio] == elem && arr[medio-1] < arr[medio] {
		return medio
	}
	if arr[medio] < elem {
		return busquedaBinaria(arr, elem, medio+1, fin)
	}
	return busquedaBinaria(arr, elem, ini, medio-1)
}


/*La complejidad de este algoritmo es O(log(n) * (2 log(n)) ), lo que seria O(2log²(n)) que por propiedad de logaritmos
nos quedaria O(4log(n)) y como la notación desprecia las constantes entonces la complejidad del algoritmo es:
O(log(n))
*/

/*Bárbara está trabajando para una marca de ropa.
Está analizando el inventario, y nota que los proveedores le trajeron los guantes por unidades en vez de a pares,
como sería lo esperable. El problema es que ahora ni siquiera sabe si van a poder venderlos todos. 
Cuenta ahora con un arreglo de guantes, y necesita saber si para cada color puede formar
pares de guantes para poder venderlos (asumir que son todos de talle único). 
Implementar una función que reciba un arreglo de guantes, y devuelva true si no queda ningún guante suelto 
sin formar pareja, o false en caso contrario.
No se pueden juntar guantes de colores diferentes. Suponer que cada guante es un struct con un campo color, 
que es un enumerativo definido en algún lado. 
Indicar y justificar la complejidad de la función implementada. 
Si los guantes pudieran ser de diferentes talles, ¿cómo modificarías el algoritmo implementado para que resuelva 
el problema (para vender una pareja de guantes deben coincidir en talle y color)? 
¿Cambiaría la complejidad del algoritmo?*/

type Guante struct {
	color int
}


func paresGuantes(arr []Guante) bool {
	hash := d.CrearHash[int, int]()
	for _, guante := range arr {
		if hash.Pertenece(guante.color) {
			v := hash.Obtener(guante.color)
			hash.Guardar(guante.color, v+1)
		} else {
			hash.Guardar(guante.color, 1)
		}
	}

	for i := hash.Iterador(); i.HaySiguiente(); i.Siguiente() {
		_, cdad := i.VerActual()
			if cdad % 2 != 0 {
				return false
			}
		
	}
	return true
}

/*La complejidad del algoritmo es O(n) porque lo primero que hago es recorrer todo el array de guantes y guardarme los elementos en
el hash que es una operacion constante, y depsues recorro el diccionario que no va atener mas claves que el largo del array
y verifico la cdad de guantes de ese color que es una operacion constante, por lo tanto toda la funcion tiene una complejidad
lineal en funcion de la cdad de guantes que hay en el arreglo*/

/*En clase vimos que se puede implementar un heap con un árbol izquierdista, o su representación equivalente en 
arreglo. Esta última, siendo mucho más sencilla de implementar. 
¿Por qué no implementamos también el Árbol Binario de Búsqueda con una representación en arreglo, 
en vez de implementarlo con, valga la redundancia, árboles?*/

/* Primero, sabemos que implementar el heap con un arreglo funciona porque 1) es un arbol izquierdista y 2) dado un 
elemento del heap, sabemos que su hijo izquierdo es un heap, y su hijo derecho es un heap. POr lo que la formula
para encontrar los hijos en el arreglo funcionan. 
Ahora que pasa con el ABB, El abb tiene una propiedad similar pero no la misma al heap, que es: "mi hijo izquierdo es menor a mi
y mi hijo derecho es mayor a mi."

Por un lado, tratar de mantener un arbol izquierdista con una comparacion que depende del usuario, seguramente no sea
muy simple, mas allá de eso, el heap trabaja con las operatorias 'upheap' y 'downheap' donde upheap sube el elemento
hasta que encuentre un padre mayor y downheap baja el elemento hasta que no tenga hijos o su hijo sea menor
con el abb esto no seria posible porque podria terminar en un lugar donde no corresponderia y nos arruinaria el orden

*/



func main() {
	array := []int{0,0,0,2,2,2,2,3}
	h := frequency(array)
	for i:= h.Iterador(); i.HaySiguiente(); i.Siguiente() {
		fmt.Println(i.VerActual())
	}

}


