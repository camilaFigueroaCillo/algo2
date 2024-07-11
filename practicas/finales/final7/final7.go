package main

import ( 
	"fmt"
	h "tdas/cola_prioridad"
	d "tdas/diccionario"
	l "tdas/lista"
)

func buscarNoRepe(arr []int, ini, fin, elem int) int {
	if ini == fin {
		return arr[ini]
	}

	m := (ini+fin)/2

	if arr[m] != elem {
		return arr[m]
	}

	izq := buscarNoRepe(arr, ini, m-1, elem)
	der := buscarNoRepe(arr, m+1, fin, elem)

	if izq == elem {
		return der
	}
	return izq
}

func elementoRepetido(arr []int) int {
	var repe int
	if arr[0] == arr[1] || arr[0] == arr[2] {
		repe = arr[0]
	} else { 
		repe = arr[1]
	}
	return buscarNoRepe(arr, 0, len(arr)-1, repe)
}

func main() {
	
	
	fmt.Println(sonAnagramas("hola", "aloh"))
	fmt.Println(sonAnagramas("hoa", "aloh"))
	fmt.Println(sonAnagramas("holi", "aloh"))

}

/*
Implementar una función que reciba un arreglo A de n enteros y un número k y devuelva un nuevo arreglo en el 
que para cada posición i de dicho arreglo, contenga el resultado de la multiplicación de los primeros k máximos 
del arreglo A entre las posición [0;i] (incluyendo a i). 
Las primeras k − 1 posiciones del arreglo a devolver deben tener como valor -1.
Por ejemplo, para el arreglo [1, 5, 3, 4, 2, 8] y k = 3, el resultado debe ser [-1, -1, 15, 60, 60, 160]. Indicar
y justificar la complejidad del algoritmo implementado.
*/

func maxArray(arr []int, k int) []int {
	heap := h.CrearHeapArr[int](arr[:k], func(a, b int) int {return b-a})
	res := make([]int, len(arr))
	producto := 1
	for i := 0; i <= k-1; i++ {
		if i < k-1 {
			res[i] = -1
		}
		producto = producto * arr[i]
		if i == k-1 {
			res[i] = producto
		}
	}
	for i := k; i < len(arr); i++ {
		menor := heap.VerMax()
		if arr[i] > menor {
			heap.Encolar(arr[i])
			heap.Desencolar()
			producto = (producto / menor)*arr[i]
		}
		res[i] = producto
	}
	return res
}

/*Implementar un algoritmo que reciba dos cadenas (strings) y determine si son anagramas entre sí. Indicar y justificar
la complejidad del algortmo implementado.*/

func sonAnagramas(c1, c2 string) bool {
	if len(c1) != len(c2) {
		return false
	}
	hash := d.CrearHash[string, int]()
	for i, char := range c1 {
		if hash.Pertenece(string(char)) {
			v := hash.Obtener(string(char))
			hash.Guardar(string(char), v+1)
		} 
		if !hash.Pertenece(string(char)) {
			hash.Guardar(string(char), 1)
		}
		if hash.Pertenece(string(c2[i])) {
			v := hash.Obtener(string(c2[i]))
			hash.Guardar(string(c2[i]), v+1)
		}
		if !hash.Pertenece(string(c2[i])) {
			hash.Guardar(string(c2[i]), 1)
		}
	}

	for i := hash.Iterador(); i.HaySiguiente(); i.Siguiente() {
		_, v := i.VerActual()
		if v % 2 != 0 {
			return false
		}
	}
	return true
}

//La complejidad de este algoritmo es O(n), siendo n la cdad de caracteres de la palabra

/*Implementar una primitiva de árbol binario de búsqueda que devuelva un diccionario en el cual las claves 
sean los niveles (int) y los datos sean listas de todos las claves del ABB que se encuentran en dicho nivel. 
Indicar y justificar la complejidad del algoritmo implementado.*/

func (abb *Abb[K, V]) dictClaves() d.Diccionario[int, l.Lista[K]] {
	hash := d.CrearHash[int, l.Lista[K]]()
	altura := abb.Altura()
	for i := 0; i <= altura; i++ {
		hash.Guardar(i, l.CrearListaEnlazada[K]())
	}
	abb.llenarDiccionario(hash, 0, abb.raiz)
	return hash
}

func (abb *Abb[K, V]) llenarDiccionario(hash d.Diccionario[int, l.Lista[K]], nivel int, nodo *nodoABB[K, V]) {
	if nodo == nil {
		return
	}
	abb.llenarDiccionario(hash, nivel+1, nodo.izquierdo)
	abb.llenarDiccionario(hash, nivel+1, nodo.derecho)
	lista := hash.Obtener(nivel)
	lista.InsertarUltimo(nodo.clave)
	hash.Guardar(nivel, lista)
}

func (abb *Abb[K,V]) Altura() int {
	if abb.raiz == nil {
		return 0
	}
	izq := abb.raiz.izquierdo.Altura()
	der := abb.raiz.derecho.Altura()
	return maximo(izq, der)+1
}

func maximo(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//La complejidad del algoritmo es O(N) siendo n la cdad de nodos del arbol porque recorro todos los nodos una vez yr ealizo
//Operaciones constantes