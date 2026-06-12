package storage

import (
	"sync"
	"time"

	"Proyecto_Aplicaciones_Web_II/internal/models"
)

// ══════════════════════════════════════════════
// IMPLEMENTACIONES EN MEMORIA
// ══════════════════════════════════════════════

// ── Rutas ─────────────────────────────────────

type MemRutaStore struct {
	mu     sync.RWMutex
	rutas  []models.Ruta
	nextID uint
}

func NewMemRutaStore() *MemRutaStore {
	return &MemRutaStore{nextID: 1}
}

func (s *MemRutaStore) CrearRuta(r models.Ruta) (models.Ruta, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	r.ID = s.nextID
	r.CreadoEn = time.Now()
	s.nextID++
	s.rutas = append(s.rutas, r)
	return r, nil
}

func (s *MemRutaStore) ObtenerRutas() ([]models.Ruta, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.rutas, nil
}

func (s *MemRutaStore) ObtenerRutaPorID(id uint) (models.Ruta, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, r := range s.rutas {
		if r.ID == id {
			return r, nil
		}
	}
	return models.Ruta{}, ErrNotFound
}

func (s *MemRutaStore) ActualizarRuta(id uint, nuevo models.Ruta) (models.Ruta, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, r := range s.rutas {
		if r.ID == id {
			nuevo.ID = id
			nuevo.CreadoEn = r.CreadoEn
			s.rutas[i] = nuevo
			return nuevo, nil
		}
	}
	return models.Ruta{}, ErrNotFound
}

func (s *MemRutaStore) EliminarRuta(id uint) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, r := range s.rutas {
		if r.ID == id {
			s.rutas = append(s.rutas[:i], s.rutas[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

// ── Puntos ────────────────────────────────────

type MemPuntoStore struct {
	mu     sync.RWMutex
	puntos []models.Punto
	nextID uint
}

func NewMemPuntoStore() *MemPuntoStore {
	return &MemPuntoStore{nextID: 1}
}

func (s *MemPuntoStore) CrearPunto(p models.Punto) (models.Punto, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	p.ID = s.nextID
	p.CreadoEn = time.Now()
	s.nextID++
	s.puntos = append(s.puntos, p)
	return p, nil
}

func (s *MemPuntoStore) ObtenerPuntos() ([]models.Punto, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.puntos, nil
}

func (s *MemPuntoStore) ObtenerPuntoPorID(id uint) (models.Punto, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, p := range s.puntos {
		if p.ID == id {
			return p, nil
		}
	}
	return models.Punto{}, ErrNotFound
}

func (s *MemPuntoStore) ActualizarPunto(id uint, nuevo models.Punto) (models.Punto, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, p := range s.puntos {
		if p.ID == id {
			nuevo.ID = id
			nuevo.CreadoEn = p.CreadoEn
			s.puntos[i] = nuevo
			return nuevo, nil
		}
	}
	return models.Punto{}, ErrNotFound
}

func (s *MemPuntoStore) EliminarPunto(id uint) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, p := range s.puntos {
		if p.ID == id {
			s.puntos = append(s.puntos[:i], s.puntos[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

// ── Transportistas ────────────────────────────

type MemTransportistaStore struct {
	mu             sync.RWMutex
	transportistas []models.Transportista
	nextID         uint
}

func NewMemTransportistaStore() *MemTransportistaStore {
	return &MemTransportistaStore{nextID: 1}
}

func (s *MemTransportistaStore) CrearTransportista(t models.Transportista) (models.Transportista, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	// Validar placa única
	for _, existing := range s.transportistas {
		if existing.PlacaVehiculo == t.PlacaVehiculo {
			return models.Transportista{}, ErrPlacaDuplicada
		}
	}
	t.ID = s.nextID
	t.CreadoEn = time.Now()
	s.nextID++
	s.transportistas = append(s.transportistas, t)
	return t, nil
}

func (s *MemTransportistaStore) ObtenerTransportistas() ([]models.Transportista, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.transportistas, nil
}

func (s *MemTransportistaStore) ObtenerTransportistaPorID(id uint) (models.Transportista, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, t := range s.transportistas {
		if t.ID == id {
			return t, nil
		}
	}
	return models.Transportista{}, ErrNotFound
}

func (s *MemTransportistaStore) ActualizarTransportista(id uint, nuevo models.Transportista) (models.Transportista, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, t := range s.transportistas {
		if t.ID == id {
			nuevo.ID = id
			nuevo.CreadoEn = t.CreadoEn
			s.transportistas[i] = nuevo
			return nuevo, nil
		}
	}
	return models.Transportista{}, ErrNotFound
}

func (s *MemTransportistaStore) EliminarTransportista(id uint) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, t := range s.transportistas {
		if t.ID == id {
			s.transportistas = append(s.transportistas[:i], s.transportistas[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

// ── Entregas ──────────────────────────────────

type MemEntregaStore struct {
	mu       sync.RWMutex
	entregas []models.EntregaPedido
	nextID   uint
}

func NewMemEntregaStore() *MemEntregaStore {
	return &MemEntregaStore{nextID: 1}
}

func (s *MemEntregaStore) CrearEntrega(e models.EntregaPedido) (models.EntregaPedido, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	e.ID = s.nextID
	e.CreadoEn = time.Now()
	s.nextID++
	s.entregas = append(s.entregas, e)
	return e, nil
}

func (s *MemEntregaStore) ObtenerEntregas() ([]models.EntregaPedido, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.entregas, nil
}

func (s *MemEntregaStore) ObtenerEntregaPorID(id uint) (models.EntregaPedido, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, e := range s.entregas {
		if e.ID == id {
			return e, nil
		}
	}
	return models.EntregaPedido{}, ErrNotFound
}

func (s *MemEntregaStore) ActualizarEntrega(id uint, nuevo models.EntregaPedido) (models.EntregaPedido, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, e := range s.entregas {
		if e.ID == id {
			nuevo.ID = id
			nuevo.CreadoEn = e.CreadoEn
			s.entregas[i] = nuevo
			return nuevo, nil
		}
	}
	return models.EntregaPedido{}, ErrNotFound
}

func (s *MemEntregaStore) EliminarEntrega(id uint) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, e := range s.entregas {
		if e.ID == id {
			s.entregas = append(s.entregas[:i], s.entregas[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}
