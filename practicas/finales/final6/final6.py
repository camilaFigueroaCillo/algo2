"""Responder (justificando) las siguientes preguntas:
a. ¿Qué complejidad espacial tiene un iterador interno preorder de un árbol binario?
b. ¿Qué complejidad espacial tiene un iterador interno por niveles de un árbol binario?
c. Si dos árboles binarios tienen mismo preorder y postorder entonces, ¿Deben tener la misma estructura interna (es
decir, ser el mismo árbol)?"""

#La complejidad espacial del iterador interno preorder de un arbol binario es logn en el stack de la recursion.
#Porque cuando termino las llamadas del lado izquierdo, voy volviendo en la recursion y a bajo nivel
#Esas direcciones de los nodos que ya visite se van desapilando mientras que se van apilando aquellas que me quedan por ver

#La complejidad espacial del iterador interno por niveles es O(N), porque en el ultimo nivel voy a tener encoladas
#todas las hojas del arbol.

#Falso, yo podria tener un AB en forma de lista A -> B -> C y un AB en zigzag A 
                                                                            #/ 
                                                                            #B
                                                                            #\C
#Que tienen el mismo preorder y postorder y no tienen la misma estructura interna

"""Implementar un algoritmo que reciba un grafo no dirigido y conexo, y determine si el mismo es en realidad 
biconexo.
Un grafo es biconexo si para cada par de vértices existen al menos dos caminos que los unen 
(es decir, los dos forman parte de al menos un ciclo). 
Indicar y justificar la complejidad del algoritmo implementado."""

#Solucion, si existen puntos de articulacion entonces el grafo no es biconexo

def es_biconexo(grafo):
    ptos = dfs_pa(grafo)
    return not ptos

def dfs_pa(grafo):
    vis = set()
    pa = set()
    for v in grafo:
        if v not in vis:
            dfs_tarjan(grafo, v, vis, {v:None}, {v:0}, {}, pa, True)
    return pa

def dfs_tarjan(grafo, v, vis, padres, orden, mb, pa, es_raiz):
    mb[v] = orden[v]
    hijos = 0
    vis.add(v)
    for w in grafo.adyacentes(v):
        if w not in vis:
            orden[w] = orden[v]+1
            hijos += 1
            padres[w] = v
            dfs_tarjan(grafo, w, vis, orden, mb, pa, False)
            
            if mb[w] >= orden[v] and not es_raiz:
                pa.add(v)
            
            mb[v] = min(mb[v], mb[w])

        elif padres[v] != w:
            mb[v] = min(mb[v], orden[w])
        
    if es_raiz and hijos > 1:
        pa.add(v)


        
"""Realizar un seguimiento de aplicar Radix Sort para ordenar el siguiente arreglo de países de forma alfabética. Considerar
que no todos los países tienen el mismo largo pero sí es similar (explicar cómo se resuelve esto):
[Brasil(6), Argentina(9), Venezuela(9), Ecuador(7), Bolivia(7), Surinam(7), Uruguay(7), Colombia(8), 
Guyana(6), Paraguay(8)] """

#Argentina, Venezuela, Brasil, Ecuador, Bolivia, Surinam, Uruguay, Colombia, Guyana, Paraguay
#Colombia, Venezuela, Argentina, Paraguay, Brasil, Ecuador, Bolivia, Surinam, Uruguay, Guyana
#Paraguay, Bolivia, Venezuela, Colombia, Argentina, Surinam, Ecuador, Uruguay, Brasil, Guyana
#Surinam, Uruguay, Guyana, Colombia, Bolivia, Brasil, Ecuador, Argentina, Paraguay, Venezuela
#Ecuador, Paraguay, Brasil, Col mbia, Surinam, Guyana, Argentina, Uruguay, Bolivia, Venezuela
#Ecuador, Paraguay, Gu yana, Argentina, Venezuela, Surinam, Bolivia, Uruguay, Colombia, Brasil
#Brasil, Argentina, Bolivia, Colombia, Venezuela, Paraguay, Surinam, Ecuador, Uruguay, Guyana
#Paraguay, Ecuador, Venezuela, Bolivia, Colombia, Brasil, Argentina, Uruguay, Surinam, Guyana
#Argentina,Bolivia,BRasil, Colombia, Ecuador, Guyana, Paraguay, Surinam, Uruguay, Venezuela

#Yo lo resolveria haciendo un bucket con los paises que tienen la misma cantidad de letras y ordeno por ese 
# digito, lo junto con los del digito anterior a ese y ordeno ese array parcial y asi sucesivamente 
# hasta juntar todo el array.
# 



