import random

class Grafo:

    def __init__ (self, es_dirigido=False, vertices_init=[]):
        self.es_dirigido = es_dirigido
        self.vertices = {}
        for v in vertices_init:
            self.agregar_vertice(v)

    def agregar_vertice(self, v):
        if v in self.vertices:
            raise ValueError(f"Ya hay un vertice {v} en el grafo")
        self.vertices[v] = {}

    def borrar_vertice(self, v):
        if not v in self.vertices:
            raise ValueError("Vertice no pertenece al grafo")
        vertices = self.vertices.keys()
        for w in vertices:
            if w != v:
                self.vertices[w].pop(v, -1)
        self.vertices.pop(v)

    def agregar_arista(self, v, w, peso=1):
        if not v in self.vertices or not w in self.vertices:
            raise ValueError("Vertice no pertenece al grafo")
        vertices = (v, w)
        for i in range(2):
            self.vertices[vertices[i]][vertices[1-i]] = peso
            if self.es_dirigido:
                break
        
    def borrar_arista(self, v, w):
        if not v in self.vertices or not w in self.vertices:
            raise ValueError("Vertice no pertenece al grafo")
        if not (self.estan_unidos(v, w) or self.estan_unidos(w, v)):
            raise ValueError("Vertices no unidos")
        vertices = (v, w)
        for i in range(2):
            self.vertices[vertices[i]].pop(vertices[1-i])
            if self.es_dirigido:
                break
    
    def estan_unidos(self, v, w):
        if not v in self.vertices or not w in self.vertices:
            raise ValueError("Vertice no pertenece al grafo")
        return w in self.vertices[v]


    def peso_arista(self, v, w):
        if not self.estan_unidos(v, w):
            raise ValueError("Vertices no unidos")
        return self.vertices[v][w]

    def obtener_vertices(self):
        return list(self.vertices.keys())
    
    def vertice_aleatorio(self):
        return random.choice(list(self.obtener_vertices()))

    def adyacentes(self, v):
        if v not in self.vertices:
            raise ValueError("Vertice no pertenece al grafo")
        return list(self.vertices[v].keys())
    
    def __str__(self):
        string = ""
        for v, d in self.vertices.items():
            vert = f"{v}: \n"
            for k, p in d.items():
                vert += f"\t {k}: {str(p)} \n"
            string += vert
        return string
    
    def __iter__(self):
        self.claves = iter(self.vertices)
        return self
    
    def __next__(self):
        return next(self.claves)