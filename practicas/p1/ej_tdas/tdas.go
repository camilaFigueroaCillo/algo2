package main

import (
	"fmt"
	TDACola "tdas/cola"
	TDAPila "tdas/pila"
	//l "guias/p1/ej_tdas/lista"

)

/*Implementar una función que dada una pila, determine si la misma se encuentra ordenada
(es decir, se ingresaron los elementos de menor a mayor). La pila debe quedar en el mismo
estado al original al terminar la ejecución de la función. Indicar y justificar la complejidad
de la función.*/

func EstaOrdenada(p TDAPila.Pila[int]) bool {
	if p.EstaVacia() {
		return true
	}
	a := p.Desapilar()
	if !p.EstaVacia() && a <= p.VerTope() {
		p.Apilar(a)
		return false
	}
	r := EstaOrdenada(p)
	p.Apilar(a)
	return r
}

//Este algoritmo es O(n) porque va desapilando elemento por elemento y comparandolo con el tope de la pila 
//Una vez que termina de analizar si la pila esta ordenada, vuelve a apilar los elementos en la pila original
//Esto seria O(2n), pero como se desprecian las constantes, este algoritmo tiene una complejidad de O(n)

/*Implementar una primitiva para una Cola implementada como una estructura en arreglo (como la vista en clase), 
Filtrar[T](func condicion(T) bool) Cola[T] que devuelva una nueva cola para la cual los elementos de la cola 
original dan true en la función condicion pasada por parámetro. La cola original debe quedar intacta, y los 
elementos de la final deben tener el orden relativo que tenían en la original. Indicar y justificar la 
complejidad del algoritmo implementado.*/

/*Este algoritmo es lineal debido a que primero se desencolan todos los elementos y se encolan a una auxiliar O(2n) 
luego, se van desencolando los elementos de la auxiliar, se evalua con la funcion y si da true se encolan ambas colas y si da false
solo en la cola original
por ultimo se retorna la nueva estructura, en total O(3n) + O(1) , como se desprecian las constantes == O(N) */

func Filter[T any](f func(elem T) bool, c TDACola.Cola[T]) TDACola.Cola[T] {
	cola := TDACola.CrearColaEnlazada[T]()
	aux := TDACola.CrearColaEnlazada[T]()
	
	for !c.EstaVacia() {
		a := c.Desencolar()
		aux.Encolar(a)
	}
	//c esta vacia y tengo aux llena
	//devuelvo los elementos a c original y filtro
	
	for !aux.EstaVacia() {
		a := aux.Desencolar()
		if f(a) {
			cola.Encolar(a)
		}
		c.Encolar(a)
	}

	return cola

}

/*Dada una implementación de Pila, implementar una función func AgregarAlFondo[T any](p Pila[T], elem T) 
recursiva que agregue el elemento al fondo de la pila sin usar TDAs o estructuras auxiliares. Justificar el 
orden de complejidad de la función */

func AgregarAlFondo[T any](p TDAPila.Pila[T], elem T) {
	if p.EstaVacia() {
		p.Apilar(elem)
		return
	}
	a := p.Desapilar()
	AgregarAlFondo(p, elem)
	p.Apilar(a)
}


/*Implementar una función que reciba una Pila genérica y modifique su contenido tal que los elementos sean intercambiados
de a pares. No se pueden usar estructuras auxiliares. Indicar y justificar el orden del algoritmo.
Ejemplo:
Pila original
Tope <- A,B1,B2,C1,C2,D1,D2
Pila tras la función
Tope <- A,B2,B1,C2,C1,D2,D1*/

func alternarPila[T any](pila TDAPila.Pila[T]) {
	c := 0
	pilaInversa := TDAPila.CrearPilaDinamica[T]()
	for !pila.EstaVacia() {
		elem := pila.Desapilar()
		pilaInversa.Apilar(elem)
		c++
	}
	for !pilaInversa.EstaVacia() {
		elem  := pilaInversa.Desapilar()
		pila.Apilar(elem)
	}
	if c % 2 != 0 {
		tope := pila.Desapilar()
		invertirPila(pila)
		pila.Apilar(tope)
		return
	}
	invertirPila(pila)
}

