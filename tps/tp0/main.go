package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tp0/ejercicios"
)

func procesarArchivo(ruta string) []int {
	archivo, _ := os.Open(ruta)
	defer archivo.Close()
	s := bufio.NewScanner(archivo)
	vector := []int{}
	for s.Scan() {
		elemento, _ := strconv.Atoi(s.Text())
		vector = append(vector, elemento)
	}
	return vector
}

func imprimirVectorOrd(vector []int) {
	ejercicios.Seleccion(vector)
	for i := 0; i < len(vector); i++ {
		fmt.Println(vector[i])
	}
}

func main() {
	const ruta1 = "archivo1.in"
	const ruta2 = "archivo2.in"
	vector1 := procesarArchivo(ruta1)
	vector2 := procesarArchivo(ruta2)
	n := ejercicios.Comparar(vector1, vector2)
	if n == 1 {
		imprimirVectorOrd(vector1)
	} else {
		imprimirVectorOrd(vector2)
	}
}
