package services

import (
	"busqueda_hotel-api/dtos"
	"busqueda_hotel-api/go_solr"
	e "busqueda_hotel-api/utils/errors"
	"errors"

	"github.com/vanng822/go-solr/solr"
)

type hotelService struct{}

type hotelServiceInterface interface {
	GetHotelById(id string) (dtos.HotelDto, e.ApiError)
	InsertHotel(hotelDto dtos.HotelDto) e.ApiError
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
	query := solr.NewQuery()
	query.Q("id:" + id)
	search := solrClient.Search(query)
	response, err := search.Result(nil)

	if err != nil {
		// Maneja el error de la consulta a Solr
		return dtos.HotelDto{}, e.NewInternalServerApiError("error in Solr query", err)
	}

	if response.Results.NumFound == 0 {
		return dtos.HotelDto{}, e.NewBadRequestApiError("hotel not found")
	}

	hotel := response.Results.Docs[0]
	imagesInterface, ok := hotel["images"].([]interface{})

	if !ok {
		return dtos.HotelDto{}, e.NewInternalServerApiError("error in Solr data type conversion", errors.New("type conversion error"))
	}

	var images []string

	for _, image := range imagesInterface {
		if str, ok := image.(string); ok {
			images = append(images, str)
		} else {
			return dtos.HotelDto{}, e.NewInternalServerApiError("error in Solr data type conversion", errors.New("type conversion error"))
		}
	}

	amenitiesInterface, ok := hotel["amenities"].([]interface{})
	if !ok {
		return dtos.HotelDto{}, e.NewInternalServerApiError("error in Solr data type conversion", errors.New("type conversion error"))
	}

	var amenities []string

	for _, amenity := range amenitiesInterface {
		if str, ok := amenity.(string); ok {
			amenities = append(amenities, str)
		} else {
			return dtos.HotelDto{}, e.NewInternalServerApiError("error in Solr data type conversion", errors.New("type conversion error"))
		}
	}

	hotelDto := dtos.HotelDto{
		ID:          hotel["id"].(string),
		Nombre:      hotel["nombre"].(string),
		Descripcion: hotel["descripcion"].(string),
		Email:       hotel["email"].(string),
		Ciudad:      hotel["ciudad"].(string),
		Images:      images,
		CantHab:     int(hotel["cant_hab"].(float64)),
		Amenities:   amenities,
	}

	return hotelDto, nil
}

func (s *hotelService) InsertHotel(hotelDto dtos.HotelDto) e.ApiError {

	solrClient := go_solr.GetSolrClient()

	/*hotelDocument := make(solr.Document)

	hotelDocument["id"] = hotelDto.ID
	hotelDocument["nombre"] = hotelDto.Nombre
	hotelDocument["descripcion"] = hotelDto.Descripcion
	hotelDocument["email"] = hotelDto.Email
	hotelDocument["ciudad"] = hotelDto.Ciudad
	hotelDocument["images"] = hotelDto.Images
	hotelDocument["cant_hab"] = hotelDto.CantHab
	hotelDocument["amenities"] = hotelDto.Amenities*/

	hotelDocument := map[string]interface{}{
		"add": []interface{}{
			map[string]interface{}{
				"id":          hotelDto.ID,
				"nombre":      hotelDto.Nombre,
				"descripcion": hotelDto.Descripcion,
				"email":       hotelDto.Email,
				"ciudad":      hotelDto.Ciudad,
				"images":      hotelDto.Images,
				"cant_hab":    hotelDto.CantHab,
				"amenities":   hotelDto.Amenities,
			},
		},
	}

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

/*hotelDocument := map[string]interface{}{
	"add": []interface{}{
		map[string]interface{}{
			"id":          hotel.ID,
			"name":        hotel.Name,
			"city":        hotel.City,
			"description": hotel.Description,
			//"thumbnail":   hotel.Thumbnail,
			"amenities": hotel.Amenities,
			"images":    hotel.Images,
		},
	},
}

// Inserta el nuevo documento en Solr
_, err := db.SolrClient.Update(hotelDocument, true) // El segundo parámetro "true" realiza una confirmación inmediata
if err != nil {
	return err
}
return nil*/

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
