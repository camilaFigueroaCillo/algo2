package main

import (
	"fmt"
	"strconv"
)

/*
Se tiene un arreglo con las distancias en años luz entre la Tierra y distintos cuerpos celestes en la Via Lactea y otras
galaxias ceranas descubiertos hasta el momento. Debido a un error cometido por el sysadmin del Planetario, estas
distancias hoy están completamente desordenadas.
Esta persona necesita volver a ordenar los registros antes de que
descubran el error para no perder su trabajo. Como no cursó Algoritmos 2 no sabe cómo hacer, por lo que nos pide
ayuda a nosotros.
Diseñar un algoritmo de ordenamiento que pueda ordenar un arreglo de CuerpoCeleste donde
cada uno tiene un campo distancia int que funcione lo más rápido posible. Indicar y justificar el orden del
algoritmo.
De los datos, se sabe lo siguiente:
• Hay alrededor de 500 mil de registros
• Las distancias son muy dispares, yendo desde unos pocos años luz (como Alfa Centauri, la estrella mas cercana a
la Tierra que está a 4 años luz) a un poco más de 20 millones (como ADS 7251, una de las estrellas que conforman
la constelación Osa Mayor).
*/

type CuerpoCeleste struct {
	n string
	dist int
}

func ordenarCCelestes(arr []CuerpoCeleste) {
	//counting(arreglo, rango, digito)
	countingSort(arr, 10, 1) //unidad
	countingSort(arr, 10, 10) //decena
	countingSort(arr, 10, 100) //centena
	countingSort(arr, 10, 1000) // unidad de mil
	countingSort(arr, 10, 10000) //decena de mil
	countingSort(arr, 10, 100000) //centenda de mil
	countingSort(arr, 10, 1000000) //unidad de millon
	countingSort(arr, 10, 10000000) //decena de millon
}


func countingSort(arr []CuerpoCeleste, rango, digito int) {
	frecuencias := make([]int, rango)
	sumasAcum := make([]int, rango)
	res := make([]CuerpoCeleste, len(arr))

	for _, e := range arr {
		valor := (e.dist / digito) % 10
		frecuencias[valor]++
	}

	for i := 1; i < len(frecuencias); i++ {
		sumasAcum[i] = sumasAcum[i-1]+frecuencias[i-1]
	}

	for _, e := range arr {
		valor := (e.dist/digito)%10
		pos := sumasAcum[valor]
		res[pos] = e
		sumasAcum[valor]++
	}

	copy(arr, res)
}

/*El Wordle es un popular juego que nos desafía a encontrar la palabra correcta cada día. 
Una característica de este juego es que todas las palabras usan las letras desde la A a la Z, y siempre son 
palabras de 5 letras. 
Implementar un algoritmo de ordenamiento lineal que, dado un arreglo de todas las palabras que pueden ser 
solución de este juego, las ordene alfabéticamente. 
En caso de usar ordenamientos auxiliares, todos deben ser implementados. 
Indicar y justificar la complejidad del algoritmo (no se aceptarán resultados parciales).*/


func ordenarAlfabeticamente(arr []string) {
	countingSortW(arr, 27, 5)
	countingSortW(arr, 27, 4)
	countingSortW(arr, 27, 3)
	countingSortW(arr, 27, 2)
	countingSortW(arr, 27, 1)

}

func countingSortW(arr []string, rango, digito int) {
	frecuencias := make([]int, rango)
	sumasAcum := make([]int, rango)
	res := make([]string, len(arr))

	for _, e := range arr {
		valor := (e[digito] - '0')
		frecuencias[valor]++
	}

	for i := 1; i < len(frecuencias); i++ {
		sumasAcum[i] = sumasAcum[i-1]+frecuencias[i-1]
	}

	for _, e := range arr {
		valor := (e[digito] - '0')
		pos := sumasAcum[valor]
		res[pos] = e
		sumasAcum[valor]++
	}

	copy(arr, res)
}


/*Llega el primer gran torneo de Algortnite. 
Las 1000 personas más hábiles en este videojuego participarán en una cantidad de rondas, cada una con una 
batalla todos contra todos. 
Para cada ronda se registra el número de participante (de 0 a 999) y apodo que ha ganado. 
Dado un arreglo que registra quién ganó cada ronda, desde la primera hasta la última, se desea obtener un 
arreglo ordenado por el número de participante que ganó cada ronda. 
Por ejemplo, si gana 1 vez el participante 60 (algrtmz), tres veces seguidas el participante 50 (mb) y una vez 
el 19 (algorw), 
la entrada es [(60,algrtmz), (50,mb), (50,mb), (50,mb), (19,algorw)]  
el arreglo ordenado será
[(19,algorw), (50,mb), (50,mb), (50,mb), (60,algrtmz)]
*/

func ordenarParticipantes(arr [][]string){
	counting(arr, 10, 1)
	counting(arr, 10, 10)
	counting(arr, 10, 100)
}

func counting(arr [][]string, rango, digito int) {

	frecuencias := make([]int, rango)
	sumasAcum := make([]int, rango)
	res := make([][]string, len(arr))

	for _, e := range arr {
		// e = ["int", string]
		d, _ := strconv.Atoi(e[0])
		valor := (d/digito) % 10
		frecuencias[valor]++
	}

	for i := 1; i < len(frecuencias); i++ {
		sumasAcum[i] = sumasAcum[i-1]+frecuencias[i-1]
	}

	for _, e := range arr {
		d, _ := strconv.Atoi(e[0])
		valor := (d/digito) % 10
		pos := sumasAcum[valor]
		res[pos] = e
		sumasAcum[valor]++
	}
	copy(arr, res)
}



func main() {
	//p := [][]string{["60","algrtmz"], ["50","mb"], ["50","mb"], ["50","mb"], ["19","algorw"]}
	//fmt.Println(p)
	//ordenarParticipantes(p)
	return
}