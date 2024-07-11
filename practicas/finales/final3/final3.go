package main 

import (
	"fmt"
	"strconv"
	l "tdas/lista"
)

/*Dado un arreglo de n enteros positivos que casi representa una progresión aritmética creciente 
(es una progresión aritmética a la que le falta un elemento), 
implementar un algoritmo que devuelva el elemento faltante de manera eficiente
(complejidad logarítmica). 
Se puede suponer que el arreglo tiene al menos 4 elementos. 
Justificar la complejidad del algoritmo implementado. Por ejemplo, si la sucesión es [5, 8, 14, 17, 20, 23] 
tiene que devolver 11.*/

func progresionAritmetica(arr []int) int {
	var cdad int
	c1 := arr[1] - arr[0]
	c2 := arr[2] - arr[1]
	c3 := arr[3] - arr[2]

	if c1 == c2 || c1 == c3 {
		cdad = c1
	} else if c2 == c3 {
		cdad = c2
	}
	return buscarFaltante(arr, cdad, 0, len(arr))
}

func buscarFaltante(arr []int, cdad, ini, fin int) int {
	if ini == fin {
		return arr[ini]+cdad
	}

	m := (ini+fin)/2

	if arr[m] - arr[m-1] != cdad {
		return arr[m-1] + cdad
	}

	if arr[m] == arr[ini]+cdad*(m-ini) {
		return buscarFaltante(arr, cdad, m, fin)
	}
	return buscarFaltante(arr, cdad, ini, m)
}

/*La complejidad de este algoritmo es O(log n) porque es un algortimo recursivo con la ecuacion de recurrencia de la forma
	T(n) = 1T(n/2) + O(n°) 
	entonces A = 1, B = 2 y C= 0
	por el teorema maestro podemos ver que como log(1) == 0 entonces
	la complejidad del algoritmo es O(n°log(n)) == O(log(n))
	 
*/


/*Una lista que representa un número es una lista en la que en cada lugar tiene un dígito, y dicho número se lee
lugar a lugar. 
Por ejemplo, la lista 1 -> 8 -> 1 -> 2 representa al número 1812. 
Implementar una función que reciba dos listas que represeten números, y devuelva una lista que represente a la 
suma de estos. 
Por ejemplo, si recibe las listas 9 -> 1 -> 5 y 9 -> 6 debe devolver la lista 1 -> 0 -> 1 -> 1. 
Las listas recibidas por parámetro no deben verse modificadas. 
Indicar y justificar la complejidad del algoritmo implementado.*/


func sumarListas(l1, l2 l.Lista[int]) l.Lista[int] {
	var n1 string
	var n2 string

	for i := l1.Iterador(); i.HaySiguiente(); i.Siguiente() {
		n1 = n1 + strconv.Itoa(i.VerActual())
	}

	for i := l2.Iterador(); i.HaySiguiente(); i.Siguiente() {
		n2 = n2 + strconv.Itoa(i.VerActual())
	}

	num1, _ := strconv.Atoi(n1)
	num2, _ := strconv.Atoi(n2)

	suma := strconv.Itoa(num1+num2)
	listaSuma := l.CrearListaEnlazada[int]()
	for _, d := range suma {
		digit, _ := strconv.Atoi(string(d))
		listaSuma.InsertarUltimo(digit)
	}
	return listaSuma
}

/*La complejidad del algoritmo es O(n+m) siendo n el largo de la lista 1 y m el largo de la lista 2*/

func main() {
	pro := []int{5, 8, 14, 17, 20, 23}

	fmt.Println(progresionAritmetica(pro))

	n1 := []int{9,1,5}
	l1 := l.CrearListaEnlazada[int]()
	for _, n := range n1 {
		l1.InsertarUltimo(n)
	}
	n2 := []int{9, 6}
	l2 := l.CrearListaEnlazada[int]()
	for _, n := range n2 {
		l2.InsertarUltimo(n)
	}

	lista := sumarListas(l1, l2)
	for i := lista.Iterador(); i.HaySiguiente(); i.Siguiente() {
		fmt.Print(i.VerActual())
	}
}