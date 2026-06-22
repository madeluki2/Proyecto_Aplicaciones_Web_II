package storage

import (
	"Proyecto_Aplicaciones_Web_II/internal/models"

	"gorm.io/gorm"
)

type SQLitePesca struct {
	db *gorm.DB
}

func NuevaSQLitePesca(db *gorm.DB) *SQLitePesca {
	return &SQLitePesca{
		db: db,
	}
}

// Usuarios
func (s *SQLitePesca) ListarUsuarios() []models.Usuario {
	var usuarios []models.Usuario
	s.db.Find(&usuarios)
	return usuarios
}

func (s *SQLitePesca) BuscarUsuarioPorID(id int) (models.Usuario, bool) {
	var usuario models.Usuario
	result := s.db.First(&usuario, id)

	if result.Error != nil {
		return models.Usuario{}, false
	}

	return usuario, true
}

func (s *SQLitePesca) CrearUsuario(
	u models.Usuario,
) models.Usuario {
	s.db.Create(&u)
	return u
}

func (s *SQLitePesca) ActualizarUsuario(
	id int,
	datos models.Usuario,
) (models.Usuario, bool) {

	var usuario models.Usuario

	result := s.db.First(&usuario, id)

	if result.Error != nil {
		return models.Usuario{}, false
	}

	datos.ID = id

	s.db.Save(&datos)

	return datos, true
}

func (s *SQLitePesca) BorrarUsuario(id int) bool {

	result := s.db.Delete(
		&models.Usuario{},
		id,
	)

	return result.RowsAffected > 0
}

// --------------------------- Pescadores --------------------------

func (s *SQLitePesca) ListarPescadores() []models.Pescador {
	var pescadores []models.Pescador
	s.db.Find(&pescadores)
	return pescadores
}

func (s *SQLitePesca) BuscarPescadorPorID(id int) (models.Pescador, bool) {
	var pescador models.Pescador

	result := s.db.First(&pescador, id)

	if result.Error != nil {
		return models.Pescador{}, false
	}

	return pescador, true
}

func (s *SQLitePesca) CrearPescador(
	p models.Pescador,
) models.Pescador {

	s.db.Create(&p)

	return p
}

func (s *SQLitePesca) ActualizarPescador(
	id int,
	datos models.Pescador,
) (models.Pescador, bool) {

	var pescador models.Pescador

	result := s.db.First(&pescador, id)

	if result.Error != nil {
		return models.Pescador{}, false
	}

	datos.ID = id

	s.db.Save(&datos)

	return datos, true
}

func (s *SQLitePesca) BorrarPescador(id int) bool {

	result := s.db.Delete(
		&models.Pescador{},
		id,
	)

	return result.RowsAffected > 0
}

// --------------------------- Embarcaciones --------------------------

func (s *SQLitePesca) ListarEmbarcaciones() []models.Embarcacion {
	var embarcaciones []models.Embarcacion
	s.db.Find(&embarcaciones)
	return embarcaciones
}

func (s *SQLitePesca) BuscarEmbarcacionPorID(id int) (models.Embarcacion, bool) {
	var embarcacion models.Embarcacion

	result := s.db.First(&embarcacion, id)

	if result.Error != nil {
		return models.Embarcacion{}, false
	}

	return embarcacion, true
}

func (s *SQLitePesca) CrearEmbarcacion(
	e models.Embarcacion,
) models.Embarcacion {

	s.db.Create(&e)

	return e
}

func (s *SQLitePesca) ActualizarEmbarcacion(
	id int,
	datos models.Embarcacion,
) (models.Embarcacion, bool) {

	var embarcacion models.Embarcacion

	result := s.db.First(&embarcacion, id)

	if result.Error != nil {
		return models.Embarcacion{}, false
	}

	datos.ID = id

	s.db.Save(&datos)

	return datos, true
}

func (s *SQLitePesca) BorrarEmbarcacion(id int) bool {

	result := s.db.Delete(
		&models.Embarcacion{},
		id,
	)

	return result.RowsAffected > 0
}

// --------------------------- Especies --------------------------

func (s *SQLitePesca) ListarEspecies() []models.Especie {
	var especies []models.Especie
	s.db.Find(&especies)
	return especies
}

