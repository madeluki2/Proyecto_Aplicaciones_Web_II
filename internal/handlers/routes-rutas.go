package handlers

import (
	"github.com/go-chi/chi/v5"

	"Proyecto_Aplicaciones_Web_II/internal/storage"
)

// RegistrarRutas monta todos los subrouters del módulo Rutas de Distribución
// sobre el router Chi que recibe como parámetro.
// Se llama desde main.go pasando el router principal.
func RegistrarRutas(r chi.Router) {

	// ── Instanciar stores ─────────────────────────
	rutaStore := storage.NewMemRutaStore()
	puntoStore := storage.NewMemPuntoStore()
	transportistaStore := storage.NewMemTransportistaStore()
	entregaStore := storage.NewMemEntregaStore()

	// ── Instanciar handlers ───────────────────────
	rutaHandler := NewRutaHandler(rutaStore)
	puntoHandler := NewPuntoHandler(puntoStore)
	transportistaHandler := NewTransportistaHandler(transportistaStore)
	entregaHandler := NewEntregaHandler(entregaStore)

	// ── Subrouters ────────────────────────────────

	// /api/v1/rutas
	r.Route("/rutas", func(r chi.Router) {
		r.Post("/", rutaHandler.Crear)
		r.Get("/", rutaHandler.ObtenerTodos)
		r.Get("/{id}", rutaHandler.ObtenerUno)
		r.Put("/{id}", rutaHandler.Actualizar)
		r.Delete("/{id}", rutaHandler.Eliminar)
	})

	// /api/v1/puntos
	r.Route("/puntos", func(r chi.Router) {
		r.Post("/", puntoHandler.Crear)
		r.Get("/", puntoHandler.ObtenerTodos)
		r.Get("/{id}", puntoHandler.ObtenerUno)
		r.Put("/{id}", puntoHandler.Actualizar)
		r.Delete("/{id}", puntoHandler.Eliminar)
	})

	// /api/v1/transportistas
	r.Route("/transportistas", func(r chi.Router) {
		r.Post("/", transportistaHandler.Crear)
		r.Get("/", transportistaHandler.ObtenerTodos)
		r.Get("/{id}", transportistaHandler.ObtenerUno)
		r.Put("/{id}", transportistaHandler.Actualizar)
		r.Delete("/{id}", transportistaHandler.Eliminar)
	})

	// /api/v1/entregas
	r.Route("/entregas", func(r chi.Router) {
		r.Post("/", entregaHandler.Crear)
		r.Get("/", entregaHandler.ObtenerTodos)
		r.Get("/{id}", entregaHandler.ObtenerUno)
		r.Put("/{id}", entregaHandler.Actualizar)
		r.Delete("/{id}", entregaHandler.Eliminar)
	})
}
