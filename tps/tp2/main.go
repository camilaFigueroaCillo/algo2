package main

import (
	"bufio"
	"os"
	diccionario "tdas/diccionario"
	lectura "tp2/lectura"
	validacion "tp2/validacion"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	ips := diccionario.CrearABB[validacion.Ip, bool](validacion.CmpIp)
	frecuenciasSitios := diccionario.CrearHash[string, int]()
	for s.Scan() {
		lectura.LeerInput(s.Text(), ips, frecuenciasSitios)
	}
}
