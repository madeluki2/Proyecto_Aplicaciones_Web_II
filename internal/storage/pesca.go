package storage

import (
	"sync"

	"Proyecto_API/internal/models"
)

// MemoriaPesca guarda todos los datos en RAM mientras la aplicación esté en ejecución. No es persistente.
type MemoriaPesca struct {
	usuarios      []models.Usuario
	nextUsuarioID int

	pescadores     []models.Pescador
	nextPescadorID int

	embarcaciones     []models.Embarcacion
	nextEmbarcacionID int

	capturas      []models.Captura
	nextCapturaID int

	bodegas      []models.Bodega
	nextBodegaID int

	stocks      []models.Stock
	nextStockID int

	especies      []models.Especie
	nextEspecieID int

	mu sync.Mutex
}

// NuevaMemoriaPesca crea una nueva instancia de MemoriaPesca con datos iniciales vacíos, comenzando los IDs en 1.
func NuevaMemoriaPesca() *MemoriaPesca {
	return &MemoriaPesca{
		usuarios:          []models.Usuario{},
		nextUsuarioID:     1,
		pescadores:        []models.Pescador{},
		nextPescadorID:    1,
		embarcaciones:     []models.Embarcacion{},
		nextEmbarcacionID: 1,
		capturas:          []models.Captura{},
		nextCapturaID:     1,
		bodegas:           []models.Bodega{},
		nextBodegaID:      1,
		stocks:            []models.Stock{},
		nextStockID:       1,
		especies:          []models.Especie{},
		nextEspecieID:     1,
	}
}
