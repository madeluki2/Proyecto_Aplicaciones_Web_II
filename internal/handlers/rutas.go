package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"Proyecto_Aplicaciones_Web_II/internal/models"
	"Proyecto_Aplicaciones_Web_II/internal/storage"
)

// ══════════════════════════════════════════════
// RutaHandler
// ══════════════════════════════════════════════

type RutaHandler struct {
	store storage.RutaStore
}

func NewRutaHandler(s storage.RutaStore) *RutaHandler {
	return &RutaHandler{store: s}
}

// POST /api/v1/rutas
func (h *RutaHandler) Crear(w http.ResponseWriter, r *http.Request) {
	var ruta models.Ruta
	if err := json.NewDecoder(r.Body).Decode(&ruta); err != nil {
		respondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	// Validación básica de campos requeridos
	if ruta.Nombre == "" || ruta.Origen == "" || ruta.Destino == "" {
		respondError(w, http.StatusBadRequest, "nombre, origen y destino son obligatorios")
		return
	}
	if ruta.Estado == "" {
		ruta.Estado = "activo"
	}
	creado, err := h.store.CrearRuta(ruta)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, creado)
}

// GET /api/v1/rutas
func (h *RutaHandler) ObtenerTodos(w http.ResponseWriter, r *http.Request) {
	rutas, err := h.store.ObtenerRutas()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, rutas)
}

// GET /api/v1/rutas/{id}
func (h *RutaHandler) ObtenerUno(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "ID inválido")
		return
	}
	ruta, err := h.store.ObtenerRutaPorID(id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			respondError(w, http.StatusNotFound, "ruta no encontrada")
			return
		}
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, ruta)
}

// PUT /api/v1/rutas/{id}
func (h *RutaHandler) Actualizar(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "ID inválido")
		return
	}
	var ruta models.Ruta
	if err := json.NewDecoder(r.Body).Decode(&ruta); err != nil {
		respondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if ruta.Nombre == "" || ruta.Origen == "" || ruta.Destino == "" {
		respondError(w, http.StatusBadRequest, "nombre, origen y destino son obligatorios")
		return
	}
	actualizado, err := h.store.ActualizarRuta(id, ruta)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			respondError(w, http.StatusNotFound, "ruta no encontrada")
			return
		}
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, actualizado)
}

// DELETE /api/v1/rutas/{id}
func (h *RutaHandler) Eliminar(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "ID inválido")
		return
	}
	if err := h.store.EliminarRuta(id); err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			respondError(w, http.StatusNotFound, "ruta no encontrada")
			return
		}
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"mensaje": "ruta eliminada correctamente"})
}

// ══════════════════════════════════════════════
// PuntoHandler
// ══════════════════════════════════════════════

type PuntoHandler struct {
	store storage.PuntoStore
}

func NewPuntoHandler(s storage.PuntoStore) *PuntoHandler {
	return &PuntoHandler{store: s}
}

// POST /api/v1/puntos
func (h *PuntoHandler) Crear(w http.ResponseWriter, r *http.Request) {
	var punto models.Punto
	if err := json.NewDecoder(r.Body).Decode(&punto); err != nil {
		respondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if punto.Nombre == "" || punto.Direccion == "" || punto.RutaID == 0 {
		respondError(w, http.StatusBadRequest, "nombre, direccion y ruta_id son obligatorios")
		return
	}
	if punto.Estado == "" {
		punto.Estado = "activo"
	}
	creado, err := h.store.CrearPunto(punto)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, creado)
}

// GET /api/v1/puntos
func (h *PuntoHandler) ObtenerTodos(w http.ResponseWriter, r *http.Request) {
	puntos, err := h.store.ObtenerPuntos()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, puntos)
}

// GET /api/v1/puntos/{id}
func (h *PuntoHandler) ObtenerUno(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "ID inválido")
		return
	}
	punto, err := h.store.ObtenerPuntoPorID(id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			respondError(w, http.StatusNotFound, "punto no encontrado")
			return
		}
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, punto)
}

// PUT /api/v1/puntos/{id}
func (h *PuntoHandler) Actualizar(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "ID inválido")
		return
	}
	var punto models.Punto
	if err := json.NewDecoder(r.Body).Decode(&punto); err != nil {
		respondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if punto.Nombre == "" || punto.Direccion == "" || punto.RutaID == 0 {
		respondError(w, http.StatusBadRequest, "nombre, direccion y ruta_id son obligatorios")
		return
	}
	actualizado, err := h.store.ActualizarPunto(id, punto)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			respondError(w, http.StatusNotFound, "punto no encontrado")
			return
		}
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, actualizado)
}

// DELETE /api/v1/puntos/{id}
func (h *PuntoHandler) Eliminar(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "ID inválido")
		return
	}
	if err := h.store.EliminarPunto(id); err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			respondError(w, http.StatusNotFound, "punto no encontrado")
			return
		}
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"mensaje": "punto eliminado correctamente"})
}

