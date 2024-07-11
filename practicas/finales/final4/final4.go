package main

import (
	"fmt"
	"strconv"
	h "tdas/cola_prioridad"
)

/*Definimos a la reducción de una secuencia de números como la operación de sacar los dos elementos más pequeños
de la misma, y volver a guardar el resultado de 2 * mínimo + segundo_mínimo. 
Esto sólo puede hacerse si hay al menos dos elementos. 
Ejemplo: [1, 7, 2, 3] -> [7, 3, 4]. 
De esta forma, la secuencia queda con un elemento menos.
Implementar un algoritmo que reciba un arreglo y 
devuelva el único valor que quedaría en el arreglo si aplicaramos dicha reducción hasta que quede un único 
elemento en el arreglo. 
Indicar y justificar la complejidad del algoritmo.*/

func reduccion(arr []int) int {
	heap := h.CrearHeapArr[int](arr, func(a, b int) int {return b-a}) 
	for heap.Cantidad() >= 2 {
		a := heap.Desencolar()
		b := heap.Desencolar()
		heap.Encolar(2*a+b)
	}
	return heap.Desencolar()
}

//La complejidad del algoritmo es O(nlogn), siendo n la cdad de elementos del array

func main() {
	arr := []int{1, 7, 2, 3}
	fmt.Println(reduccion(arr))
}


/*Así como en Go tenemos slices de arreglos, nos gustaría tener un slice sobre una lista enlazada. 
Es decir, poder tener “una porción” de dicha lista. 
Por ejemplo, si yo hago slice := lista.Slice(5, 9) y luego hago slice.VerPrimero(),
dicha operación me devuelva el equivalente a haber iterado hasta la posición 5 en la original, y 
devolver ese elemento (pero lo hace en tiempo constante). 
También podemos usar las primitivas que modifican la lista, obtener el largo (en el ejemplo sería 4), 
usar el iterador interno y externo sobre dicho slice, así como crear otro slice para ese slice. 
Es decir, Slice debe implementar Lista. 
Explicar cómo implementarías dicha estructura slice (campos, y cómo serían las primitivas) 
para que esto funcione como se indica, de forma eficiente. 
No deben hacerse menciones a modificar la lista enlazada (salvo, por supuesto, el agregado de la primitiva Slice)
o su iterador externo, ya que esto no es ni debe ser necesario. 
No es necesario implementar, si mencionar los puntos claves sobre referencias que se tendrían,
cómo funcionarían las primitivas que modifiquen la lista (tal que esta quede correctamente actualizada) 
y cómo sería la iteración externa e interna. 
Se puede asumir que, tal como con los iteradores, no se pueden tener dos slices con uno
modificando la lista, y lo mismo con un slice y un iterador, ni tampoco usar las primitivas de 
modificación de la lista mientras haya un slice en uso.*/


/*Implementar un algoritmo que permita ordenar canciones, por su respectivo año de publicación, 
de forma eficiente, tomando en cuenta canciones desde el año 1800. 
Indicar y justificar la complejidad del algoritmo (no se toman como válidas respuestas parciales). 
Si quisiéramos considerar hasta incluso las canciones que escuchaban los dinosaurios, ¿el algoritmo propusto 
seguiría siendo eficiente? Si lo es, justificar. Si no lo es, mencionar otro algoritmo que sea mejor.*/

type Cancion struct {
	nombre string
	año int
}


func ordenarCanciones(arr []Cancion) {
	for i := 1; i <= 1000; i = i*10 {
		countingSort(arr, func(c Cancion) int {return c.año}, 4, i)
	}
}

func countingSort(arr []Cancion, criterio func(a Cancion) int, rango, digito int) {
	freq := make([]int, rango)
	sumasAcum := make([]int, rango)
	res := make([]Cancion, len(arr))

	for _, e := range arr {
		valor := criterio(e)/digito % 10
		freq[valor]++
	} 

	for i:=1; i < len(freq); i++ {
		sumasAcum[i] = freq[i-1] + sumasAcum[i-1]
	}

	for _, e := range arr {
		valor := criterio(e)/digito%10
		pos := sumasAcum[valor]
		res[pos] = e
		sumasAcum[valor]++
	}

	copy(arr, res)
}

/*Para ordenar implemente un radix por digito de los años, del menos significativo al mas significativo, 
es decir de las unidades a las unidades de mil.
La complejidad de este algoritmo es O(4*n) lo que terminaria siendo O(n)
Si quisiesemos ordenar canciones desde la era de los dinosaurios, tambien funcionaria el radixsort por digito
la unica diferencia que haria es usar un bucketSort para separar entre los años antes de cristo y despues de cristo
entonces ordeno un balde luego otro y luego junto los resultados
*/


