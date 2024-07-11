from collections import deque
import random

class Grafo:
    def __init__(self, es_dirigido= False, vertices_init=[]):
        self.grafo = {}
        self.es_dirigido = es_dirigido
        for v in vertices_init:
            self.grafo[v] = {}

    def Iterador(self, v):
        return iteradorGrafo(self.grafo, v)
    
    def agregar_vertice(self, v):
        if v in self.grafo:
            raise ValueError(f"Ya hay un vertice {v} en el grafo")
        self.grafo[v] = {}

    def borrar_vertice(self, v):
        if not v in self.grafo:
            raise ValueError("Vertice no pertenece al grafo")
        vertices = self.grafo.keys()
        for w in vertices:
            if w != v:
                self.grafo[w].pop(v, -1)
        self.grafo.pop(v)

    def agregar_arista(self, v, w, peso=1):
        if not v in self.grafo or not w in self.grafo:
            raise ValueError("Vertice no pertenece al grafo")
        vertices = (v, w)
        for i in range(2):
            self.grafo[vertices[i]][vertices[1-i]] = peso
            if self.es_dirigido:
                break
        
    def borrar_arista(self, v, w):
        if not v in self.grafo or not w in self.grafo:
            raise ValueError("Vertice no pertenece al grafo")
        if not (self.estan_unidos(v, w) or self.estan_unidos(w, v)):
            raise ValueError("Vertices no unidos")
        vertices = (v, w)
        for i in range(2):
            self.grafo[vertices[i]].pop(vertices[1-i])
            if self.es_dirigido:
                break
    
    def estan_unidos(self, v, w):
        if not v in self.grafo or not w in self.grafo:
            raise ValueError("Vertice no pertenece al grafo")
        return w in self.grafo[v]


    def peso_arista(self, v, w):
        if not self.estan_unidos(v, w):
            raise ValueError("Vertices no unidos")
        return self.grafo[v][w]

    def obtener_vertices(self):
        return list(self.grafo.keys())
    
    def vertice_aleatorio(self):
        return random.choice(list(self.obtener_vertices()))

    def adyacentes(self, v):
        if v not in self.grafo:
            raise ValueError("Vertice no pertenece al grafo")
        return list(self.grafo[v].keys())
    
    def __str__(self):
        string = ""
        for v, d in self.grafo.items():
            vert = f"{v}: \n"
            for k, p in d.items():
                vert += f"\t {k}: {str(p)} \n"
            string += vert
        return string


class iteradorGrafoBFS:
    def __init__(self, grafo, vertice_inicial):
        self.grafo = grafo
        self.visitados = {vertice_inicial}
        self.cola = deque()
        self.cola.append(vertice_inicial)
        self.actual = None

    def VerActual(self):
        if not self.HaySiguiente():
            raise Exception("No hay mas vertices para ver")
        self.actual = self.cola.popleft()
        self.visitados.add(self.actual)
        return self.actual

    def Siguiente(self):
        if not self.HaySiguiente():
            raise Exception("No hay mas vertices para ver")
        for w in self.grafo[self.actual].keys():
            if w not in self.visitados:
                self.cola.append(w)
        

    def HaySiguiente(self):
        return len(self.visitados) != len(self.grafo.keys())