package main

import (
	"bufio"
	"fmt"
	"os"
	calculadora "tp1/calculadora"
)

func calcular(texto string) {
	cola := calculadora.EncolarOperacion(texto)
	pila := calculadora.ApilarResultados(cola)
	if !pila.EstaVacia() {
		total := pila.Desapilar()
		if pila.EstaVacia() {
			fmt.Println(total)
		} else {
			fmt.Println("ERROR")
		}
	} else {
		fmt.Println("ERROR")
	}
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		calcular(s.Text())
	}

}
