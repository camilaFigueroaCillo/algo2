from collections import deque
from random import shuffle
from pila import Pila

D_RANK = 0.85
ITERACIONES_LABEL = 5
IT_PAGERANK = 10

def bfs(grafo, origen, destino = None):
    padres = {origen: None}
    dist = {origen: 0}
    vis = set()
    q = deque()
    q.append(origen)
    vis.add(origen)
    while q:
        v = q.popleft()
        for w in grafo.adyacentes(v):
            if w not in vis:
                vis.add(w)
                padres[w] = v
                dist[w] = dist[v] + 1
                q.append(w)
            if w == destino:
                return padres, dist, v
    return padres, dist, None

def armar_camino(padres, origen, destino):
    camino = []
    v = destino
    while v != origen:
        if not v:
            break
        camino.append(v)
        camino.append("->")
        v = padres[v]
    camino.append(origen)
    return " ".join(camino[::-1])

def vertices_entrada(grafo):
    vertices = {}
    for v in grafo:
        vertices[v] = []
    for v in grafo:
        for w in grafo.adyacentes(v):
            vertices[w].append(v)
    return vertices

def grado_salida(grafo):
    return {v: len(grafo.adyacentes(v)) for v in grafo}

def page_rank(grafo):
    entrada = vertices_entrada(grafo)
    N = len(grafo.obtener_vertices())
    pr = {v: 1/N for v in grafo}
    salida = grado_salida(grafo)

    for _ in range(IT_PAGERANK):

        new_pr = {}

        for v in grafo:
            new_pr[v] = (1-D_RANK)/N
            for w in entrada[v]: 
                new_pr[v] += D_RANK * pr[w] / salida[w]
           
        pr = new_pr

    return sorted(pr, key=lambda v:pr[v], reverse=True)

def label_propagation(grafo):
    entrada = vertices_entrada(grafo)
    i = 1
    labels = {}
    vertices = list(grafo.obtener_vertices())
    for v in vertices:
        labels[v] = i
        i += 1
    shuffle(vertices)
    for _ in range(ITERACIONES_LABEL):
        new_labels = {}
        for v in vertices:
            new_labels[v] = max_frec(entrada[v], labels)
        labels = new_labels
    return labels

def max_frec(entrada, labels):
    maxima = 0
    frec = 0
    labels_entrada = {}
    for w in entrada:
        labels_entrada[labels[w]] = labels_entrada.get(labels[w], 0) + 1 
        frec = max(frec, labels_entrada[labels[w]])
        if frec == labels_entrada[labels[w]]:
            maxima = labels[w]
    return maxima

def armar_comunidades(grafo):
    labels = label_propagation(grafo)
    comunidades = {}
    for v, label in labels.items():
        if label in comunidades:
            comunidades[label].append(v)
        else:
            comunidades[label] = [v]
    return comunidades

def bfs_rango(grafo, origen, n):
    visitados = set()
    q = deque()
    q.append(origen)
    dist = {origen: 0}
    while q:
        v = q.popleft()
        if dist[v] == n:
            continue
        for w in grafo.adyacentes(v):
            if w == origen:
                continue
            if not w in visitados:
                visitados.add(w)
                dist[w] = dist[v] + 1
                q.append(w)
    return visitados

def dfs_cfc(grafo, v, visitados, orden, mas_bajo, pila, apilados, cfcs, contador_global):
    orden[v] = mas_bajo[v] = contador_global[0]
    contador_global[0] += 1
    visitados.add(v)
    pila.apilar(v)
    apilados.add(v)
    for w in grafo.adyacentes(v):
        if w not in visitados:
            dfs_cfc(grafo, w, visitados, orden, mas_bajo, pila, apilados, cfcs, contador_global)
        if w in apilados:
            mas_bajo[v] = min(mas_bajo[v], mas_bajo[w])

    if orden[v] == mas_bajo[v]:
        nueva_cfc = []
        while True: 
            w = pila.desapilar()
            apilados.remove(w)
            nueva_cfc.append(w)
            if w == v:
                break
        cfcs.append(nueva_cfc)

def cfcs_grafo(grafo):
    resultados = []
    visitados = set()
    contador_global = [0]
    for v in grafo.obtener_vertices():
        if v not in visitados:
            dfs_cfc(grafo, v, visitados, {}, {}, Pila(), set(), resultados, contador_global)
    return resultados