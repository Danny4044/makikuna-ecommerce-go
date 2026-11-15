package carrito

import "testing"

func TestCalcularSubtotal(t *testing.T) {
	catalogo := map[string]ProductoMinimo{
		"PUL-001": {ID: "PUL-001", PrecioCentavos: 1250, Stock: 5, Activo: true},
		"COL-001": {ID: "COL-001", PrecioCentavos: 2000, Stock: 3, Activo: true},
	}
	items := []Item{{ProductoID: "PUL-001", Cantidad: 2}, {ProductoID: "COL-001", Cantidad: 1}}
	got, err := CalcularSubtotal(items, catalogo)
	if err != nil {
		t.Fatalf("no se esperaba error: %v", err)
	}
	if want := int64(4500); got != want {
		t.Fatalf("subtotal = %d, se esperaba %d", got, want)
	}
}

func TestCalcularSubtotalRechazaStockInsuficiente(t *testing.T) {
	catalogo := map[string]ProductoMinimo{
		"PUL-001": {ID: "PUL-001", PrecioCentavos: 1250, Stock: 1, Activo: true},
	}
	_, err := CalcularSubtotal([]Item{{ProductoID: "PUL-001", Cantidad: 2}}, catalogo)
	if err == nil {
		t.Fatal("se esperaba un error por stock insuficiente")
	}
}
