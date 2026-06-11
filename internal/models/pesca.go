package models

type Especie struct {
	ID               int    `json:"id"`
	NombreComun      string `json:"nombre_comun"`
	NombreCientifico string `json:"nombre_cientifico"`
	UnidadMedida     string `json:"unidad_medida"`
	Temporada        string `json:"temporada"`
	Estado           bool   `json:"estado"`
}
