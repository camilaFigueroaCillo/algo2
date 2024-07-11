"""Escribir un algoritmo que permita obtener los puntos de articulación de un grafo no dirigido. Indicar y justificar la
complejidad del mismo. Aplicar el algoritmo al grafo del dorso, comenzando desde el vértice G."""

def PA(grafo):
    vis = set()
    res = []
    for v in grafo:
        if v not in vis:
            pa_tarjan(grafo, v, vis, {v: 0}, {v: 0}, {v: None}, res, True)
    return res

def pa_tarjan(grafo, v, vis, orden, mb, padres, res, es_raiz):
    hijos = 0
    orden[v] = mb[v]
    vis.add(v)
    for w in grafo.adyacentes(v):
        if w not in vis:
            padres[w] = v
            hijos += 1
            orden[w] = orden[v]+1
            pa_tarjan(grafo, w, vis, orden, mb, res, False)

            if mb[w] >= orden[v] and not es_raiz:
                res.append(v)
        
            mb[v] = min(mb[v], mb[w])
        elif padre[v] != w:
            mb[v] = min(mb[v], orden[w])
    
    if es_raiz and hijos > 1:
        res.append(v)

#La complejidad del algoritmo es O(V+E) porque por cada vertice veo sus adyacentes una sola vez y realizo operaciones constantes

        
        
