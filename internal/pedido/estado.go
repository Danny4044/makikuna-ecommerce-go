package pedido

// Estado identifica la situación de un pedido.
type Estado string

const (
	Pendiente   Estado = "pendiente"
	Pagado      Estado = "pagado"
	Preparando  Estado = "preparando"
	Enviado     Estado = "enviado"
	Entregado   Estado = "entregado"
	Cancelado   Estado = "cancelado"
)

// PuedeCambiarEstado valida una transición permitida del flujo del pedido.
func PuedeCambiarEstado(actual, siguiente Estado) bool {
	switch actual {
	case Pendiente:
		return siguiente == Pagado || siguiente == Cancelado
	case Pagado:
		return siguiente == Preparando || siguiente == Cancelado
	case Preparando:
		return siguiente == Enviado
	case Enviado:
		return siguiente == Entregado
	default:
		return false
	}
}
