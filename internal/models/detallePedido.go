package models

// DetallePedido representa cada producto (especie) dentro de un pedido
type DetallePedido struct {
	ID             int     `json:"id"`
	PedidoID       int     `json:"pedido_id"`
	EspecieID      int     `json:"especie_id"` // referencia externa al módulo de Pesca
	CantidadKg     float64 `json:"cantidad_kg"`
	PrecioUnitario float64 `json:"precio_unitario"`
	Subtotal       float64 `json:"subtotal"`
}
