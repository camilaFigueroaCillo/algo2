from guias.finales.final7.grafo import Grafo
from collections import deque

"""Implementar una función que reciba un grafo pesado, en el que se asegura que todos los pesos son 1 o 2, 
y un vértice v, y devuelva los caminos mínimos desde v hacia los demás vértices en dicho grafo, 
en tiempo O(V + E).  Justificar la complejidad de la función implementada. 
Dado que dicho algoritmo es lineal, ¿por qué no lo aplicamos para el caso general, en vez de utilizar otros 
algoritmos (por ejemplo, Dijkstra)?"""

#No se aplica en el caso general porque no se tiene informacion sobre los pesos de las aristas.
#En este caso se sabia que los pesos eran 1 y 2, pero en el supuesto caso de que el peso sea un flotante, no se podria representar
#O si el peso de una arista fuese 10000 estaria consumiendo una cantidad de espacio muy grande y no conviene

def caminos_minimos(grafo, v):
    nuevo_grafo = Grafo(False, grafo.obtener_vertices())
    for v in grafo:
        for w in grafo.adyacentes(v):
            if grafo.peso_arista(v, w) == 1:
                nuevo_grafo.agregar_arista(v, w)
            else:
                v_aux = f"aux{v}{w}"
                nuevo_grafo.agregar_vertice(v_aux)
                nuevo_grafo.agregar_arista(v, v_aux)
                nuevo_grafo.agregar_arista(v_aux, w)
    
    padres = bfs(nuevo_grafo, v)
    vertices = set(grafo.obtener_vertices())
    return filtrar_padres(padres, vertices)

def bfs(grafo, v):
    cola = deque()
    visitados = {v}
    padres = {v: None}
    cola.append(v)
    while cola:
        ver = cola.popleft()
        for w in grafo.adyacentes(ver):
            if w not in visitados:
                visitados.add(w)
                padres[w] = v
                cola.append(w)
    return padres

def filtrar_padres(padres, vertices):
    nuevos_padres = {}
    for v in padres:
        if not padres[v]:
            nuevos_padres[v] = None
        elif padres[v] in vertices:
            nuevos_padres[v] = padres[v]
        else:
            nuevos_padres[v] = padres[padres[v]]
    return nuevos_padres

"""Cuando programamos un módulo en Go, tenemos un archivo go.mod que nos indica las dependencias del proyecto.
Asímismo, esas dependencias tienen sus propios go.mod que nos indican sus propias dependencias. Para compilar
nuestro proyecto, Go debe traer (i.e. descargar) y compilar todas nuestras dependencias, así como las dependencias
transitivas. 
Si el módulo A depende del B, es necesario sí o sí compilar primero el módulo B antes que el A.
Explicar detalladamente cómo modelarías este problema con grafos, y cómo obtendrías un orden correcto para 
compilar el proyecto entero (de forma correcta). 
Indicar la complejidad de lo definido, en función de las variables del problema."""

#Este problema se puede modelar con un grafo dirigido de manera que cada vértice sea cada modulo y que las aristas 
# indiquen las dependencias, en este caso: "Si modulo A depende del B" entonces existiria un vertice B, un vertice A
#y una arista que vaya de A hasta B. Luego, para obtener el orden correcto de compilacion, se utiliza un orden topológico
#el cual daria en tiempo lineal el orden de compilación de los archivos.
#Entonces sea M a la cdad de modulos, y D la cdad de dependencias, estariamos diciendo que la complejidad
#del algoritmo es O(M+D).

def g_ent(grafo):
    grados = {v: 0 for v in grafo}
    for v in grafo:
        for w in grafo.adyacentes(v):
            grados[w] += 1
    return grados 

def orden_topologico(grafo):
    g = g_ent(grafo)
    q = deque()
    res = []
    for v in grafo:
        if g[v] == 0:
            q.append(v)
    while q:
        v = q.popleft()
        res.append(v)
        for w in grafo.adyacentes(v):
            g[w] -= 1
            if g[w] == 0:
                q.append(w)
    return res






        




    
