package main

import (
	"fmt"
	"math"
)

type Personas struct {
	nombre string
	fecha int
}

func fechasCumple(jugadores []Personas, fecha int) []Personas {
	if  len(jugadores) == 0 {
		return []Personas{}
	}
	i := busquedaBinaria(jugadores, fecha, 0, len(jugadores)-1)
	if jugadores[i].fecha > fecha {
		f := busquedaBinaria(jugadores, jugadores[i].fecha+1, 0, len(jugadores)-1)
		return jugadores[i:f]
	} else {
		f := busquedaBinaria(jugadores, fecha+1, 0, len(jugadores)-1)
		return jugadores[i:f]
	}
	
}

func busquedaBinaria(jugadores []Personas, fecha, ini, fin int) int {
	if jugadores[ini].fecha == fecha {
		return ini
	}
	
	medio := (ini+fin)/2

	if jugadores[medio].fecha < fecha && jugadores[medio+1].fecha > fecha {
		return medio+1
	}
	if jugadores[medio].fecha == fecha && jugadores[medio-1].fecha < fecha {
		return medio
	}
	if jugadores[medio].fecha == fecha || jugadores[medio].fecha > fecha {
		return busquedaBinaria(jugadores, fecha, ini, medio-1)
	}
	return busquedaBinaria(jugadores, fecha, medio+1, fin)
}

//Implementar un algoritmo que reciba un arreglo de enteros y determine el par cuya suma sea mayor

func sumaPar(arr []int) (int, int) {
	p1, p2, _ := parMayor(arr)
	return p1, p2
}

func parMayor(arr []int) (int, int, int) {
	if len(arr) == 1 {
		return arr[0], 0, arr[0]
	}
	if len(arr) == 2 {
		return arr[0], arr[1], arr[0]+arr[1]
	}
	
	m := len(arr)/2
	n1Izq, n2Izq, sumaI := parMayor(arr[:m+1])
	n1Der, n2Der, sumaD := parMayor(arr[m:])

	if sumaI > sumaD {
		return n1Izq, n2Izq, sumaI
	}
	return n1Der, n2Der, sumaD
}

/*Implementar un algoritmo que reciba un arreglo de enteros de tamaño n, ordenado descendente-
mente, y determine en tiempo O(log n) si existe algún valor i tal que arr[i] = - nˆ2. Justificar
la complejidad del algoritmo.*/

func arrPotencia(arr []int) int {
	return potenciaArr(arr, 0, len(arr)-1)
}

func potenciaArr(arr []int, i, f int) int {
	if i >= f {
		if arr[f] == -int(math.Pow(float64(len(arr)), float64(2))) {
			return f
		}
		return -1
	}

	m := (i+f)/2
	if arr[m] == -int(math.Pow(float64(len(arr)), float64(2))) {
		return m
	}

	if arr[m] < -int(math.Pow(float64(len(arr)), float64(2))) {
		return potenciaArr(arr, i, m-1)
	}
	return potenciaArr(arr, m+1, f)
}




func main(){
	a := Personas{"a", 29}
	b := Personas{"b", 29}
	e := Personas{"e", 30}
	f := Personas{"f", 30}
	g := Personas{"g", 30}
	h := Personas{"h", 30}
	i := Personas{"i", 31}
	j := Personas{"j", 31}
	k := Personas{"k", 31}
	l := Personas{"l", 31}
	m := Personas{"m", 32}
	n := Personas{"n", 32}
	o := Personas{"o", 32}
	p := Personas{"p", 32}

	jugadores_1 := []Personas{a, b, e, f, g, h, i, j, k, l, m, n, o, p}
	jugadores_2 := []Personas{a, b, i, j, k, l, m, n, o, p}
	jugadores_3 := []Personas{e, f, g, h, i, j, k, l, m, n, o, p}


	fmt.Println(fechasCumple(jugadores_1, 30))
	fmt.Println(fechasCumple(jugadores_2, 30))
	fmt.Println(fechasCumple(jugadores_3, 30))

	arr := []int{4,8,9,1,3,6,7,1,-81} //9 
	fmt.Println(sumaPar(arr))
	fmt.Println(arrPotencia(arr))

}






