package storage

import (
	"sync"

	"Proyecto_Aplicaciones_Web_II/internal/models"
)

// Memoria mantiene todos los datos del módulo de Gestión de Pedidos
type Memoria struct {
	clientes      []models.Cliente
	pedidos       []models.Pedido
	detalles      []models.DetallePedido
	nextClienteID int
	nextPedidoID  int
	nextDetalleID int
	mu            sync.Mutex
}

// NewMemoria crea un almacén vacío y listo para usar
func NewMemoria() *Memoria {
	return &Memoria{
		clientes:      []models.Cliente{},
		pedidos:       []models.Pedido{},
		detalles:      []models.DetallePedido{},
		nextClienteID: 1,
		nextPedidoID:  1,
		nextDetalleID: 1,
	}
}

// Seed carga datos iniciales de prueba en memoria
func (m *Memoria) Seed() {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Clientes de prueba
	m.clientes = []models.Cliente{
		{ID: 1, UsuarioID: 1, TipoCliente: "restaurante", NombreNegocio: "Sushi Koi", Direccion: "Av. Flavio Reyes", Telefono: "0991234567", Estado: "activo"},
		{ID: 2, UsuarioID: 2, TipoCliente: "mayorista", NombreNegocio: "Distribuidora El Puerto", Direccion: "Calle 10 de Agosto", Telefono: "0997654321", Estado: "activo"},
	}
	m.nextClienteID = 3

	// Pedidos de prueba
	m.pedidos = []models.Pedido{
		{ID: 1, ClienteID: 1, Fecha: "2026-06-10", Estado: "pendiente", Total: 150.00},
		{ID: 2, ClienteID: 2, Fecha: "2026-06-11", Estado: "en_proceso", Total: 320.50},
	}
	m.nextPedidoID = 3

	// Detalles de prueba
	m.detalles = []models.DetallePedido{
		{ID: 1, PedidoID: 1, EspecieID: 1, CantidadKg: 10.0, PrecioUnitario: 8.50, Subtotal: 85.00},
		{ID: 2, PedidoID: 1, EspecieID: 2, CantidadKg: 5.0, PrecioUnitario: 13.00, Subtotal: 65.00},
	}
	m.nextDetalleID = 3
}

// -------------------- CLIENTES --------------------

// ListarClientes devuelve todos los clientes guardados en memoria
func (m *Memoria) ListarClientes() []models.Cliente {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]models.Cliente, len(m.clientes))
	copy(copia, m.clientes)
	return copia
}

// BuscarClientePorID devuelve el cliente con el ID dado
func (m *Memoria) BuscarClientePorID(id int) (models.Cliente, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, c := range m.clientes {
		if c.ID == id {
			return c, true
		}
	}
	return models.Cliente{}, false
}

// CrearCliente agrega un nuevo cliente con ID incremental
func (m *Memoria) CrearCliente(cliente models.Cliente) models.Cliente {
	m.mu.Lock()
	defer m.mu.Unlock()
	cliente.ID = m.nextClienteID
	m.clientes = append(m.clientes, cliente)
	m.nextClienteID++
	return cliente
}

// ActualizarCliente reemplaza los datos del cliente con el ID dado
func (m *Memoria) ActualizarCliente(id int, cliente models.Cliente) (models.Cliente, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, c := range m.clientes {
		if c.ID == id {
			cliente.ID = c.ID
			if cliente.TipoCliente == "" {
				cliente.TipoCliente = c.TipoCliente
			}
			if cliente.NombreNegocio == "" {
				cliente.NombreNegocio = c.NombreNegocio
			}
			if cliente.Direccion == "" {
				cliente.Direccion = c.Direccion
			}
			if cliente.Telefono == "" {
				cliente.Telefono = c.Telefono
			}
			if cliente.Estado == "" {
				cliente.Estado = c.Estado
			}
			m.clientes[i] = cliente
			return cliente, true
		}
	}
	return models.Cliente{}, false
}

// EliminarCliente remueve el cliente con el ID dado
func (m *Memoria) EliminarCliente(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, c := range m.clientes {
		if c.ID == id {
			m.clientes = append(m.clientes[:i], m.clientes[i+1:]...)
			return true
		}
	}
	return false
}

