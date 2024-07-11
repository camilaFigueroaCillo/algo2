package lectura

import (
	diccionario "tdas/diccionario"
	operar "tp2/operaciones"
	validacion "tp2/validacion"
)

func LeerInput(comando string, ips diccionario.DiccionarioOrdenado[validacion.Ip, bool], frecuenciasSitios diccionario.Diccionario[string, int]) {
	linea, operacion, err := validacion.ValidarEntrada(comando)
	if err != nil {
		return
	}
	if operacion == validacion.Agregar_archivo {
		ruta, err := validacion.ValidarAgregarArchivo(linea)
		if err != nil {
			return
		}
		operar.ProcesarArchivo(ruta, frecuenciasSitios, ips)
	} else if operacion == validacion.Ver_mas_visitados {
		n, err := validacion.ValidarVerMasVisitados(linea)
		if err != nil {
			return
		}
		operar.VerMasVisitados(frecuenciasSitios, n)
	} else {
		ip1, ip2, err := validacion.ValidarVerVisitantes(linea)
		if err != nil {
			return
		}
		operar.VerVisitantes(ips, &ip1, &ip2)
	}
}
