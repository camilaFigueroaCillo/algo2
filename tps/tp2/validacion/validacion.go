package validacion

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type operacion int

const (
	Agregar_archivo = operacion(iota)
	Ver_mas_visitados
	Ver_visitantes
	Operacion_invalida
)

func ValidarEntrada(texto string) ([]string, operacion, error) {
	if texto == "" {
		return ErrorValidacion([]string{})
	}
	linea := strings.Split(texto, " ")
	if len(linea) < 2 || len(linea) > 3 {
		return ErrorValidacion(linea)
	}
	if len(linea) == 2 {
		if linea[0] == "agregar_archivo" {
			return linea, Agregar_archivo, nil
		} else if linea[0] == "ver_mas_visitados" {
			return linea, Ver_mas_visitados, nil
		} else {
			return ErrorValidacion(linea)
		}
	} else {
		if linea[0] != "ver_visitantes" {
			return ErrorValidacion(linea)
		} else {
			return linea, Ver_visitantes, nil
		}

	}
}

func ErrorValidacion(linea []string) ([]string, operacion, error) {
	mensaje := []string{"Error en comando", linea[0]}
	fmt.Fprintf(os.Stderr, "%s\n", strings.Join(mensaje, " "))
	return linea, Operacion_invalida, errors.New("error")
}

func ValidarAgregarArchivo(linea []string) (string, error) {
	if !strings.Contains(linea[1], ".") {
		l, _, e := ErrorValidacion(linea)
		return l[0], e
	}
	ruta := strings.Split(linea[1], ".")
	if ruta[1] != "log" {
		l, _, e := ErrorValidacion(linea)
		return l[0], e
	}
	return linea[1], nil
}

func ValidarVerMasVisitados(linea []string) (int, error) {
	n, err := strconv.Atoi(linea[1])
	if err != nil || n < 0 {
		_, _, e := ErrorValidacion(linea)
		return n, e
	}
	return n, nil
}

func ValidarVerVisitantes(linea []string) (Ip, Ip, error) {
	ip_1, error1 := ParsearIP(linea[1])
	ip_2, error2 := ParsearIP(linea[2])
	if error1 != nil || error2 != nil {
		_, _, e := ErrorValidacion(linea)
		return ip_1, ip_2, e
	}
	return ip_1, ip_2, nil
}