func invertirPila[T any](pila TDAPila.Pila[T]) {
	if pila.EstaVacia() {
		return
	}
	a := pila.Desapilar()
	b := pila.Desapilar()
	invertirPila(pila)
	pila.Apilar(a)
	pila.Apilar(b)
}

/*La complejidad de este algoritmo es O(n) porque iteramos la pila 2 veces para saber su cantidad y luego iteramos la pila recursivamente,
pasando por todos sus elementos 3 veces lo que nos daria una compleidad O(3N) Pero como las constantes se desprecian 
la complejidad termina siendo O(n)
*/

/*Implementar la primitiva de listaEnlazada Extend[T any](otra *listaEnlazada[T]) que extien-
da la lista con todos los elementos de la otra lista que se pasa por parámetro. Indicar y justificar el
orden del algoritmo.*/

/*El Supermercado Noche& tiene la mala costumbre de cerrar cajas de atención cuando todavía tienen clientes 
haciendo fila. No piensan cambiar esta costumbre, pero quieren implementar la redistribución de clientes de la 
forma más prolija posible. 
Dado un slice de colas representando las cajas que continúan en servicio, y un segundo slice de colas
representando las cajas que cierran, implementar una función que redistribuya a todas las personas de las cajas 
que cierran en las diferentes cajas que siguen en servicio, ubicando en cada caja la cantidad de personas lo más
parecida posible. 
Para apaciguar quejas desmedidas, hay que evitar que ningún cliente, experimente cualquiera de las siguientes
situaciones: 
Llamaremos Juan a un cliente en particular, pero esto debe cumplirse para cualquier cliente.
• Juan no debe tener delante suyo alguien que estaba detrás de Juan originalmente.
• En la nueva caja, Juan no debe tener delante suyo a alguien que tenía más personas por delante que él. 

Es decir, si Juan tenía por delante K personas y Lara (también a moverse) tenía L siendo L > K, 
Lara no puede estar delante de Juan en la misma nueva caja. 
Este criterio no aplica para comparar con quienes que ya estaban en la caja.
Indicar y justificar el orden de complejidad de la función.*/


type Caja struct {
	pos int
	encolados int
}

func superNoche(cajasAbiertas, cajasCerradas []TDACola.Cola[string]) {
	
	abiertas := make([]Caja, len(cajasAbiertas))

	for i := range cajasAbiertas {
		c := Caja{i, 0}
		abiertas[i] = c
	} //creo un struct de cajas con la cdad de encolados

	for {
		personas := []string{}

		for _, fila := range cajasCerradas {
			if !fila.EstaVacia() {
				personas = append(personas, fila.Desencolar())
			}
		} //suponiendo que cajas cerradas son 3, entonces en personas tendriamos los 3 primeros
		
		if len(personas) == 0 { //si no apendee nada a personas significa que todas las filas estan vacias
			break              //entonces termino
		}

		for _, p := range personas {
			for i := range abiertas {
				if (i+1) < len(abiertas) && abiertas[i].encolados > abiertas[i+1].encolados {
					continue
				}
				cajasAbiertas[abiertas[i].pos].Encolar(p)
				abiertas[i].encolados++
				break
			}
		}
	}
}


func main() {
	arr := []string{"D2", "D1", "C2", "C1", "B2", "B1", "A"}

	caja_ab1 := TDACola.CrearColaEnlazada[string]()
	caja_ab3 := TDACola.CrearColaEnlazada[string]()
	caja_ab2 := TDACola.CrearColaEnlazada[string]()
	caja_cerrada := TDACola.CrearColaEnlazada[string]()
	for _, e := range arr {
		caja_cerrada.Encolar(e)
	}
	c := []TDACola.Cola[string]{caja_cerrada}
	a := []TDACola.Cola[string]{caja_ab1, caja_ab2, caja_ab3}
	superNoche(a, c)
	
	for i, caja := range a {
		fmt.Printf("caja abierta %d\n", i )
		for !caja.EstaVacia() {
			fmt.Println(caja.Desencolar())
		}
	}

	
	
}




