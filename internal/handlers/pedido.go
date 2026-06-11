package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"

	"Proyecto_Aplicaciones_Web_II/internal/models"
	"Proyecto_Aplicaciones_Web_II/internal/storage"
)

// Server unifica todos los handlers del módulo de Gestión de Pedidos
type Server struct {
	storage *storage.Memoria
}

// NewServer es el constructor que inyecta el almacén en el servidor
func NewServer(s *storage.Memoria) *Server {
	return &Server{storage: s}
}

// -------------------- CLIENTES --------------------

// ListarClientes devuelve todos los clientes registrados (GET)
func (s *Server) ListarClientes(w http.ResponseWriter, _ *http.Request) {
	clientes := s.storage.ListarClientes()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(clientes); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// ObtenerCliente devuelve un cliente por su ID (GET)
func (s *Server) ObtenerCliente(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	// Convertimos el ID de texto a entero
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "el id debe ser un número entero", http.StatusBadRequest)
		return
	}

	cliente, encontrado := s.storage.BuscarClientePorID(id)
	if !encontrado {
		http.Error(w, "cliente no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(cliente); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// CrearCliente registra un nuevo cliente (POST)
func (s *Server) CrearCliente(w http.ResponseWriter, r *http.Request) {
	var nuevo models.Cliente
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validación de campos obligatorios
	if strings.TrimSpace(nuevo.NombreNegocio) == "" || strings.TrimSpace(nuevo.TipoCliente) == "" || strings.TrimSpace(nuevo.Telefono) == "" || strings.TrimSpace(nuevo.Direccion) == "" {
		http.Error(w, "NombreNegocio, TipoCliente, Telefono y Direccion son obligatorios", http.StatusBadRequest)
		return
	}

	// El estado siempre inicia como activo
	nuevo.Estado = "activo"

	nuevo = s.storage.CrearCliente(nuevo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(nuevo); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// ActualizarCliente modifica los datos de un cliente existente (PUT)
func (s *Server) ActualizarCliente(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "el id debe ser un número entero", http.StatusBadRequest)
		return
	}

	var datos models.Cliente
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validamos que al menos un campo venga para actualizar
	if strings.TrimSpace(datos.NombreNegocio) == "" && strings.TrimSpace(datos.TipoCliente) == "" && strings.TrimSpace(datos.Telefono) == "" && strings.TrimSpace(datos.Direccion) == "" && strings.TrimSpace(datos.Estado) == "" {
		http.Error(w, "debe enviar al menos un campo para actualizar", http.StatusBadRequest)
		return
	}

	actualizado, encontrado := s.storage.ActualizarCliente(id, datos)
	if !encontrado {
		http.Error(w, "cliente no encontrado para actualizar", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(actualizado); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// EliminarCliente remueve un cliente por su ID (DELETE)
func (s *Server) EliminarCliente(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "el id debe ser un número entero", http.StatusBadRequest)
		return
	}

	seBorro := s.storage.EliminarCliente(id)
	if !seBorro {
		http.Error(w, "cliente no encontrado para eliminar", http.StatusNotFound)
		return
	}

	// 204 No Content: eliminación exitosa sin cuerpo de respuesta
	w.WriteHeader(http.StatusNoContent)
}

// -------------------- PEDIDOS --------------------

// ListarPedidos devuelve todos los pedidos registrados (GET)
func (s *Server) ListarPedidos(w http.ResponseWriter, _ *http.Request) {
	pedidos := s.storage.ListarPedidos()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(pedidos); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// ObtenerPedido devuelve un pedido por su ID (GET)
func (s *Server) ObtenerPedido(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	// Convertimos el ID de texto a entero
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "el id debe ser un número entero", http.StatusBadRequest)
		return
	}

	pedido, encontrado := s.storage.BuscarPedidoPorID(id)
	if !encontrado {
		http.Error(w, "pedido no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(pedido); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// CrearPedido registra un nuevo pedido (POST)
func (s *Server) CrearPedido(w http.ResponseWriter, r *http.Request) {
	var nuevo models.Pedido
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validación de campos obligatorios
	if nuevo.ClienteID == 0 || strings.TrimSpace(nuevo.Fecha) == "" {
		http.Error(w, "ClienteID y Fecha son obligatorios", http.StatusBadRequest)
		return
	}

	// El estado siempre inicia como pendiente
	nuevo.Estado = "pendiente"
	// El total inicia en cero hasta agregar detalles
	nuevo.Total = 0

	nuevo = s.storage.CrearPedido(nuevo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(nuevo); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// ActualizarPedido modifica los datos de un pedido existente (PUT)
func (s *Server) ActualizarPedido(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "el id debe ser un número entero", http.StatusBadRequest)
		return
	}

	var datos models.Pedido
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validamos que al menos un campo venga para actualizar
	if strings.TrimSpace(datos.Estado) == "" && strings.TrimSpace(datos.Fecha) == "" && datos.Total == 0 {
		http.Error(w, "debe enviar al menos un campo para actualizar", http.StatusBadRequest)
		return
	}

	actualizado, encontrado := s.storage.ActualizarPedido(id, datos)
	if !encontrado {
		http.Error(w, "pedido no encontrado para actualizar", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(actualizado); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// EliminarPedido cancela y remueve un pedido por su ID (DELETE)
func (s *Server) EliminarPedido(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "el id debe ser un número entero", http.StatusBadRequest)
		return
	}

	seBorro := s.storage.EliminarPedido(id)
	if !seBorro {
		http.Error(w, "pedido no encontrado para eliminar", http.StatusNotFound)
		return
	}

	// 204 No Content: eliminación exitosa sin cuerpo de respuesta
	w.WriteHeader(http.StatusNoContent)
}

// -------------------- DETALLES DE PEDIDO --------------------

// ListarDetalles devuelve todos los detalles de pedido registrados (GET)
func (s *Server) ListarDetalles(w http.ResponseWriter, _ *http.Request) {
	detalles := s.storage.ListarDetalles()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(detalles); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// ObtenerDetalle devuelve un detalle de pedido por su ID (GET)
func (s *Server) ObtenerDetalle(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	// Convertimos el ID de texto a entero
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "el id debe ser un número entero", http.StatusBadRequest)
		return
	}

	detalle, encontrado := s.storage.BuscarDetallePorID(id)
	if !encontrado {
		http.Error(w, "detalle de pedido no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(detalle); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// CrearDetalle registra un nuevo detalle dentro de un pedido (POST)
func (s *Server) CrearDetalle(w http.ResponseWriter, r *http.Request) {
	var nuevo models.DetallePedido
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validación de campos obligatorios
	if nuevo.PedidoID == 0 || nuevo.EspecieID == 0 || nuevo.CantidadKg == 0 || nuevo.PrecioUnitario == 0 {
		http.Error(w, "PedidoID, EspecieID, CantidadKg y PrecioUnitario son obligatorios", http.StatusBadRequest)
		return
	}

	// Calculamos el subtotal automáticamente
	nuevo.Subtotal = nuevo.CantidadKg * nuevo.PrecioUnitario

	nuevo = s.storage.CrearDetalle(nuevo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(nuevo); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// ActualizarDetalle modifica los datos de un detalle de pedido existente (PUT)
func (s *Server) ActualizarDetalle(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "el id debe ser un número entero", http.StatusBadRequest)
		return
	}

	var datos models.DetallePedido
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validamos que al menos un campo venga para actualizar
	if datos.CantidadKg == 0 && datos.PrecioUnitario == 0 && datos.EspecieID == 0 {
		http.Error(w, "debe enviar al menos un campo para actualizar", http.StatusBadRequest)
		return
	}

	// Recalculamos el subtotal si vienen los dos campos necesarios
	if datos.CantidadKg != 0 && datos.PrecioUnitario != 0 {
		datos.Subtotal = datos.CantidadKg * datos.PrecioUnitario
	}

	actualizado, encontrado := s.storage.ActualizarDetalle(id, datos)
	if !encontrado {
		http.Error(w, "detalle de pedido no encontrado para actualizar", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(actualizado); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// EliminarDetalle remueve un detalle de pedido por su ID (DELETE)
func (s *Server) EliminarDetalle(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "el id debe ser un número entero", http.StatusBadRequest)
		return
	}

	seBorro := s.storage.EliminarDetalle(id)
	if !seBorro {
		http.Error(w, "detalle de pedido no encontrado para eliminar", http.StatusNotFound)
		return
	}

	// 204 No Content: eliminación exitosa sin cuerpo de respuesta
	w.WriteHeader(http.StatusNoContent)
}
