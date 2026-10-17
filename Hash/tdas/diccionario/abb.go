package diccionario

const (
	_PANIC_DICCIONARIO_VACIO = "El diccionario esta vacio"
	_PANIC_CLAVE_NO_PERTENECE_DICCIONARIO = "La clave no pertenece al diccionario"
)

type diccionarioOrdenado[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cmp      func(K, K) int
	cantidad int
}

type nodoAbb[K comparable, V any] struct {
	clave     K
	dato      V
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	abb := new(diccionarioOrdenado[K, V])
	abb.cmp = funcion_cmp
	return abb
}

func crearNodo[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	nodo := new(nodoAbb[K, V])
	nodo.clave = clave
	nodo.dato = dato
	return nodo
}

func (abb *diccionarioOrdenado[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb *diccionarioOrdenado[K, V]) buscarNodo(raiz **nodoAbb[K, V], clave K) **nodoAbb[K, V] {
	if *raiz == nil {
		return raiz
	}
	comp := abb.cmp(clave, (*raiz).clave)
	if comp == 0 {
		return raiz
	} else if comp < 0 {
		return abb.buscarNodo(&(*raiz).izquierdo, clave)
	} else {
		return abb.buscarNodo(&(*raiz).derecho, clave)
	}
}

func (abb *diccionarioOrdenado[K, V]) Guardar(clave K, dato V) {
	if abb == nil {
		panic(_PANIC_DICCIONARIO_VACIO)
	}
	nodo := abb.buscarNodo(&abb.raiz, clave)
	if *nodo != nil {
		(*nodo).dato = dato
	} else {
		nuevoNodo := crearNodo(clave, dato)
		*nodo = nuevoNodo
		abb.cantidad++
	}
}

func buscarMinimoDerecho[K comparable, V any](nodo **nodoAbb[K, V]) **nodoAbb[K, V] {
	if (*nodo).izquierdo == nil {
		return nodo
	}
	return buscarMinimoDerecho(&(*nodo).izquierdo)
}

func (abb *diccionarioOrdenado[K, V]) Borrar(clave K) V {
	if abb == nil {
		panic(_PANIC_DICCIONARIO_VACIO)
	}
	nodo := abb.buscarNodo(&abb.raiz, clave)
	if *nodo != nil {
		abb.cantidad--
		dato := (*nodo).dato
		if (*nodo).izquierdo == nil && (*nodo).derecho == nil {
			*nodo = nil
			return dato
		} else if (*nodo).izquierdo != nil && (*nodo).derecho == nil {
			*nodo = (*nodo).izquierdo
			return dato
		} else if (*nodo).izquierdo == nil && (*nodo).derecho != nil {
			*nodo = (*nodo).derecho
			return dato
		} 
		menorDerecho := buscarMinimoDerecho(&(*nodo).derecho)
		(*nodo).clave = (*menorDerecho).clave
		(*nodo).dato = (*menorDerecho).dato
		(*menorDerecho) = (*menorDerecho).derecho
		return dato
	}
	panic(_PANIC_CLAVE_NO_PERTENECE_DICCIONARIO)
}