func (s *SQLitePesca) BuscarEspeciePorID(id int) (models.Especie, bool) {
	var especie models.Especie

	result := s.db.First(&especie, id)

	if result.Error != nil {
		return models.Especie{}, false
	}

	return especie, true
}

func (s *SQLitePesca) CrearEspecie(
	e models.Especie,
) models.Especie {

	s.db.Create(&e)

	return e
}

func (s *SQLitePesca) ActualizarEspecie(
	id int,
	datos models.Especie,
) (models.Especie, bool) {

	var especie models.Especie

	result := s.db.First(&especie, id)

	if result.Error != nil {
		return models.Especie{}, false
	}

	datos.ID = id

	s.db.Save(&datos)

	return datos, true
}

func (s *SQLitePesca) BorrarEspecie(id int) bool {

	result := s.db.Delete(
		&models.Especie{},
		id,
	)

	return result.RowsAffected > 0
}

func (s *SQLitePesca) ListarCapturas() []models.Captura {
	var capturas []models.Captura
	s.db.Find(&capturas)
	return capturas
}

func (s *SQLitePesca) BuscarCapturaPorID(id int) (models.Captura, bool) {
	var captura models.Captura

	result := s.db.First(&captura, id)

	if result.Error != nil {
		return models.Captura{}, false
	}

	return captura, true
}

func (s *SQLitePesca) CrearCaptura(c models.Captura) models.Captura {
	s.db.Create(&c)
	return c
}

func (s *SQLitePesca) ActualizarCaptura(id int, datos models.Captura) (models.Captura, bool) {

	var captura models.Captura

	result := s.db.First(&captura, id)

	if result.Error != nil {
		return models.Captura{}, false
	}

	datos.ID = id

	s.db.Save(&datos)

	return datos, true
}

func (s *SQLitePesca) BorrarCaptura(id int) bool {

	result := s.db.Delete(
		&models.Captura{},
		id,
	)

	return result.RowsAffected > 0
}

// --------------------------- Bodegas --------------------------

func (s *SQLitePesca) ListarBodegas() []models.Bodega {
	var bodegas []models.Bodega
	s.db.Find(&bodegas)
	return bodegas
}

func (s *SQLitePesca) BuscarBodegaPorID(id int) (models.Bodega, bool) {
	var bodega models.Bodega

	result := s.db.First(&bodega, id)

	if result.Error != nil {
		return models.Bodega{}, false
	}

	return bodega, true
}

func (s *SQLitePesca) CrearBodega(
	b models.Bodega,
) models.Bodega {

	s.db.Create(&b)

	return b
}

func (s *SQLitePesca) ActualizarBodega(
	id int,
	datos models.Bodega,
) (models.Bodega, bool) {

	var bodega models.Bodega

	result := s.db.First(&bodega, id)

	if result.Error != nil {
		return models.Bodega{}, false
	}

	datos.ID = id

	s.db.Save(&datos)

	return datos, true
}

func (s *SQLitePesca) BorrarBodega(id int) bool {

	result := s.db.Delete(
		&models.Bodega{},
		id,
	)

	return result.RowsAffected > 0
}

// --------------------------- Stocks --------------------------

func (s *SQLitePesca) ListarStocks() []models.Stock {
	var stocks []models.Stock
	s.db.Find(&stocks)
	return stocks
}

func (s *SQLitePesca) BuscarStockPorID(id int) (models.Stock, bool) {
	var stock models.Stock

	result := s.db.First(&stock, id)

	if result.Error != nil {
		return models.Stock{}, false
	}

	return stock, true
}

func (s *SQLitePesca) CrearStock(
	st models.Stock,
) models.Stock {

	s.db.Create(&st)

	return st
}

func (s *SQLitePesca) ActualizarStock(
	id int,
	datos models.Stock,
) (models.Stock, bool) {

	var stock models.Stock

	result := s.db.First(&stock, id)

	if result.Error != nil {
		return models.Stock{}, false
	}

	datos.ID = id

	s.db.Save(&datos)

	return datos, true
}

func (s *SQLitePesca) BorrarStock(id int) bool {

	result := s.db.Delete(
		&models.Stock{},
		id,
	)

	return result.RowsAffected > 0
}

var _ AlmacenPesca = (*SQLitePesca)(nil)
