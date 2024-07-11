#!/usr/bin/python3
from grafo import Grafo
import sys
import procesar  

def main():
    delincuentes = Grafo(True)
    procesar.cargar_datos(sys.argv[1], delincuentes)
    importantes = []
    comunidades = {}
    componentes = []
    for linea in sys.stdin:
        if not linea.strip(): 
            continue
        procesar.leerEntrada(linea, delincuentes, importantes, comunidades, componentes)
    return

main()
