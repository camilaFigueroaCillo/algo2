from collections import deque

""" En una facultad contamos únicamente con 2 proyectores. Diferentes docentes de distintos cursos quieren 
utilizarlos (dos cursos no pueden usar el mismo proyector si coindicen en horarios). 
Teniendo la información de los horarios de cada curso, se pide definir si existe una forma de organizar la 
asignación para que todos tengan algún proyector (no importa de momento cuál es la asignación, sólo si existe). 
Modelar este problema utilizando grafos, e implementar un algoritmo que reciba un grafo de las características 
descriptas y resuelva el problema. 
Indicar y justificar la complejidad del algoritmo implementado en función de las variables del problema. """

# Para este problema, se modela un grafo dirigido cuyos vertices representen los horarios de inicio y fin de cada curso, 
# Es decir, cada curso, se representara con 2 vertices en el grafo, un vertice que represente su horario de inicio, y otro
# Vertice que represente su horario de fin. En consiguiente, las aristas unirían estos vertices y tambien existirian
# Aristas que unan el horario de fin de un curso al horario de inicio de otro curso, si el horario de fin es menor o igual al horario de
# inicio del curso 2.
# Luego, para resolver el problema de si existe un orden en el cual todos los cursos puedan usar el proyector, se buscaria que los vertices con grado de entrada 0 no sean mas que 2
# Ya que si tengo mas de dos vertices con grado de entrada 0, significaria que en ese momento hay mas de 2 cursos que necesitan el proyector

def g_e(grafo):
    g = {v:0 for v in grafo}
    for v in grafo:
        for w in grafo.adyacentes(v):
            g[w] += 1
    return g

def existe_manera(grafo):
    
    q = deque()
    encolados = 0
    g = g_e(grafo)
    
    for v in g:
        if g[v] == 0:
            encolados += 1
            q.append(v)
    
    while q:
        if encolados > 2:
            return False
        
        v = q.popleft()
        encolados -= 1
        for w in grafo.adyacentes(v):
            g[w] -= 1
            if g[w] == 0:
                q.append(w)
                encolados += 1
    
    return True
            

    
