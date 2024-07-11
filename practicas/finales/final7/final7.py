from collections import deque
from grafo import Grafo

"""Tenemos un plan de estudios que nos indica las correlatividades de las materias que debemos realizar. Suponer que no
hay electivas, ni correlativas por cantidad de créditos. Tenemos un alumno, al que llamaremos agus9900, que quiere
recibirse lo antes posible (es decir, en la mínima cantidad de cuatrimestres). Modelar este problema con grafos, e
implementar una función que reciba dicho grafo y devuelva una lista de listas, donde en la lista i diga qué materias hay
que cursar en el i-ésimo cuatrimestre, de tal manera de tomar la menor cantidad de cuatrimestres en recibirse. Por
supuesto, siempre debe suceder que para toda materia de la lista i, todas sus correlativas deben haberse cursado en
cuatrimestres anteriores. Pueden asumir que agus9900 es tan genio que aprobó todas las cursadas y todos los finales
(en el mismo cuatrimestre de haberlas cursado). Indicar y justificar la complejidad del algoritmo implementado en
función de la cantidad de materias del plan de estudios, y la cantidad de correlatividades.
Tarea para el hogar: para aprender más del lore de FIUBA, pueden leer sobre el legendario"""

#Para resolver este problema, modelo un grafo dirigido cuyos vertices representen las materias y las agregar_aristas representan
#Las correlatividades
#Entonces para Analisis Algebra Proba, existe agregar_arista desde analisis a proba, desde algebra a proba pero no desde analisis a algebra ni viceversa
#Luego realizo un orden topológico para obtener el orden de los que puedo ir cursando

def grados_entrada(grafo):
    g = {v: 0 for v in grafo}
    for v in grafo:
        for w in grafo.adyacentes(v):
            g[w] += 1
    return g

def orden_topologico(grafo):
    q = deque()
    g_ent = grados_entrada(grafo)
    orden = {}
    res = {}
    for v in g_ent:
        if g_ent[v] == 0:
            q.append(v)
            orden[v] = 0

    while q:
        v = q.popleft()
        res[orden[v]] = res.get(orden[v], [])
        res[orden[v]].append(v)
        for w in grafo.adyacentes(v):
            g_ent[w] -= 1
            orden[w] = orden[v]+1
            if g_ent[w] == 0:
                q.append(w)

    resultado = []
    for k, lista in res.items():
        resultado.append(lista)
    
    return resultado 

def _ej_topologico():

    MATERIAS = ["Física I", "Física II", "Física III", "Algoritmos y Programación I", "Algoritmos y Programación II", "Algoritmos y Programación III", "Análisis Matemático II", 'Álgebra II', "Análisis Matemático III", "Probabilidad y Estadística", "Matemática Discreta", "Teoría de Algoritmos I", "Teoría de Algoritmos II", "Química", "Laboratorio", "Estructura del Computador", "Análisis Numérico I", "Organización de Computadoras", "Taller de Programación I", "Organización de Datos", "Taller de Programación II", "Estructura de las Organizaciones", "Modelos y Optimización I", "Sistemas Operativos", "Análisis de la Información", "Técnicas de Diseño", "Base de Datos", "Introducción a los Sistemas Distribuidos"]
    g = Grafo(True, MATERIAS)
    g.agregar_arista("Física I", "Física II")
    g.agregar_arista("Análisis Matemático II", "Física II")
    g.agregar_arista("Algoritmos y Programación I", "Algoritmos y Programación II")
    g.agregar_arista("Algoritmos y Programación II", "Algoritmos y Programación III")
    g.agregar_arista("Algoritmos y Programación II", "Teoría de Algoritmos I")
    g.agregar_arista("Teoría de Algoritmos I", "Teoría de Algoritmos II")
    g.agregar_arista("Matemática Discreta", "Teoría de Algoritmos I")
    g.agregar_arista("Álgebra II", "Física III")
    g.agregar_arista("Física II", "Física III")
    g.agregar_arista("Química", "Física III")
    g.agregar_arista("Física II", "Laboratorio")
    g.agregar_arista("Física II", "Estructura del Computador")
    g.agregar_arista("Algoritmos y Programación II", "Estructura del Computador")
    g.agregar_arista("Álgebra II", "Estructura del Computador")
    g.agregar_arista("Algoritmos y Programación II", "Análisis Numérico I")
    g.agregar_arista("Álgebra II", "Análisis Numérico I")
    g.agregar_arista("Análisis Matemático II", "Análisis Numérico I")
    g.agregar_arista("Álgebra II", "Probabilidad y Estadística")
    g.agregar_arista("Análisis Matemático II", "Probabilidad y Estadística")
    g.agregar_arista("Álgebra II", "Análisis Matemático III")
    g.agregar_arista("Análisis Matemático II", "Análisis Matemático III")
    g.agregar_arista("Estructura del Computador", "Organización de Computadoras")
    g.agregar_arista("Estructura del Computador", "Organización de Datos")
    g.agregar_arista("Algoritmos y Programación II", "Organización de Datos")
    g.agregar_arista("Laboratorio", "Organización de Computadoras")
    g.agregar_arista("Estructura del Computador", "Taller de Programación I")
    g.agregar_arista("Análisis Numérico I", "Taller de Programación I")
    g.agregar_arista("Algoritmos y Programación II", "Taller de Programación I")
    g.agregar_arista("Organización de Datos", "Estructura de las Organizaciones")
    g.agregar_arista("Análisis Matemático III", "Modelos y Optimización I")
    g.agregar_arista("Física II", "Modelos y Optimización I")
    g.agregar_arista("Química", "Modelos y Optimización I")
    g.agregar_arista("Taller de Programación I", "Modelos y Optimización I")
    g.agregar_arista("Organización de Datos", "Sistemas Operativos")
    g.agregar_arista("Taller de Programación I", "Análisis de la Información")
    g.agregar_arista("Algoritmos y Programación III", "Análisis de la Información")
    g.agregar_arista("Análisis de la Información", "Técnicas de Diseño")
    g.agregar_arista("Sistemas Operativos", "Técnicas de Diseño")
    g.agregar_arista("Organización de Datos", "Base de Datos")
    g.agregar_arista("Análisis de la Información", "Base de Datos")
    g.agregar_arista("Organización de Computadoras", "Introducción a los Sistemas Distribuidos")
    g.agregar_arista("Física III", "Introducción a los Sistemas Distribuidos")
    g.agregar_arista("Sistemas Operativos", "Introducción a los Sistemas Distribuidos")
    g.agregar_arista("Taller de Programación I", "Taller de Programación II")
    g.agregar_arista("Modelos y Optimización I", "Taller de Programación II")
    g.agregar_arista("Algoritmos y Programación III", "Taller de Programación II")
    return g