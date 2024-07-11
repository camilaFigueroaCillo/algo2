import comandos

def cargar_datos(ruta, grafo):
    vertices = set()
    with open(ruta, 'r') as f:
        for mensaje in f:
            mensaje = mensaje.strip().split("\t")
            origen, destino = mensaje
            if origen not in vertices:
                grafo.agregar_vertice(origen)
                vertices.add(origen)
            if destino not in vertices:
                grafo.agregar_vertice(destino)
                vertices.add(destino)
            if not grafo.estan_unidos(origen, destino):
                grafo.agregar_arista(origen, destino)
    return

def leerEntrada(linea, grafo, importantes, comunidades, componentes):
    linea = linea.strip().split()
    return validar_operar(linea, grafo, importantes, comunidades, componentes)

def validar_operar(linea, grafo, importantes, comunidades, componentes): 
    comando = linea[0]
    if comando == "min_seguimientos":
        comandos.min_seguimientos(grafo, linea[1], linea[2])
    elif comando == "mas_imp":
        comandos.mas_importantes(grafo, int(linea[1]), importantes)
    elif comando == "persecucion":
        deli = linea[1].split(",")
        comandos.persecucion(grafo, deli, int(linea[2]), importantes)
    elif comando == "comunidades":
        comandos.comunidades(grafo, int(linea[1]), comunidades)
    elif comando == "divulgar":
        comandos.divulgar(grafo, linea[1], int(linea[2]))
    elif comando == "divulgar_ciclo":
        comandos.divulgar_ciclo(grafo, linea[1])
    elif comando == "cfc":
        comandos.cfc(grafo, componentes)
