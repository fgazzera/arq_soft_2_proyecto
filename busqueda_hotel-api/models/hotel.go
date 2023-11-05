package model

type Hotel struct {
	ID          string   `json:"id"`
	Nombre      string   `json:"nombre"`
	Descripcion string   `json:"descripcion"`
	Email       string   `json:"email"`
	Ciudad      string   `json:"ciudad"`
	Images      []string `json:"images"`
	CantHab     int      `json:"cant_hab"`
	Amenities   []string `json:"amenities"`
}

type Hotels []Hotel
