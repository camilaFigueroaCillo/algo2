from collections import deque
from guias.finales.final7.grafo import Grafo


""" 
Se quiere implementar una búsqueda similar a la Búsqueda Binaria: la Búsqueda Ternaria. 
Esta, en vez de partir el arreglo en 2, parte el arreglo en 3 tercios. 
Verifica si el elemento buscado está en la posición del primer tercio, así como
también en la posición del segundo tercio (en vez de únicamente en la mitad, como hace Búsqueda binaria). 
Cuando no se trata de ninguno de estos, llama recursivamente para el segmento que corresponda 
(el primer tercio, segundo o tercero, según cómo sea el elemento buscado respecto al elemento del primer tercio 
y del segundo).

a. Determinar y justificar el orden de la Búsqueda Ternaria.

b. Si en vez de dividir en 3 partes, ahora decidiéramos dividir en n partes (siendo n el tamaño de arreglo), ¿cuál sería
la complejidad del algoritmo? ¿A qué algoritmo se asemeja dicha implementación?

c. Dado los resultados anteriores, ¿tiene sentido implementar la búsqueda K-aria, para k > 2? Justificar.
*/

// a. Como la busqueda ternaria es un algoritmo recursivo con ecuacion de recurrencia de la forma:
// T(n) = 1*(n/3)+O(1)
//Por el teorema maestro podemos saber el orden de la funcion:
// A: 1
// B: 3
//C: 0
//Entonces como log3(1) == 0 -> O(n°log3(n)) = O(n°log(n)) == O(log(n))

//b. Si dividiesemos el arreglo en n, supongamos tenemos [1,2,3,4] ahora tendriamos [1], [2], [3], [4], luego nos
//quedaria verificar que el elemento que estamos buscando este en cada subarreglo, como hacemos en la busqueda 
//binaria cuando verificamos el medio , esta operatoria es lo mismo que si buscasemos
//de manera lineal en todo el arreglo sin haberlo dividido antes, por lo tanto, la complejidad de la busqueda N-aria
//terminaria siendo O(n)

//c. Tendria sentido, si y solo si k es despreciable respecto a n, porque a medida que k se acerca a n 
//tendriamos que verificar en cada subarreglo de unos pocos elementos, todas estas verificaciones nos llevaria a
//terminar viendo aproximadamente casi todos los elementos, por lo que sería una operatoria lineal.

/*Implementar un algoritmo que reciba un Grafo con características de árbol (no árbol binario, sino referido a árbol
de teoría de grafos) y devuelva una lista con los puntos de articulación de dicho árbol. 
Indicar y justificar la complejidad del algoritmo implementado. 
Importante: aprovechar las características del grafo que se recibe para que la solución sea lo más simple posible.*/"""

def grado_salida(grafo):
    return {v: len(grafo.adyacentes(v)) for v in grafo}

def encontrarPA(grafo):
    res = []
    salida = grado_salida(grafo)
    for v in grafo:
        if salida[v] > 1:
            res.append(v)
    return res

#La complejidad de este algoritmo es O(V) porque veo todos los vertices del grafo y realizo operaciones constantes.

            
"""Carlos es nuevo en la empresa en la que trabajan Alan y Bárbara. 
Alan va a ser el mentor de Carlos, quien debe implementar un nuevo TDA Gatito. 
Alan, revisando el trabajo que hizo Carlos, nota que este agregó una primitiva Redimensionar, pública en la 
interfaz Gatito, para que la use Bárbara. Alan lo increpa a Carlos, preguntando para qué es dicha primitiva, 
y este le contesta “Tal como dice la documentación, es para que Bárbara me diga cómo redimensionar el arreglo 
de pelos que tiene el gatito”. Alan, que conoce bien el temperamento de Bárbara, decide evitar que echen a
Carlos en su segunda semana de trabajo. En este ejercicio, te toca hacer de Alan.
Escribir una explicación de por qué esto que está haciendo Carlos está mal. Considerá que Carlos es muy testarudo
(incluso, a pesar de su propio bien), así que tu argumentación deberá ser muy clara y contundente."""

# La primitiva 'Redimensionar' publica está mal debido a que la redimension es algo interno del TDA y el usuario no deberia conocer
# como esta implementado, es el principio de la abstraccion.
# Si barbara quiere un TDA Gatito, no le interesa saber como esta implementado, solo quiere que funcione
# como ella lo pidio y que sea rápido. 
# Si en la implementación Carlos decidió que ese tda sea de manera que se deba redimensionar, entonces sera algo que carlos deba manejar 
# internamente.
#Por último, si le damos la opcion al usuario de estar en contacto con una primitiva privada del tda, podria 
#ingresar algo que pueda romperlo.



