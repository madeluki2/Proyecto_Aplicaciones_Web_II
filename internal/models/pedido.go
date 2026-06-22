package models

// Pedido representa una solicitud de compra realizada por un cliente
type Pedido struct {
	ID        int     `json:"id"`
	ClienteID int     `json:"cliente_id"`
	Fecha     string  `json:"fecha"`
	Estado    string  `json:"estado"` // pendiente / en_proceso / entregado / cancelado
	Total     float64 `json:"total"`
}
