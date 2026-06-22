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

	datos.Id = id

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

// var _ AlmacenPesca = (*SQLitePesca)(nil)
