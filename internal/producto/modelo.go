package producto

// Producto representa un accesorio disponible en el catálogo.
type Producto struct {
	ID             string
	Nombre         string
	Categoria      string
	PrecioCentavos int64
	Stock          int
	Activo         bool
}
