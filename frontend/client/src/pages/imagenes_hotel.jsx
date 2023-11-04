import React, { useContext, useEffect, useState } from 'react';
import { AuthContext } from './login/auth';
import { useParams } from 'react-router-dom';
import './estilo/imagenes_hotel.css';

const handleVolver = () => {
  window.history.back();
};

function ImagenesHotel() {
  const { hotelId } = useParams();
  const [hotelData, setHotelData] = useState('');
  const { isLoggedAdmin } = useContext(AuthContext);
  const [imagenesHotel, setImagenesHotel] = useState([]);
  const [imagenes, setImagenes] = useState([]);
  const [imagenesSeleccionadas, setImagenesSeleccionadas] = useState([]);

  const Verificacion = () => {
    if (!isLoggedAdmin) {
      window.location.href = '/login-admin';
    }
  };

  useEffect(() => {
    if (hotelId) {
      fetch(`http://localhost:8090/admin/hotel/${hotelId}`)
        .then(response => response.json())
        .then(data => {
          setHotelData(data);
        })
        .catch(error => {
          console.error('Error al obtener los datos del hotel:', error);
        });
    }
  }, [hotelId]);

  useEffect(() => {
    if (hotelId) {
      fetch(`http://localhost:8090/admin/imagenes/hotel/${hotelId}`)
        .then(response => response.json())
        .then(data => {
            if (data.length > 1) {
                let imagenesArray = []
                for (let i = 1; i < data.length; i++) {
                    imagenesArray.push(data[i])
                }
                setImagenesHotel(imagenesArray);
            }
        })
        .catch(error => {
          console.error('Error al obtener los datos del hotel:', error);
        });
    }
  }, [hotelId]);

  const handleChange = event => {
    const files = event.target.files;
    const imagenesArray = [];

    for (let i = 0; i < files.length; i++) {
        imagenesArray.push(files[i]);
    }

    setImagenes(imagenesArray);
  };

  const AgregarImagenes = async () => {
    if (imagenesSeleccionadas.length > 0) {
        try {
            for (let i = 0; i < imagenesSeleccionadas.length; i++) {
                const imagen = imagenesSeleccionadas[i];
                fetch(`http://localhost:8090/admin/imagen/delete/${imagen.id}`,
                {
                    method: 'DELETE',
                });
            }
            setImagenesSeleccionadas([]);
        } catch (error) {
            console.log("No se pudieron borrar las imagenes:", error);
        }
    }

    if (imagenes.length > 0) {
      const formDataWithImages = new FormData();
      for (let i = 0; i < imagenes.length; i++) {
        formDataWithImages.append('imagen', imagenes[i]);
      }
      
      const req = await fetch(
        `http://localhost:8090/admin/hotel/${hotelId}/add-imagenes`,
        {
          method: 'POST',
          body: formDataWithImages
        }
      );
      const res = await req.json();

      if (res) {
        window.location.href = '/agregar-imagenes';
      } else {
        console.error('Error en el registro:', res);
        alert('Imagen no registrada');
      }
    } else {
      window.location.href = '/agregar-imagenes';
    }
  };

  return (
    <div className="registration-container" onLoad={Verificacion}>
      <h2 className="nombre">{hotelData['nombre']}</h2>
      <form onSubmit={AgregarImagenes} className="registration-form">
        <label>
          Imagen:
          <input type="file" name="imagen" onChange={handleChange} multiple />
        </label>
        <div className="image-grid">
          {imagenesHotel.map((imagen) => (
            <img key={imagen.id} src={`http://localhost:8090/${imagen.url}`} alt={`Imagen ${imagen.id}`} 
            onClick={() => {
                const isSelected = imagenesSeleccionadas.includes(imagen);
                if (isSelected) {
                  const updatedSelection = imagenesSeleccionadas.filter(
                    img => img !== imagen
                  );
                  setImagenesSeleccionadas(updatedSelection);
                } else {
                  setImagenesSeleccionadas([...imagenesSeleccionadas, imagen]);
                }
              }}
              className={imagenesSeleccionadas.includes(imagen) ? 'selected' : ''} />
          ))}
        </div>
        <button type="submit">Guardar Imágenes</button>
      </form>
      <button className="botonBack" onClick={handleVolver}>
        🔙
      </button>
    </div>
  );
}

export default ImagenesHotel;
