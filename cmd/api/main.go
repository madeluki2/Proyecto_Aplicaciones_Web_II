package main

import (
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
}
