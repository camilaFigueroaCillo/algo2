package operaciones

import (
	"fmt"
	diccionario "tdas/diccionario"
	validacion "tp2/validacion"
)

func VerVisitantes(ips diccionario.DiccionarioOrdenado[validacion.Ip, bool], desde, hasta *validacion.Ip) {
	fmt.Printf("Visitantes:\n")
	ips.IterarRango(desde, hasta, func(clave validacion.Ip, dato bool) bool {
		k := validacion.ConvertirAString(clave)
		fmt.Printf("\t%s\n", k)
		return true
	})
	fmt.Printf("OK\n")
}
