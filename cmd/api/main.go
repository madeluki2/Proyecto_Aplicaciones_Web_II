package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"Proyecto_Aplicaciones_Web_II/internal/handlers"
)

func main() {
	// ── Router principal ──────────────────────────
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// ── Montar módulos bajo /api/v1 ───────────────
	r.Route("/api/v1", func(r chi.Router) {
		handlers.RegistrarRutas(r) // ← módulo rutas de distribución
		// handlers.RegistrarPescas(r)  ← otros módulos de tus compañeros
		// handlers.RegistrarPedidos(r)
	})

	// ── Arrancar servidor ─────────────────────────
	fmt.Println("Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