// -------------------- PEDIDOS --------------------

// ListarPedidos devuelve todos los pedidos guardados en memoria
func (m *Memoria) ListarPedidos() []models.Pedido {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]models.Pedido, len(m.pedidos))
	copy(copia, m.pedidos)
	return copia
}

// BuscarPedidoPorID devuelve el pedido con el ID dado
func (m *Memoria) BuscarPedidoPorID(id int) (models.Pedido, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, p := range m.pedidos {
		if p.ID == id {
			return p, true
		}
	}
	return models.Pedido{}, false
}

// CrearPedido agrega un nuevo pedido con ID incremental
func (m *Memoria) CrearPedido(pedido models.Pedido) models.Pedido {
	m.mu.Lock()
	defer m.mu.Unlock()
	pedido.ID = m.nextPedidoID
	m.pedidos = append(m.pedidos, pedido)
	m.nextPedidoID++
	return pedido
}

// ActualizarPedido reemplaza los datos del pedido con el ID dado
func (m *Memoria) ActualizarPedido(id int, pedido models.Pedido) (models.Pedido, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, p := range m.pedidos {
		if p.ID == id {
			pedido.ID = p.ID
			if pedido.Estado == "" {
				pedido.Estado = p.Estado
			}
			if pedido.Fecha == "" {
				pedido.Fecha = p.Fecha
			}
			if pedido.Total == 0 {
				pedido.Total = p.Total
			}
			if pedido.ClienteID == 0 {
				pedido.ClienteID = p.ClienteID
			}
			m.pedidos[i] = pedido
			return pedido, true
		}
	}
	return models.Pedido{}, false
}

// EliminarPedido remueve el pedido con el ID dado
func (m *Memoria) EliminarPedido(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, p := range m.pedidos {
		if p.ID == id {
			m.pedidos = append(m.pedidos[:i], m.pedidos[i+1:]...)
			return true
		}
	}
	return false
}

// -------------------- DETALLES DE PEDIDO --------------------

// ListarDetalles devuelve todos los detalles guardados en memoria
func (m *Memoria) ListarDetalles() []models.DetallePedido {
	m.mu.Lock()
	defer m.mu.Unlock()
	copia := make([]models.DetallePedido, len(m.detalles))
	copy(copia, m.detalles)
	return copia
}

// BuscarDetallePorID devuelve el detalle con el ID dado
func (m *Memoria) BuscarDetallePorID(id int) (models.DetallePedido, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, d := range m.detalles {
		if d.ID == id {
			return d, true
		}
	}
	return models.DetallePedido{}, false
}

// CrearDetalle agrega un nuevo detalle de pedido con ID incremental
func (m *Memoria) CrearDetalle(detalle models.DetallePedido) models.DetallePedido {
	m.mu.Lock()
	defer m.mu.Unlock()
	detalle.ID = m.nextDetalleID
	m.detalles = append(m.detalles, detalle)
	m.nextDetalleID++
	return detalle
}

// ActualizarDetalle reemplaza los datos del detalle con el ID dado
func (m *Memoria) ActualizarDetalle(id int, detalle models.DetallePedido) (models.DetallePedido, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, d := range m.detalles {
		if d.ID == id {
			detalle.ID = d.ID
			if detalle.CantidadKg == 0 {
				detalle.CantidadKg = d.CantidadKg
			}
			if detalle.PrecioUnitario == 0 {
				detalle.PrecioUnitario = d.PrecioUnitario
			}
			if detalle.Subtotal == 0 {
				detalle.Subtotal = d.Subtotal
			}
			if detalle.EspecieID == 0 {
				detalle.EspecieID = d.EspecieID
			}
			if detalle.PedidoID == 0 {
				detalle.PedidoID = d.PedidoID
			}
			m.detalles[i] = detalle
			return detalle, true
		}
	}
	return models.DetallePedido{}, false
}

// EliminarDetalle remueve el detalle con el ID dado
func (m *Memoria) EliminarDetalle(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, d := range m.detalles {
		if d.ID == id {
			m.detalles = append(m.detalles[:i], m.detalles[i+1:]...)
			return true
		}
	}
	return false
}
