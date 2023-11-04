import React, { useContext, useEffect, useState, useCallback } from 'react';
import { AuthContext } from './login/auth';
import './estilo/editar_hoteles.css';

const handleVolver = () => {
  window.history.back();
};

const EditarHoteles = () => {
  const [hotels, setHotels] = useState([]);
  const [imagenes, setImagenes] = useState([]);
  const { isLoggedAdmin } = useContext(AuthContext);

  const getHotels = async () => {
    try {
      const request = await fetch("http://localhost:8090/admin/hoteles");
      const response = await request.json();
      setHotels(response);
    } catch (error) {
      console.log("No se pudieron obtener los hoteles:", error);
    }
  };

  const getImagenes = useCallback(async () => {
    try {
      const imagenesArray = [];
      for (let i = 0; i < hotels.length; i++) {
        const hotel = hotels[i];
        const request = await fetch(`http://localhost:8090/admin/imagenes/hotel/${hotel.id}`);
        const response = await request.json();
        if (response.length > 0) {
          imagenesArray.push({url: response[0].url, hotel_id: response[0].hotel_id});
        }
      }
      setImagenes(imagenesArray);
    } catch (error) {
      console.log("No se pudieron obtener las imagenes de loshoteles:", error);
    }
  }, [hotels]);

  useEffect(() => {
    getHotels();
  }, []);

  useEffect(() => {
    getImagenes();
  }, [getImagenes]);

  const Verificacion = () => {
    if (!isLoggedAdmin) {
      window.location.href = '/login-admin';
    }
  };

  const handleEditHotel = (hotelId) => {
    window.location.href = `/editar-hotel/${hotelId}`;
  };

  return (
    <body className="bodyinicio" onLoad={Verificacion}>
      <div className="containerIni">
        <div className="hotels-container">
          {hotels.length ? (
            hotels.map((hotel) => {
              const imagen = imagenes.find((imagen) => imagen.hotel_id === hotel.id);
              return (
              <div className="hotel-card" key={hotel.id}>
                {imagen ? (
                  <img src={`http://localhost:8090/${imagen.url}`} alt={hotel.nombre} className="hotel-image" />
                ) : (
                  <div className="hotel-image-placeholder" />
                )}
                <div className="hotel-info">
                  <h4>{hotel.nombre}</h4>
                  <p>{hotel.email}</p>
                </div>
                <div className="hotel-description">
                    <label htmlFor={`description-${hotel.id}`}>Descripcion:</label>
                    <p id={`description-${hotel.id}`}>{hotel.descripcion}</p>
                </div>
                <button className="edit-button" onClick={() => handleEditHotel(hotel.id)}>
                    Editar
                </button>
              </div>
            )})
          ) : (
            <p>No hay hoteles</p>
          )}
        </div>
      </div>
      <button className="botonBack" onClick={handleVolver}>
        🔙
      </button>
    </body>
  );
};

export default EditarHoteles;