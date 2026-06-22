package storage

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"Proyecto_Aplicaciones_Web_II/internal/models"
)

func NuevaConexionSQLite() (*gorm.DB, error) {

	db, err := gorm.Open(
		sqlite.Open("pesca.db"),
		&gorm.Config{},
	)

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.Usuario{},
		&models.Pescador{},
		&models.Embarcacion{},
		&models.Especie{},
		&models.Captura{},
		&models.Bodega{},
		&models.Stock{},
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}
