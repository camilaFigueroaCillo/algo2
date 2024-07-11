from guias.finales.final7.grafo import Grafo
from collections import deque

"""¿Qué implica que un algoritmo de ordenamiento sea estable? Explicar detalladamente por qué el algoritmo de
ordenamiento auxiliar de RadixSort debe ser sí o sí estable. Dar algún ejemplo en el que se evidencie esta necesidad."""
#Que un algoritmom de ordenamiento sea estable, significa que mantiene el orden relativo en el que estaban los elementos
#antes de ordenarlos
#POr ejemplo, si tengo un arreglo de cartas desordenado donde tengo primero una carta negra que vale 4 y luego una
#carta blanca que tambien vale cuatro, cuando las ordene, primero va a estar la carta roja que valia 4 y luego va a estar 
#la carta blanca que valia 4.
#RadixSort debe utilizar un algoritmo de ordenamiento estable porque si no lo ultilizaramos, en cada ordenamiento
#por los criterios de menor a mayor importancia, estariamos perdiendo el orden relativo anterior.
#Ej: Si quisiesemos ordenar el arreglo [347, 521, 148] y en lugar de usar counting utilizaramos Quicksort que no es
#estable, en el primer ordenamiento por las cifras de las unidades nos quedaria:
#[321, 347, 348], luego
#[321, 348, 347] y despues 
#[321, 348, 347] y como vemos, no esta ordenado. 

"""En clase se ha demostrado que ningún algoritmo de ordenamiento comparativo puede ser mejor que O(n log n), 
siendo n la cantidad de elementos del arreglo a ordenar. 
Al mismo tiempo, hemos visto algoritmos para obtener un ordenamiento topológico de un conjunto de datos que, 
modelando el problema con grafos (dirigidos), termina resolviéndose en tiempo lineal a la cantidad de vértices 
y aristas del mismo. 
Explicar por qué en problemas de ordenamientos (e.g. arreglo de
números de los que no se tiene información) no modelamos siempre a estos con grafos y luego resolvemos utilizando
alguno de los algoritmos antes mencionados.
Recomendación: pensar bien este ejercicio, hacerse dibujos, etc. . . antes de ponerse a escribir la respuesta.
No responder “no se puede resolver de esta forma”, porque sí se puede. 
Por el contrario, para plantear la respuesta conviene plantearse cómo se haría para resolver un problema de 
ordenamiento usando una resolución de orden topológico. Esta consigna no
va sobre la posibilidad, sino por qué no es la mejor alternativa para hacerlo, a pesar de parecer una idea tentadora."""

#No resolvemos ese problema de esa manera porque no mejoraria la complejidad.
#Por ejemplo, si yo planteo resolver el problema de ordenar el arreglo [14,17, 3.1418, 2.1514, 1.24] con grafos
#y digo bueno cada numero es un vertice y cada arista representa "V1" es menor a "V2", para agregar estas aristas
# a cada vertice voy a tener que
#compararlo con todos los otros vertices y agregar las aristas correspondientes, entonces el proceso de agregar 
# todas las aristas nos esta costando O(V*V)
#Por otro lado, en espacio tambien es peor, el arreglo de enteros (en funcion de las variables del grafo), 
# nos ocupa O(V) mientras que el grafo nos ocupa O(V+E) siendo E la cdad de aristas que haya en el grafo.

#Entonces si sabemos que la cdad de vertices es igual a la cdad de elementos del arreglo, por lo que O(V) == O(N)
#Podemos decir que armar el grafo nos esta costando O(N + N*N) y todavia nos queda realizarle
#Un algoritmo de orden topologico, lo cual es mucho peor a hacer un ordenamiento O(nlogn) y ya.
#Por lo tanto, si modelar este problema con grafos nos empeora la complejidad, para que hacerlo? mejor aplicar un algoritmo
#de ordenamiento y ya.


"""Dado un grafo no dirigido, pesado con todas aristas de pesos diferentes, implementar un algoritmo que reciba 
un grafo y una arista y determine si esa arista pertenece al único árbol de tendido mínimo dentro del grafo, 
o no, en O(V + E). 
Justificar la complejidad del algoritmo implementado.
Ayuda: pensar qué sucede si la arista no es parte del MST, y qué condiciones se deben dar en el grafo respecto 
a los dos extremos de la arista en cuestión."""

def perteneceAristaMST(grafo, arista):
    v, w, peso = arista
    grafo.borrar_arista(v, w)
    suma_pesos, hay_camino = bfs(grafo, v, w)
    return (not hay_camino) or (hay_camino and suma_pesos > peso)

def bfs(grafo,origen, destino):
    visitados = set()
    q = deque()
    q.append(origen)
    visitados.add(origen)
    padres = {origen: None}
    while q:
        v = q.popleft()
        if v == destino:
            return armar_camino(grafo, padres, origen, destino), True
        for w in grafo.adyacentes(v):
            if w not in visitados:
                visitados.add(w)
                padres[w] = v
                q.append(w)
    return None, False

def armar_camino(grafo, padres, origen, destino):
    v = destino
    res = 0
    while v != origen:
        res += grafo.peso_arista(v, padres[v])
        v = padres[v]
    return res


g = Grafo(False, ["A", "B", "C", "D", "F"])
g.agregar_arista("A", "B", 1)
g.agregar_arista("A", "C", 2)
g.agregar_arista("C", "B", 3)
g.agregar_arista("C", "D", 4)
g.agregar_arista("B", "F", 5)
g.agregar_arista("B", "D", 7)

print(perteneceAristaMST(g, ("B", "D", 7)))
print(perteneceAristaMST(g, ("B", "C", 3)))
print(perteneceAristaMST(g, ("A", "B", 1)))




    