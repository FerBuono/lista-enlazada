package lista

type nodoLista[T any] struct {
	dato T
	prox *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type iterListaEnlazada[T any] struct {
	lista    *listaEnlazada[T]
	actual   *nodoLista[T]
	anterior *nodoLista[T]
}

func (l *listaEnlazada[T]) EstaVacia() bool {
	return l.largo == 0 && l.primero == nil && l.ultimo == nil
}

func (l *listaEnlazada[T]) InsertarPrimero(dato T) {
	nuevo := l.crearNodo(dato)

	if l.EstaVacia() {
		l.primero = nuevo
		l.ultimo = nuevo
	} else {
		nuevo.prox = l.primero
		if l.primero.prox == nil {
			l.primero = nuevo
			l.ultimo = l.primero.prox
		} else {
			l.primero = nuevo
		}
	}

	l.largo++
}

func (l *listaEnlazada[T]) InsertarUltimo(dato T) {
	nuevo := l.crearNodo(dato)

	if l.EstaVacia() {
		l.primero = nuevo
	} else {
		if l.primero.prox == nil {
			l.primero.prox = nuevo
		} else {
			l.ultimo.prox = nuevo
		}
	}

	l.ultimo = nuevo
	l.largo++
}

func (l *listaEnlazada[T]) BorrarPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}

	primero := l.primero.dato

	if l.primero.prox == nil {
		l.primero = nil
		l.ultimo = nil
	} else {
		l.primero = l.primero.prox
	}

	l.largo--
	return primero
}

func (l *listaEnlazada[T]) VerPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}

	return l.primero.dato
}

func (l *listaEnlazada[T]) VerUltimo() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}

	return l.ultimo.dato
}

func (l *listaEnlazada[T]) Largo() int {
	return l.largo
}

func (l *listaEnlazada[T]) Iterador() IteradorLista[T] {
	iter := new(iterListaEnlazada[T])
	iter.lista = l
	iter.actual = l.primero
	return iter
}

func (i *iterListaEnlazada[T]) VerActual() T {
	return i.actual.dato
}

func (i *iterListaEnlazada[T]) HaySiguiente() bool {
	return i.actual.prox != nil
}

func (i *iterListaEnlazada[T]) Siguiente() T {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	actual := i.actual.dato
	i.anterior = i.actual
	i.actual = i.actual.prox

	return actual
}

func (i *iterListaEnlazada[T]) Insertar(T) {

}

func (i *iterListaEnlazada[T]) Borrar() T {
	return i.actual.dato
}

func (l *listaEnlazada[T]) crearNodo(dato T) *nodoLista[T] {
	nuevo := new(nodoLista[T])
	nuevo.dato = dato
	return nuevo
}

func CrearListaEnlazada[T any]() Lista[T] {
	l := new(listaEnlazada[T])
	return l
}