// ══════════════════════════════════════════════
// TransportistaHandler
// ══════════════════════════════════════════════

type TransportistaHandler struct {
	store storage.TransportistaStore
}

func NewTransportistaHandler(s storage.TransportistaStore) *TransportistaHandler {
	return &TransportistaHandler{store: s}
}

// POST /api/v1/transportistas
func (h *TransportistaHandler) Crear(w http.ResponseWriter, r *http.Request) {
	var t models.Transportista
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		respondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if t.Nombre == "" || t.PlacaVehiculo == "" {
		respondError(w, http.StatusBadRequest, "nombre y placa_vehiculo son obligatorios")
		return
	}
	if t.Estado == "" {
		t.Estado = "activo"
	}
	creado, err := h.store.CrearTransportista(t)
	if err != nil {
		if errors.Is(err, storage.ErrPlacaDuplicada) {
			respondError(w, http.StatusBadRequest, err.Error())
			return
		}
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, creado)
}

// GET /api/v1/transportistas
func (h *TransportistaHandler) ObtenerTodos(w http.ResponseWriter, r *http.Request) {
	lista, err := h.store.ObtenerTransportistas()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, lista)
}

// GET /api/v1/transportistas/{id}
func (h *TransportistaHandler) ObtenerUno(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "ID inválido")
		return
	}
	t, err := h.store.ObtenerTransportistaPorID(id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			respondError(w, http.StatusNotFound, "transportista no encontrado")
			return
		}
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, t)
}

// PUT /api/v1/transportistas/{id}
func (h *TransportistaHandler) Actualizar(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "ID inválido")
		return
	}
	var t models.Transportista
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		respondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if t.Nombre == "" || t.PlacaVehiculo == "" {
		respondError(w, http.StatusBadRequest, "nombre y placa_vehiculo son obligatorios")
		return
	}
	actualizado, err := h.store.ActualizarTransportista(id, t)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			respondError(w, http.StatusNotFound, "transportista no encontrado")
			return
		}
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, actualizado)
}

// DELETE /api/v1/transportistas/{id}
func (h *TransportistaHandler) Eliminar(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "ID inválido")
		return
	}
	if err := h.store.EliminarTransportista(id); err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			respondError(w, http.StatusNotFound, "transportista no encontrado")
			return
		}
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"mensaje": "transportista eliminado correctamente"})
}

// ══════════════════════════════════════════════
// EntregaHandler
// ══════════════════════════════════════════════

type EntregaHandler struct {
	store storage.EntregaStore
}

func NewEntregaHandler(s storage.EntregaStore) *EntregaHandler {
	return &EntregaHandler{store: s}
}

// POST /api/v1/entregas
func (h *EntregaHandler) Crear(w http.ResponseWriter, r *http.Request) {
	var e models.EntregaPedido
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		respondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if e.PedidoID == 0 || e.PuntoID == 0 || e.TransportistaID == 0 {
		respondError(w, http.StatusBadRequest, "pedido_id, punto_id y transportista_id son obligatorios")
		return
	}
	if e.Estado == "" {
		e.Estado = "pendiente"
	}
	creado, err := h.store.CrearEntrega(e)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, creado)
}

// GET /api/v1/entregas
func (h *EntregaHandler) ObtenerTodos(w http.ResponseWriter, r *http.Request) {
	lista, err := h.store.ObtenerEntregas()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, lista)
}

// GET /api/v1/entregas/{id}
func (h *EntregaHandler) ObtenerUno(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "ID inválido")
		return
	}
	e, err := h.store.ObtenerEntregaPorID(id)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			respondError(w, http.StatusNotFound, "entrega no encontrada")
			return
		}
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, e)
}

// PUT /api/v1/entregas/{id}
func (h *EntregaHandler) Actualizar(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "ID inválido")
		return
	}
	var e models.EntregaPedido
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		respondError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}
	if e.PedidoID == 0 || e.PuntoID == 0 || e.TransportistaID == 0 {
		respondError(w, http.StatusBadRequest, "pedido_id, punto_id y transportista_id son obligatorios")
		return
	}
	actualizado, err := h.store.ActualizarEntrega(id, e)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			respondError(w, http.StatusNotFound, "entrega no encontrada")
			return
		}
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, actualizado)
}

// DELETE /api/v1/entregas/{id}
func (h *EntregaHandler) Eliminar(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "ID inválido")
		return
	}
	if err := h.store.EliminarEntrega(id); err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			respondError(w, http.StatusNotFound, "entrega no encontrada")
			return
		}
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, map[string]string{"mensaje": "entrega eliminada correctamente"})
}

// ──────────────────────────────────────────────
// Helper privado: parsear {id} del path
// ──────────────────────────────────────────────

func parseID(r *http.Request) (uint, error) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}
