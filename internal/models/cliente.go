package models

// Cliente representa a un restaurante, intermediario o mayorista que realiza pedidos
type Cliente struct {
	ID            int    `json:"id"`
	UsuarioID     int    `json:"usuario_id"`
	TipoCliente   string `json:"tipo_cliente"` // restaurante / intermediario / mayorista
	NombreNegocio string `json:"nombre_negocio"`
	Direccion     string `json:"direccion"`
	Telefono      string `json:"telefono"`
	Estado        string `json:"estado"`
}
