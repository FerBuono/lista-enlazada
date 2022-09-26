package lista_test

import (
	TDAlista "lista"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListaVacia(t *testing.T) {
	lista := TDAlista.CrearListaEnlazada[int]()
	require.NotNil(t, lista)
}
