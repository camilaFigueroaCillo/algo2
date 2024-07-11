package validacion

import (
	"errors"
	"strconv"
	"strings"
)

type Ip struct {
	campo1 int
	campo2 int
	campo3 int
	campo4 int
}

func ParsearIP(linea string) (Ip, error) {
	elementos := strings.Split(linea, ".")
	if len(elementos) != 4 {
		return Ip{}, errors.New("error")
	}
	campos := [4]int{}
	for i, elem := range elementos {
		campo, err := strconv.Atoi(elem)
		if err != nil || campo < 0 || campo > 255 {
			return Ip{}, errors.New("error")
		}
		campos[i] = campo
	}
	return Ip{campos[0], campos[1], campos[2], campos[3]}, nil
}

func CmpIp(ip_1 Ip, ip_2 Ip) int {
	if ip_1.campo1 != ip_2.campo1 {
		return ip_1.campo1 - ip_2.campo1
	}
	if ip_1.campo2 != ip_2.campo2 {
		return ip_1.campo2 - ip_2.campo2
	}
	if ip_1.campo3 != ip_2.campo3 {
		return ip_1.campo3 - ip_2.campo3
	}
	if ip_1.campo4 != ip_2.campo4 {
		return ip_1.campo4 - ip_2.campo4
	}
	return 0
}

func RadixSortIp(arr []Ip) []Ip {
	countingSortIp(arr, 256, obtenerCampo, 4)
	countingSortIp(arr, 256, obtenerCampo, 3)
	countingSortIp(arr, 256, obtenerCampo, 2)
	countingSortIp(arr, 256, obtenerCampo, 1)
	return arr
}

func countingSortIp(arr []Ip, rango int, campo func(Ip, int) int, n int) {
	frecuencias := make([]int, rango)
	sumasAcumuladas := make([]int, rango)
	resultado := make([]Ip, len(arr))

	for _, ip := range arr {
		campo := campo(ip, n)
		frecuencias[campo]++
	}

	for i := 1; i < rango; i++ {
		sumasAcumuladas[i] = sumasAcumuladas[i-1] + frecuencias[i-1]
	}

	for _, ip := range arr {
		campo := campo(ip, n)
		posicion := sumasAcumuladas[campo]
		resultado[posicion] = ip
		sumasAcumuladas[campo]++
	}

	copy(arr, resultado)
}

func obtenerCampo(ip Ip, n int) int {
	campos := []int{ip.campo1, ip.campo2, ip.campo3, ip.campo4}
	return campos[n-1]
}

func ConvertirAString(ip Ip) string {
	campo1 := strconv.Itoa(ip.campo1)
	campo2 := strconv.Itoa(ip.campo2)
	campo3 := strconv.Itoa(ip.campo3)
	campo4 := strconv.Itoa(ip.campo4)
	campos := []string{campo1, campo2, campo3, campo4}
	return strings.Join(campos, ".")
}
