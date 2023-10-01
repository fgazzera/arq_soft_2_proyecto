package dtos

type HotelDto struct {
	ID          string   `json:"id"`
	Nombre      string   `json:"nombre"`
	Description string   `json:"description"`
	Email       string   `json:"email"`
	Ciudad      string   `json:"ciudad"`
	Images      []string `json:"images"`
	CantHab     int      `json:"cant_hab"`
	Amenities   []string `json:"amenities"`
}
