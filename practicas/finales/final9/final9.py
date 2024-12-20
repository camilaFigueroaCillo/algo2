"""
El 10/10 un nuevo algoritmo de ordenamiento ha sido inventado: el MessiSort. Así como no podemos entender 
cómo hace Messi para jugar como lo hace, vamos a asumir que no podemos entender cómo hace este algoritmo 
para ordenar.
El creador del algoritmo (que nada tiene que ver con el astro argentino), declara que el mismo ordena en tiempo
mejor a O(n log n). 
¿Tenés algo para decir sobre su afirmación? Si esto no fuera cierto, ¿podría utilizarse como algoritmo
auxiliar de RadixSort?
"""

#Para que ordene mejor que O(nlogn) entonces el algoritmo deberia ser no comparativo, porque de otra manera, esta afirmación sería falsa,
#puesto que no hay manera de ordenar comparativamente en tiempo mejor a O(nlogn).
#Se podría utilizar como algoritmo auxiliar de RadixSort si el agoritmo es estable, si no lo es, entonces no porque Radix, no ordena si 
# no se usa un algoritmo auxiliar estable.


