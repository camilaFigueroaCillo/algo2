"""
El 10/10 un nuevo algoritmo de ordenamiento ha sido inventado: el MessiSort. Así como no podemos entender 
cómo hace Messi para jugar como lo hace, vamos a asumir que no podemos entender cómo hace este algoritmo 
para ordenar.
El creador del algoritmo (que nada tiene que ver con el astro argentino), declara que el mismo ordena en tiempo
mejor a O(n log n). 
¿Tenés algo para decir sobre su afirmación? Si esto no fuera cierto, ¿podría utilizarse como algoritmo
auxiliar de RadixSort?
"""

#Lo que puedo decir sobre la afirmacion es que para que ordene mejor que O(nlogn) entonces el algoritmo deberia ser 
#no comparativo, porque de otra manera, esta afirmación sería falsa, puesto que no hay manera de ordenar comparativamente
#En tiempo mejor a O(nlogn).
#Sem podría utilizar como algoritmo auxiliar de RadixSort si y solo si el agoritmo es estable, si no lo es, entonces no
#Porque Radix, no ordena si no usamos un algoritmo auxiliar estable.


