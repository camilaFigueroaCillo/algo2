package diccionario

import (
	"fmt"
)

type estadoClaveDato int

const (
	TAMANIO_INICIAL              = 17
	FNV_OFFSET_BASIS      uint64 = 14695981039346656037
	FNV_PRIME             uint64 = 1099511628211
	REDIMENSION                  = 2
	COMPARACION_EXTENSION        = 0.7
	COMPARACION_REDUCCION        = 0.15
)

const (
	VACIO = estadoClaveDato(iota)
	BORRADO
	OCUPADO
)

type celdaHash[K comparable, V any] struct {
	clave  K
	dato   V
	estado estadoClaveDato
}

type hashCerrado[K comparable, V any] struct {
	tabla    []celdaHash[K, V]
	cantidad int // cantidad de ocupados
	tam      int // cant celdas
	borrados int
}

type iterDiccionario[K comparable, V any] struct {
	hash      *hashCerrado[K, V]
	posActual int
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	tabla := crearTablaHash[K, V](TAMANIO_INICIAL)
	return &hashCerrado[K, V]{tabla, 0, TAMANIO_INICIAL, 0}
}

func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {
	posicion, estado := buscarPosicion(hash.tabla, clave)
	if estado != OCUPADO {
		hash.cantidad++
	}
	hash.tabla[posicion] = celdaHash[K, V]{clave, dato, OCUPADO}
	if hash.factorDeCarga() >= COMPARACION_EXTENSION {
		hash.redimensionar(hash.tam * REDIMENSION)
	}
}

func (hash *hashCerrado[K, V]) Pertenece(clave K) bool {
	_, estado := buscarPosicion(hash.tabla, clave)
	return estado == OCUPADO
}

func (hash *hashCerrado[K, V]) Obtener(clave K) V {
	posicion, estado := buscarPosicion(hash.tabla, clave)
	if estado != OCUPADO {
		panic("La clave no pertenece al diccionario")
	}
	return hash.tabla[posicion].dato
}

func (hash *hashCerrado[K, V]) Borrar(clave K) V {
	posicion, estado := buscarPosicion(hash.tabla, clave)
	if estado != OCUPADO {
		panic("La clave no pertenece al diccionario")
	}
	hash.borrados++
	hash.cantidad--
	dato := hash.tabla[posicion].dato
	hash.tabla[posicion].estado = BORRADO
	if hash.factorDeCarga() <= COMPARACION_REDUCCION && hash.tam >= REDIMENSION*TAMANIO_INICIAL {
		hash.redimensionar(hash.tam / REDIMENSION)
	}
	return dato
}

func (hash *hashCerrado[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash *hashCerrado[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	for _, celda := range hash.tabla {
		if celda.estado != OCUPADO {
			continue
		}
		if !visitar(celda.clave, celda.dato) {
			break
		}
	}
}

func (hash *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	posInicial := buscarPosOcupada[K, V](hash.tabla, -1)
	return &iterDiccionario[K, V]{hash, posInicial}
}

func (iter *iterDiccionario[K, V]) HaySiguiente() bool {
	if iter.posActual > len(iter.hash.tabla) {
		panic("El iterador termino de iterar")
	}
	return iter.posActual != len(iter.hash.tabla)
}

func (iter *iterDiccionario[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	clave := iter.hash.tabla[iter.posActual].clave
	dato := iter.hash.tabla[iter.posActual].dato
	return clave, dato
}

func (iter *iterDiccionario[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iter.posActual = buscarPosOcupada[K, V](iter.hash.tabla, iter.posActual)
}

func buscarPosOcupada[K comparable, V any](tabla []celdaHash[K, V], posicion int) int {
	for j := posicion + 1; j < len(tabla); j++ {
		if tabla[j].estado == OCUPADO {
			return j
		}
	}
	return len(tabla)
}

func (hash *hashCerrado[K, V]) redimensionar(capacidad int) {
	nuevaTabla := crearTablaHash[K, V](capacidad)
	ocupados := 0
	for _, celda := range hash.tabla {
		if celda.estado != OCUPADO {
			continue
		}
		posicion, _ := buscarPosicion(nuevaTabla, celda.clave)
		nuevaTabla[posicion] = crearCeldaHash(celda.clave, celda.dato, OCUPADO)
		ocupados++
	}
	hash.tabla, hash.tam, hash.cantidad, hash.borrados = nuevaTabla, capacidad, ocupados, 0

}

func (hash *hashCerrado[K, V]) factorDeCarga() float64 {
	return (float64(hash.cantidad + hash.borrados)) / float64(hash.tam)
}

func buscarPosicion[K comparable, V any](tabla []celdaHash[K, V], clave K) (int, estadoClaveDato) {
	valorHasheado := funcionHashing(clave, len(tabla))
	i := valorHasheado
	for j := 0; j <= len(tabla); j++ {
		celda := tabla[i]
		if celda.clave == clave && celda.estado == OCUPADO {
			return i, OCUPADO
		} else if celda.clave == clave {
			return i, BORRADO
		} else if celda.estado == VACIO {
			return i, VACIO
		}
		i++
		if i == len(tabla) {
			i = 0
		}
	}
	return 0, VACIO
}

func crearTablaHash[K comparable, V any](largo int) []celdaHash[K, V] {
	tabla := make([]celdaHash[K, V], largo)
	return tabla
}

func crearCeldaHash[K comparable, V any](clave K, dato V, estado estadoClaveDato) celdaHash[K, V] {
	return celdaHash[K, V]{clave, dato, estado}
}

// FNV DECIMAL
// http://www.isthe.com/chongo/tech/comp/fnv/#public_domain
func funcionHashing[K comparable](clave K, largo int) int {
	var hash uint64 = FNV_OFFSET_BASIS
	bytes := []byte(fmt.Sprintf("%v", clave))
	for _, bit := range bytes {
		hash ^= uint64(bit)
		hash *= FNV_PRIME
	}
	if int(hash) < 0 {
		return -int(hash) % largo
	}
	return int(hash) % largo
}


/*Implementar para el Hash Cerrado la primitiva Limpieza(), la cual se encarga de eliminar todos los borrados,
asegurándose de dejar al Hash en un estado correcto (pista: pensar en las búsquedas de los elementos que efectivamente
se encuentran en el hash). Indicar y justificar la complejidad de la primitiva implementada.*/

func (hash *hashCerrado[K, V]) Limpieza() {
	nuevaTabla := make([]celdaHash[K,V], len(hash.tabla))
	for _, celda := range hash.tabla{
		if celda.estado == OCUPADO {
			pos := funcionHashing(celda.clave, len(nuevaTabla))
			if nuevaTabla[pos].estado == OCUPADO {
				for i := pos+1; i < len(nuevaTabla); i++ {
					if nuevaTabla[i].estado == VACIO {
						nuevaTabla[i].clave = celda.clave
						nuevaTabla[i].dato = celda.dato
						break
					}
				}	
			} else {
				nuevaTabla[pos].clave = celda.clave
				nuevaTabla[pos].dato = celda.dato
			}
		}	
	}
	hash.cantidad = hash.cantidad - hash.borrados
	hash.borrados = 0
}