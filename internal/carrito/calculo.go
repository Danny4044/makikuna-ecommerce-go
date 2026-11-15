package carrito

import "fmt"

// ProductoMinimo contiene los datos necesarios para calcular el carrito.
type ProductoMinimo struct {
	ID             string
	PrecioCentavos int64
	Stock          int
	Activo         bool
}

type Item struct {
	ProductoID string
	Cantidad   int
}

// CalcularSubtotal valida cada artículo y devuelve el subtotal en centavos.
// No modifica los datos recibidos ni accede a servicios externos.
func CalcularSubtotal(items []Item, catalogo map[string]ProductoMinimo) (int64, error) {
	var subtotal int64
	for _, item := range items {
		producto, existe := catalogo[item.ProductoID]
		if !existe {
			return 0, fmt.Errorf("producto %s no encontrado", item.ProductoID)
		}
		if !producto.Activo || item.Cantidad <= 0 {
			return 0, fmt.Errorf("producto o cantidad no válidos")
		}
		if item.Cantidad > producto.Stock {
			return 0, fmt.Errorf("stock insuficiente para %s", item.ProductoID)
		}
		subtotal += producto.PrecioCentavos * int64(item.Cantidad)
	}
	return subtotal, nil
}
