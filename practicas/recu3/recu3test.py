from recu3 import Grafo

g = Grafo(False, [1,2,3,5,4,8,6,7])
g.agregar_arista(1,2)
g.agregar_arista(2,3)
g.agregar_arista(3,4)
g.agregar_arista(4,5)
g.agregar_arista(5,6)
g.agregar_arista(6,7)
g.agregar_arista(7,8)

iterador = g.Iterador(2)
print(iterador.VerActual())
iterador.Siguiente()
print(iterador.VerActual())
iterador.Siguiente()
print(iterador.VerActual())
iterador.Siguiente()
print(iterador.VerActual())
iterador.Siguiente()
print(iterador.VerActual())
iterador.Siguiente()
print(iterador.VerActual())
iterador.Siguiente()
print(iterador.VerActual())
iterador.Siguiente()
print(iterador.VerActual())
print(iterador.HaySiguiente())



