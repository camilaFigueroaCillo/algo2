package main

import ( 
	"fmt"
)

type Personas struct {
	nombre string
	cumple int
}

func buscarCumples(dia int, integrantes []Personas) []Personas {
	
	if len(integrantes) == 0 {
		return []Personas{}
	}
	arr := buscarCumplesRec(dia, integrantes, 0, len(integrantes)-1)
	return arr
}

func buscarCumplesRec(dia int, integrantes []Personas, ini int, fin int) []Personas {
	arr := []Personas{}
	if ini == fin {
		return arr
	}

	medio := (ini+fin)/2
	if integrantes[medio].cumple == dia {
		arr = append(arr, integrantes[medio])
	}
	if integrantes[medio].cumple > dia || integrantes[medio+1].cumple > dia {
		izq := buscarCumplesRec(dia, integrantes, ini, medio)
		arr = append(arr, izq...)
		return arr
	}
	der := buscarCumplesRec(dia, integrantes, medio+1, fin)
	arr = append(arr, der...)
	return arr
}



/*func ordenarPalabras(arr string) []string {
	countingSort(arr, 27, 4)
	countingSort(arr, 27, 3)
	countingSort(arr, 27, 2)
	countingSort(arr, 27, 1)
	countingSort(arr, 27, 0)
	return arr
}


func countingSort(arr string, rango int, criterio int) {
	frecuencias := make([]string, rango)
	sumasAcum := make([]string, rango)
	res := make([]string, len(arr))

	for _,e := range arr {
		valor := int(e[criterio] -0)
		frecuencias[valor]++
	}

	for i:= 1; i < len(frecuencias); i++ {
		sumasAcum[i] = sumasAcum[i-1]+frecuencias[i-1]
	}

	for _,e := range arr {
		valor := int(e[criterio] -0)
		pos := sumasAcum[valor]
		res[pos] = e 
		sumasAcum[valor]++
	}

	copy(arr, res)

}*/

func main(){
	p1 := Personas{"camila", 1}
	p2 := Personas{"facu", 60}
	p3 := Personas{"joa", 47}
	p4 := Personas{"ani",1}
	p5 := Personas{"belu",5}
	arr := []Personas{p1,p4,p5,p3,p2}
	arr2 := buscarCumples(1, arr)
	fmt.Println(arr2)



}