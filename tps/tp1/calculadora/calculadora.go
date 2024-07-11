package calculadorapolaca

import (
	"fmt"
	"strconv"
	"strings"
	TDACola "tdas/cola"
	TDAPila "tdas/pila"
	Operaciones "tp1/operaciones"
)

const operadores = "+ - / * log sqrt ? ^"

func esOperador(op string, cadena_op string) bool {
	operadores := strings.Split(cadena_op, " ")
	for _, e := range operadores {
		if string(e) == op {
			return true
		}
	}
	return false
}

func esDigito(c string) bool {
	if c == "" || len(c) == 0 {
		return false
	}
	if len(c) > 1 && string(c[0]) == "-" {
		return esDigito(c[1:])
	}
	runes := []rune(c)
	for r := range runes {
		if !(runes[r] >= '0' && runes[r] <= '9') {
			return false
		}
	}
	return true
}

func EncolarOperacion(texto string) TDACola.Cola[string] {
	cola := TDACola.CrearColaEnlazada[string]()
	if texto == "" {
		return cola
	}
	caracteres := strings.Split(texto, " ")
	for _, valor := range caracteres {
		if esOperador(valor, operadores) || esDigito(valor) {
			cola.Encolar(string(valor))
		}
	}
	return cola
}

func evaluarOperador(pila TDAPila.Pila[int64], op string) error {
	var res int64
	if pila.EstaVacia() {
		return fmt.Errorf("pila vacía")
	}
	a := pila.Desapilar()
	if op == "sqrt" {
		if a < 0 {
			return fmt.Errorf("operación inválida")
		}
		res = Operaciones.RaizCuadrada(a)
		pila.Apilar(res)
	}

	if pila.EstaVacia() {
		return fmt.Errorf("pila vacía")
	}
	b := pila.Desapilar()

	switch op {
	case "+":
		res = Operaciones.Suma(b, a)
	case "-":
		res = Operaciones.Resta(b, a)
	case "*":
		res = Operaciones.Producto(b, a)
	case "/":
		if a == 0 {
			return fmt.Errorf("operación inválida")
		}
		res = Operaciones.Division(b, a)
	case "log":
		if a < 2 {
			return fmt.Errorf("operación inválida")
		}
		res = Operaciones.Logaritmo(b, a)
	case "^":
		if a < 0 {
			return fmt.Errorf("operación inválida")
		}
		res = Operaciones.Potencia(b, a)
	case "?":
		if pila.EstaVacia() {
			return fmt.Errorf("pila vacía")
		}
		c := pila.Desapilar()
		res = Operaciones.Ternario(c, b, a)
	}
	pila.Apilar(res)
	return nil
}

func ApilarResultados(cola TDACola.Cola[string]) TDAPila.Pila[int64] {
	pila := TDAPila.CrearPilaDinamica[int64]()
	for !cola.EstaVacia() {
		e := cola.Desencolar()
		if esOperador(e, operadores) {
			err := evaluarOperador(pila, e)
			if err != nil {
				break
			}
		} else {
			n, _ := strconv.ParseInt(e, 10, 64)
			pila.Apilar(n)
		}
	}
	return pila
}
