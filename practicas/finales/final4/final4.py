from guias.finales.final7.grafo import Grafo
from unionFind import UnionFind
from collections import deque

"""Implementar el algoritmo de Kruskal para obtener el árbol de tendido mínimo de un grafo. 
Indicar y justificar la complejidad del algoritmo. 
Hacer el seguimiento del algoritmo en el grafo del dorso."""

def kruskal(grafo):
    arbol = Grafo(False, grafo.obtener_vertices())
    aristas = sorted(obtenerAristas(grafo))
    conjuntos = UnionFind(grafo.obtener_vertices())
    for v, w, p in aristas:
        if conjuntos.find(v) == conjuntos.find(w):
            continue
        arbol.agregar_arista(v, w, p)
        conjuntos.union(v, w)
    return arbol


def obtenerAristas(grafo):
    aristas = set()
    for v in grafo:
        for w in grafo.adyacentes(v):
            if (v, w, grafo.peso_arista()) not in aristas:
                aristas.add((v, w, grafo.peso_arista(v, w)))
    return list(aristas)


#La complejidad del algoritmo es O(ElogV) porque primero creamos el nuevo grafo que nos cuesta O(V+E), luego 
#Ordenamos las aristas, esto es O(ElogV), luego vemos todas las aristas y por cada arista realizamos 2 finds, 
#y si lo requerimos una union y agregamos un arista al arbol, agregar una arista es constante mientras que la complejidad
#de unionfind es O(inv(ackermann)) que como la funcion de ackermann es una funcion que crece muy rápido
#la inversa crece muy lento, casi que se considera constante, por lo que los unico que cuesta en este algoritmo
#es ordenar las aristas por su peso.


"""Implementar una función que reciba un grafo y un vértice v, y nos devuelva una lista con todos los vértices 
accessibles desde v (no los adyacentes, sino a los que se puede llegar por algún camino). 
Indicar y justificar la complejidad."""

def vertices_accesibles(grafo, origen):
    return bfs(grafo, origen)

def bfs(grafo, origen):
    res = []
    visitados = {origen}
    q = deque()
    q.append(origen)
    while q:
        v = q.popleft()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                res.append(w)
                visitados.add(w)
                q.append(w)
    return res

#La complejidad del algoritmo es O(V+E) porque por cada vertice veo sus adyacentes y paso por cada vertice a lo sumo 1 vez


