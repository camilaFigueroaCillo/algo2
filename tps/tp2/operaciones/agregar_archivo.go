package operaciones

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	diccionario "tdas/diccionario"
	"time"
	validacion "tp2/validacion"
)

const time_layout = time.RFC3339

type ingreso struct {
	ip     validacion.Ip
	metodo string
	url    string
}

type frecuencia struct {
	tiempo   time.Time
	ingresos int
}

func ProcesarArchivo(ruta string, frecuenciasSitios diccionario.Diccionario[string, int], ips diccionario.DiccionarioOrdenado[validacion.Ip, bool]) {

	archivo, err := os.Open(ruta)
	if err != nil {
		validacion.ErrorValidacion([]string{"agregar_archivo"})
		return
	}

	entradas := diccionario.CrearHash[ingreso, frecuencia]()
	ips_sospechosas := []validacion.Ip{}

	s := bufio.NewScanner(archivo)
	for s.Scan() {
		ips_sospechosas = procesarLinea(s.Text(), entradas, ips_sospechosas, frecuenciasSitios, ips)
	}

	ips_ordenadas := validacion.RadixSortIp(ips_sospechosas)

	for _, ip := range ips_ordenadas {
		fmt.Printf("DoS: %s\n", validacion.ConvertirAString(ip))
	}
	fmt.Printf("OK\n")
}

func procesarLinea(texto string, entradas diccionario.Diccionario[ingreso, frecuencia], ips_sospechosas []validacion.Ip, frecuenciasSitios diccionario.Diccionario[string, int], ips diccionario.DiccionarioOrdenado[validacion.Ip, bool]) []validacion.Ip {
	linea := strings.Split(texto, "\t")
	ip, _ := validacion.ParsearIP(linea[0])
	tiempo, _ := time.Parse(time_layout, linea[1])
	url := strings.Split(linea[3], "/")
	entrada := ingreso{ip, linea[2], url[0]}
	if entradas.Pertenece(entrada) {
		ips_sospechosas = actualizarEntrada(entradas, entrada, tiempo, ips_sospechosas, ip)
	} else {
		entradas.Guardar(entrada, frecuencia{tiempo, 1})
	}
	actualizarFrecuenciasSitios(frecuenciasSitios, linea[3])
	ips.Guardar(ip, true)
	return ips_sospechosas
}

func actualizarEntrada(entradas diccionario.Diccionario[ingreso, frecuencia], entrada ingreso, tiempo time.Time, ips_sospechosas []validacion.Ip, ip validacion.Ip) []validacion.Ip {
	f := entradas.Obtener(entrada)
	if f.ingresos < 5 {
		diferencia_tiempo := tiempo.Sub(f.tiempo)
		if diferencia_tiempo.Seconds() < 2 {
			nueva_frecuencia := frecuencia{f.tiempo, f.ingresos + 1}
			entradas.Guardar(entrada, nueva_frecuencia)
			if f.ingresos+1 == 5 {
				ips_sospechosas = append(ips_sospechosas, ip)
			}
		} else {
			entradas.Guardar(entrada, frecuencia{tiempo, 1})
		}
	}
	return ips_sospechosas
}

func actualizarFrecuenciasSitios(frecuencias diccionario.Diccionario[string, int], url string) {
	if frecuencias.Pertenece(url) {
		f := frecuencias.Obtener(url)
		frecuencias.Guardar(url, f+1)
	} else {
		frecuencias.Guardar(url, 1)
	}
}
