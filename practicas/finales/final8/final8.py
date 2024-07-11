"""Implementar un algoritmo que reciba un grafo dirigido, acíclico y pesado, un vértice v y otro w,
 y devuelva la longitud del camino máximo. Indicar y justificar la complejidad del algoritmo implementado."""

def camino_max(grafo, origen, destino):
    h = heapMaximos()
    padres = {origen: None}
    distancias = {}
    for v in grafo:
        distancias[v] = float("-inf")
    distancias[origen] = 0
    h.Encolar(0, origen)
    while h:
        dist, v = h.Desencolar()
        for w in grafo.adyacentes(v):
            if dist + grafo.peso_arista(v, w) > distancias[w]:
                distancias[w] = distancias[v] + grafo.peso_arista(v, w)
                padres[w] = v
                h.Encolar(distancias[w], w)

    return armar_camino(padres, origen, destino, grafo)

def armar_camino(p, o, d, grafo):
    v = d
    s = 0
    while v != o:
        s += grafo.peso_arista(p[v], v)
        v = p[v]
    return s


"""Tenemos un mapa de caminos rurales que conectan diferentes ciudades, donde algunos de estos caminos se 
encuentran bloqueados por alguna razón (un árbol caido, una piedra que cayó desde una montaña, 
los festejos de boquita campeón quemando un local de comida rápida, etc. . . ). 
Cada bloqueo cuesta diferente de remover. Sabemos que algunos de estos bloqueos no nos impiden llegar desde una 
ciudad a otra, pero podría ser el caso que sí. 
Queremos implementar un algoritmo que nos determine qué bloqueos deben ser removidos para que se pueda llegar 
de cualquier ciudad a cualquier otra. 
Para esto, primero modelar el problema con grafos, y luego implementar un algoritmo que reciba dicho grafo y nos
devuelva los bloqueos a eliminar. 
Indicar y justificar la complejidad del algoritmo implementado."""

#Para este problema, modelo un grafo no dirigido cuyos vertices sean los lugares a los que quiero llegar y las aristas representan
#el camino. Si el camino esta bloqueado entonces esa arista sera pesada, si no, será no pesada. Luego para encontrar
#Que bloqueos deberia remover aplico un algoritmo de kruskal para encontrar el arbol de tendido minimo del grafo, esto
#nos dará las aristas que nos cuestan menos y con las cuales podemos conectar todas las ciudades con todas.

def caminos_rurales(grafo):
    mst = kruskal(grafo)
    aristas = obtenerAristas(mst)
    return aristas

def obtenerAristas(grafo):
    aristas = set()
    visitados = set()
    for v in grafo:
        for w in grafo.adyacentes(v):
            if w in visitados:
                continue
            aristas.add(v, w, grafo.peso_arista(v, w))
        visitados.add(v)
    return list(aristas)

def kruskal(grafo):
    aristas = sorted(obtenerAristas(grafo))
    mst = Grafo(False, grafo.obtener_vertices())
    conjuntos = UnionFind(grafo.obtener_vertices())
    for v, w, peso in aristas:
        if conjuntos.find(v) == conjuntos.find(w):
            continue
        mst.agregar_arista(v, w, peso)
        conjuntos.union(v, w)
    return mst

