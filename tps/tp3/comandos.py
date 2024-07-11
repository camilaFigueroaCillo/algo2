import biblioteca

def min_seguimientos(grafo, origen, destino):
    padres, _, anterior = biblioteca.bfs(grafo, origen, destino)
    if not anterior:
        print("Seguimiento imposible")
    else:
        camino = biblioteca.armar_camino(padres, origen, destino)
        print(camino)

def mas_importantes(grafo, cantidad, importantes):
    if not importantes:
        importantes = biblioteca.page_rank(grafo)
    imp = ""
    for i in range(cantidad-1):
        imp += f"{importantes[i]}, " 
    imp += importantes[cantidad-1]
    print(imp)

def persecucion(grafo, delincuentes, k, importantes):
    if not importantes:
        importantes = biblioteca.page_rank(grafo)
    largo_camino = float("inf")
    perseguido = None
    camino = None
    for d in delincuentes:
        padre, dist, _ = biblioteca.bfs(grafo, d)
        for i in range(k):
            if not importantes[i] in dist:
              continue
            if dist[importantes[i]] < largo_camino:
                perseguido = importantes[i]
                largo_camino = dist[perseguido] 
                camino = (padre, d, perseguido)
    camino_minimo = biblioteca.armar_camino(camino[0], camino[1], camino[2])
    print(camino_minimo)

def comunidades(grafo, n, comunidades):
    if not comunidades:
        comunidades = biblioteca.armar_comunidades(grafo)
    i = 1
    for _, comunidad in comunidades.items():
        if len(comunidad) < n:
            continue
        comunidad = ", ".join(comunidad)
        print(f"Comunidad {i}: " + comunidad)
        i += 1 

def divulgar(grafo, origen, n):
    delincuentes = biblioteca.bfs_rango(grafo, origen, n)
    print(", ".join(delincuentes))

def divulgar_ciclo(grafo, delincuente):
    camino, _, anterior = biblioteca.bfs(grafo, delincuente, delincuente)
    if not anterior:
        print("No se encontro recorrido")
    else:
        camino = biblioteca.armar_camino(camino, delincuente, anterior) +  " -> " + delincuente
        print(camino)

def cfc(grafo, componentes):
    if not componentes:
        componentes = biblioteca.cfcs_grafo(grafo) 
    i = 1
    for comp in componentes:
        cadena = ", ".join(comp)
        print(f"CFC {i}: " + cadena)
        i += 1
