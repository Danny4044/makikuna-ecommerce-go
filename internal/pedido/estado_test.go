package pedido

import "testing"

func TestPuedeCambiarEstado(t *testing.T) {
	if !PuedeCambiarEstado(Pendiente, Pagado) {
		t.Fatal("pendiente -> pagado debería estar permitido")
	}
	if PuedeCambiarEstado(Entregado, Pendiente) {
		t.Fatal("entregado -> pendiente no debería estar permitido")
	}
}
