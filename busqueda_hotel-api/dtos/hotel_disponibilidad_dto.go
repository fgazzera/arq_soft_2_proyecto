package dtos

type HotelDisponibilidadDto struct {
	ID             string   `json:"id"`
	Nombre         string   `json:"nombre"`
	Descripcion    string   `json:"descripcion"`
	Email          string   `json:"email"`
	Ciudad         string   `json:"ciudad"`
	Images         []string `json:"images"`
	CantHab        int      `json:"cant_hab"`
	Amenities      []string `json:"amenities"`
	Disponibilidad bool     `json:"disponibilidad"`
}

type HotelsDisponibilidadDto struct {
	Hotels []HotelDisponibilidadDto `json:"hotels"`
}
