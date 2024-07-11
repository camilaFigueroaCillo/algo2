package main

import (
	"fmt"
)
/*
Sabemos que, en Go, concatenar 2 strings (str1 + str2) es O(n + m) siendo n y m el largo de las cadenas. 
También sabemos que podemos obtener el array de caracteres (runes) que corresponde a un string ([]rune(str)) 
en tiempo lineal, y podemos luego acceder a una posición de un arreglo en tiempo constante. 
Asimismo, podemos obtener nuevamente el string en tiempo lineal (string(arrRunes)).

Queremos implementar un algoritmo que reciba un arreglo de k strings, y devuelva un string con todo junto, 
sin separaciones. 
Alan implementó el algoritmo del dorso, que efectivamente devuelve lo pedido. 
Bárbara lo utilizó, y demoraba demasiado, y le pidió “amablemente” que lo corrija. 
Indicar cuál es el problema del algoritmo, y cómo lo corregirías para que funcione en un tiempo acorde. 
Para tu análisis podés considerar que todas las cadenas son del mismo largo m 
(→ caracteres del arreglo final es n = k · m). 
Aclaración: no usar strings.Join, por supuesto.
func join(cadenas []string) string {
	resultado := ""
	for _, cad := range cadenas { O(k)
	resultado += cad O(m+m)
}
	return resultado 

*/

//En la implementacion de alan podemos ver que por cada cadena, concatena lo anterior con la cadena, a simple vista
//Uno diria que es O(k*(m+n)), por cada cadena realiza una operacion de complejidad O(n+m). 
//Pero, nos estariamos olvidando que la longitud de res, a medida que vamos recorriendo el arreglo de cadenas va aumentando
//POr lo tanto, con k cdad de cadenas y largo de las cadenas = m
//en la primer iteracion tendriamos O(0+m) len(res) = m
//En la 2da iteracion tendriamos O(m+m), ahora len(res) = 2m
//En la 3er iteracion tendriamos O(2m +m), ahora len(res) = 3m 
//y asi hasta k donde tendriamos O((k-1)*m + m) ahora len(res) = k*m

//En total la complejidad seria O(k*(k*m)) = O(k*n)

//Yo lo arreglaria con esta implementacion

func join(cadenas []string) string {
	largo := len(cadenas[0]) //O(1)
	res := make([]rune, len(cadenas)*largo) //O(n)
	var fin int //O(1)
	for i := 0; i < len(cadenas); i++ { //O(k)
		runas := []rune(cadenas[i]) //O(m)
		for i, runa := range runas { //O(m)
			res[i+fin] = runa
		}
		fin = fin + largo //O(1)
	}
	return string(res) //O(n)
}

//Analicemos complejidad: en el for, recorremos cada clave, por cada clave realizamos 2 operaciones que cuestan m
//Osea el largo de la cadena, entonces el for nos estaria costando O(k*m) = O(n), por lo que toda esta funcion 
//Cuesta O(n)
//



/*Existe una estructura llamada dequeue (Double-Ended Queue), que es como una pila y una cola en simultáneo: 
permite insertar al principio y al final, y eliminar tanto al principio como al final. 
Todas esas operaciones, en O(1). ¿Cómo implementarías dicha estructura? 
Definir detalladamente.*/

//Esta estructura la implementaria con una estructura enlazada, donde me voy a guardar las referencias al primero 
//Y al ultimo como en una lista enlazada.
//Entonces ver el primer elemento(como en la cola) es el primer elemento que inserte, y es O(1)
//Ver el tope(pila) es el ultimo elemento que inserte
//Si quiero encolar o apilar, lo que hago es insertar el final y cambiar la referencia al ultimo
//Si quiero desencolar elimino el primero y cambio la referecia
//Si quiero desapilar, elimino el ultimo y cambio la refe
//Y si quisiese saber si esta vacia, pregunto si la cantidad es igual a 0 y ya.(para esto siempre me guardo la refe a cdad y la voy actualizando)
//Entonces mi implemnetacion de DEQUEUE tendria como primitivas:
//VerTope(), VerPrimero(), Encolar(), Desencolar(), Apilar(), Desapilar(), EstaVacia()
//Todas funcionando en O(1)


func main() {
	arr := []string{"hola", "como", "esta"}
	fmt.Println(join(arr))

}