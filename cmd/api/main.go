package main

import (
<<<<<<< HEAD
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
=======
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"

	"Proyecto_Aplicaciones_Web_II/internal/handlers"
	"Proyecto_Aplicaciones_Web_II/internal/storage"
)

func main() {

	// 1. Almacén en memoria (Modulo Gestión de Pesca)
	storePesca := storage.NuevaMemoriaPesca()
	// 2. Router principal
	r := chi.NewRouter()
	// 3. Middleware global
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	// 4. Rutas versionadas /api/v1
	r.Route("/api/v1", func(r chi.Router) {

		// Handlers gestion de pesca
		handlers.MontarRutasPesca(r, storePesca)

	})

	// 5. Iniciar servidor
	log.Println("Servidor iniciado en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
>>>>>>> feature/gestion-pesca
}
