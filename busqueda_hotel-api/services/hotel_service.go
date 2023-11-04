package services

import (
	"busqueda_hotel-api/dtos"
	"busqueda_hotel-api/go_solr"
	e "busqueda_hotel-api/utils/errors"

	"github.com/vanng822/go-solr/solr"
)

type hotelService struct{}

type hotelServiceInterface interface {
	GetHotelById(id string) (dtos.HotelDto, e.ApiError)
	InsertHotel(hotelDto dtos.HotelDto) (e.ApiError)
	//UpdateHotelById(id string, hotelDto dtos.HotelDto) (dtos.HotelDto, e.ApiError)
}

var (
	HotelService hotelServiceInterface
)

func init() {
	HotelService = &hotelService{}
}

func (s *hotelService) GetHotelById(id string) (dtos.HotelDto, e.ApiError) {
	solrClient := go_solr.GetSolrClient()

	// Define la consulta para buscar el hotel por ID
	query := solr.NewQuery()
	query.Q("id:" + id)

	// Realiza la consulta a Solr
	//response := solrClient.Search(query)
	search := solrClient.Search(query)
	response, _ := search.Result(nil)

	/*if err != nil {
		// Maneja el error de la consulta a Solr
		return dtos.HotelDto{}, e.NewInternalServerApiError("error in Solr query", err)
	}*/

	// Comprueba si se encontraron resultados
	if response.Results.NumFound == 0 {
		return dtos.HotelDto{}, e.NewBadRequestApiError("hotel not found")
	}

	// Obtén el primer resultado (suponiendo que solo haya una coincidencia)
	hotel := response.Results.Docs[0]

	// Crea un HotelDto a partir de los datos de Solr
	hotelDto := dtos.HotelDto{
		ID:          hotel["id"].(string),
		Nombre:      hotel["nombre"].(string),
		Descripcion: hotel["descripcion"].(string),
		Email:       hotel["email"].(string),
		Ciudad:      hotel["ciudad"].(string),
		Images:      hotel["images"].([]string),
		CantHab:     int(hotel["canthab"].(float64)),
		Amenities:   hotel["amenities"].([]string),
	}

	return hotelDto, nil
}

func (s *hotelService) InsertHotel(hotelDto dtos.HotelDto) (e.ApiError) {

	solrClient := go_solr.GetSolrClient()

	// Establece los campos del documento con los datos del hotel
	hotelDocument := make(solr.Document)
	hotelDocument["id"] = hotelDto.ID
	hotelDocument["nombre"] = hotelDto.Nombre
	hotelDocument["descripcion"] = hotelDto.Descripcion
	hotelDocument["email"] = hotelDto.Email
	hotelDocument["ciudad"] = hotelDto.Ciudad
	hotelDocument["images"] = hotelDto.Images
	hotelDocument["canthab"] = hotelDto.CantHab
	hotelDocument["amenities"] = hotelDto.Amenities

	documents := []solr.Document{hotelDocument}
	_, err := solrClient.Add(documents, 100, nil)

	if err != nil {
		return e.NewInternalServerApiError("Error inserting hotel into Solr", err)
	}

	// Envía los cambios a Solr
	_, err = solrClient.Commit()

	if err != nil {
		return e.NewInternalServerApiError("Error committing changes to Solr", err)
	}

	return nil
}

/*func InsertHotelIntoSolr(hotel dtos.HotelDto) error {
    solrClient := go_solr.GetSolrClient()

    // Crear un nuevo documento para el hotel a ser insertado
    doc := solr.NewAddDocument()
    doc.Set("id", hotel.ID)
    doc.Set("nombre", hotel.Nombre)
    doc.Set("descripcion", hotel.Descripcion)
    doc.Set("email", hotel.Email)
    doc.Set("ciudad", hotel.Ciudad)
    doc.Set("images", hotel.Images)
    doc.Set("canthab", hotel.CantHab)
    doc.Set("amenities", hotel.Amenities)

    // Enviar el documento al cliente Solr para ser insertado
    _, err := solrClient.Update(doc)

    if err != nil {
        return err
    }

    // Si la inserción fue exitosa, realiza un commit para que los cambios se apliquen inmediatamente
    _, err = solrClient.Commit()

    return err
}*/

/*func (s *hotelService) UpdateHotelById(id string, hotelDto dtos.HotelDto) (dtos.HotelDto, e.ApiError) {

	var hotel model.Hotel = hotelDao.GetHotelById(id)

	if hotel.ID.Hex() == "000000000000000000000000" {
		return hotelDto, e.NewBadRequestApiError("hotel not found")
	}

	hotel.Nombre = hotelDto.Nombre
	hotel.Descripcion = hotelDto.Descripcion
	hotel.Email = hotelDto.Email
	hotel.Ciudad = hotelDto.Ciudad
	hotel.Images = hotelDto.Images
	hotel.CantHab = hotelDto.CantHab
	hotel.Amenities = hotelDto.Amenities

	hotel = hotelDao.UpdateHotel(hotel)
	hotelDto.ID = hotel.ID.Hex()

	return hotelDto, nil
}*/
