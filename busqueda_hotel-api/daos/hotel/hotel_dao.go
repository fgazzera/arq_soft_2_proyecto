package hotel

import (
	model "busqueda_hotel-api/models"
	"busqueda_hotel-api/utils/db"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	solr "github.com/rtt/Go-Solr"
)

type HotelDao interface {
	Get(id string) (*model.Hotel, error)
	Create(hotel *model.Hotel) error
	Update(hotel *model.Hotel) error
	GetAll() ([]*model.Hotel, error)
	GetByCiudad(city string) ([]*model.Hotel, error)
}

type HotelSolrDao struct{}

func NewHotelSolrDAO() HotelDao {
	return &HotelSolrDao{}
}

func (dao *HotelSolrDao) Create(hotel *model.Hotel) error {
	// Crear un mapa que representa el documento del hotel
	hotelDocument := map[string]interface{}{
		"add": []interface{}{
			map[string]interface{}{
				"id":          hotel.ID,
				"nombre":      hotel.Nombre,
				"descripcion": hotel.Descripcion,
				"email":       hotel.Email,
				"ciudad":      hotel.Ciudad,
				"images":      hotel.Images,
				"cant_hab":    float64(hotel.CantHab),
				"amenities":   hotel.Amenities,
			},
		},
	}

	// Inserta el nuevo documento en Solr
	_, err := db.SolrClient.Update(hotelDocument, true) // El segundo parámetro "true" realiza una confirmación inmediata
	if err != nil {
		return err
	}
	return nil
}

func (dao *HotelSolrDao) Update(hotel *model.Hotel) error {
	// Crear un mapa que representa los cambios en el documento del hotel
	hotelDocument := map[string]interface{}{
		"add": []interface{}{
			map[string]interface{}{
				"id":          hotel.ID,
				"nombre":      hotel.Nombre,
				"descripcion": hotel.Descripcion,
				"email":       hotel.Email,
				"ciudad":      hotel.Ciudad,
				"images":      hotel.Images,
				"cant_hab":    float64(hotel.CantHab),
				"amenities":   hotel.Amenities,
			},
		},
	}

	// Actualiza el documento del hotel en Solr utilizando la nueva API
	updateURL := "http://localhost:8983/solr/busqueda_hotel-core/update?commit=true" // Reemplaza "your-core" con el nombre de tu core en Solr

	requestBody, err := json.Marshal(hotelDocument)
	if err != nil {
		return err
	}

	resp, err := http.Post(updateURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Error al actualizar el hotel en Solr. Código de respuesta: %d", resp.StatusCode)
	}

	return nil
}

func (dao *HotelSolrDao) Get(id string) (*model.Hotel, error) {
	// Define la consulta para buscar un hotel por su ID
	query := &solr.Query{
		Params: solr.URLParamMap{
			"q":    []string{fmt.Sprintf("id:%s", id)}, // Consulta para buscar por ID
			"rows": []string{"1"},                      // Obtener solo un resultado
		},
	}

	// Realiza la consulta a Solr
	resp, err := db.SolrClient.Select(query)
	if err != nil {
		println("EL ERROR ES: ", err.Error())
		return nil, err
	}

	// Verifica si se encontró algún resultado
	if len(resp.Results.Collection) == 0 {
		// No se encontraron hoteles con el ID especificado
		return nil, fmt.Errorf("hotel not found")
	}

	// Extrae el primer resultado (debería haber solo uno)
	doc := resp.Results.Collection[0]

	println(doc.Field("images"))
	println(doc.Field("amenities"))

	// Construye un modelo de hotel a partir de los campos del documento
	hotel := &model.Hotel{
		ID:          doc.Fields["id"].(string),
		Nombre:      doc.Field("nombre").([]interface{})[0].(string),
		Descripcion: doc.Field("descripcion").([]interface{})[0].(string),
		Email:       doc.Field("email").([]interface{})[0].(string),
		Ciudad:      doc.Field("ciudad").([]interface{})[0].(string),
		Images:      getStringSliceFromInterface(doc.Field("images")),//getStringSliceFromInterface(doc.Field("images")),
		CantHab:     int(doc.Field("cant_hab").([]interface{})[0].(float64)),
		Amenities:   getStringSliceFromInterface(doc.Field("amenities")),//getStringSliceFromInterface(doc.Field("amenities")), 
	}

	return hotel, nil
}

func (dao *HotelSolrDao) GetAll() ([]*model.Hotel, error) {
	// Define la consulta para obtener todos los hoteles
	query := &solr.Query{
		Params: solr.URLParamMap{
			"q":    []string{"*:*"},  // Consulta que selecciona todos los documentos
			"rows": []string{"1000"}, // Número máximo de filas a recuperar (ajusta según tus necesidades)
		},
	}

	// Realiza la consulta a Solr
	resp, err := db.SolrClient.Select(query)
	if err != nil {
		return nil, err
	}

	// Itera a través de los resultados y construye una lista de hoteles
	var hotels []*model.Hotel
	for _, doc := range resp.Results.Collection {
		hotel := &model.Hotel{
			ID:          doc.Fields["id"].(string),
			Nombre:      doc.Field("nombre").([]interface{})[0].(string),
			Descripcion: doc.Field("descripcion").([]interface{})[0].(string),
			Email:       doc.Field("email").([]interface{})[0].(string),
			Ciudad:      doc.Field("ciudad").([]interface{})[0].(string),
			Images:      getStringSliceFromInterface(doc.Field("images")),
			CantHab:     int(doc.Field("cant_hab").([]interface{})[0].(float64)),
			Amenities:   getStringSliceFromInterface(doc.Field("amenities")),
		}
		hotels = append(hotels, hotel)
	}

	return hotels, nil
}

func (dao *HotelSolrDao) GetByCiudad(ciudad string) ([]*model.Hotel, error) {
	// Define la consulta para obtener hoteles por ciudad
	query := &solr.Query{
		Params: solr.URLParamMap{
			"q":    []string{fmt.Sprintf("ciudad:\"%s\"", ciudad)}, // Consulta con filtro por ciudad
			"rows": []string{"1000"},                               // Número máximo de filas a recuperar (ajusta según tus necesidades)
		},
	}

	// Realiza la consulta a Solr
	resp, err := db.SolrClient.Select(query)
	if err != nil {
		return nil, err
	}

	// Itera a través de los resultados y construye una lista de hoteles
	var hotels []*model.Hotel
	for _, doc := range resp.Results.Collection {
		hotel := &model.Hotel{
			ID:          doc.Fields["id"].(string),
			Nombre:      doc.Field("nombre").([]interface{})[0].(string),
			Descripcion: doc.Field("descripcion").([]interface{})[0].(string),
			Email:       doc.Field("email").([]interface{})[0].(string),
			Ciudad:      doc.Field("ciudad").([]interface{})[0].(string),
			Images:      getStringSliceFromInterface(doc.Field("images")),
			CantHab:     int(doc.Field("cant_hab").([]interface{})[0].(float64)),
			Amenities:   getStringSliceFromInterface(doc.Field("amenities")),
		}
		hotels = append(hotels, hotel)
	}

	return hotels, nil
}

func getStringSliceFromInterface(i interface{}) []string {
	result := []string{}
	if i == nil {
		return result
	}
	if slice, ok := i.([]interface{}); ok {
		result := make([]string, len(slice))
		for i, v := range slice {
			if str, ok := v.(string); ok {
				result[i] = str
			}
		}
		return result
	}
	return result
}
