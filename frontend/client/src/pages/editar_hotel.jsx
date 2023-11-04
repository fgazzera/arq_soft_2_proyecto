import React, { useContext, useEffect, useState } from 'react';
import { AuthContext } from './login/auth';
import { useParams } from 'react-router-dom';
import './estilo/editar_hotel.css'

const handleVolver = () => {
  window.history.back();
};

function EditarHotel() {
  const { hotelId } = useParams();
  const [hotelData, setHotelData] = useState('');
  const [Email, setEmail] = useState({});
  const [Nombre, setNombre] = useState({});
  const { isLoggedAdmin } = useContext(AuthContext);
  const [imagen, setImagen] = useState('');
  const [formData, setFormData] = useState({
    nombre: '',
    descripcion: '',
    email: '',
    cant_hab: '',
    amenities: ''
  });
  
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
          setFormData({nombre: data.nombre, descripcion: data.descripcion, email: data.email, cant_hab: data.cant_hab, amenities: data.amenities});
          setHotelData(data);
        })
        .catch(error => {
          console.error('Error al obtener los datos del hotel:', error);
        });
    }
  }, [hotelId]);

  const handleChange = (event) => {
    const { name, value, files } = event.target;

    if (name === "imagen") {
      setImagen(files[0]);
    } else if (name === "cant_hab" && value !== "") {
      const intValue = parseInt(value);
      setFormData((prevFormData) => ({
        ...prevFormData,
        [name]: intValue,
      }));
    } else {
      setFormData((prevFormData) => ({
        ...prevFormData,
        [name]: value,
      }));
    }
  };

  useEffect(() => {
    setEmail('');
  
    if (formData.email) {
      fetch(`http://localhost:8090/admin/hotel/email/${formData.email}`)
        .then(response => response.json())
        .then(data => {
            setEmail(data);
        })
        .catch(error => {
          console.error('Error al obtener los datos del cliente:', error);
        });
    }
  }, [formData.email]);

  useEffect(() => {
    setNombre('');
  
    if (formData.nombre) {
      fetch(`http://localhost:8090/admin/hotel/nombre/${formData.nombre}`)
        .then(response => response.json())
        .then(data => {
            setNombre(data);
        })
        .catch(error => {
          console.error('Error al obtener los datos del cliente:', error);
        });
    }
  }, [formData.nombre]);

  const ActualizarHotel = async () => {
    if (formData.email !== hotelData.email && formData.email === Email.email) {
      alert('El email ya pertenece a un hotel');
    }
    else if (formData.nombre !== hotelData.nombre && formData.nombre === Nombre.nombre) {
      alert('El nombre no esta disponible');
    }
    else
    {
      const request = await fetch(`http://localhost:8090/admin/hotel/${hotelId}`, {
      method: 'PUT',
      headers: {
      'Content-Type': 'application/json'
      },
      body: JSON.stringify(formData)
      })

      const response = await request.json()

      if (response) {
        if (imagen !== '') {
          alert("Hola")
          const formDataWithImage = new FormData();
          formDataWithImage.append("imagen", imagen);
          console.log(formDataWithImage)

          const req = await fetch(`http://localhost:8090/admin/hotel/imagen/${response.id}`, { 
            method: 'PUT',
            body: formDataWithImage
          })
          const res = await req.json();

          if (res) {
            window.location.href = '/editar-hoteles';
          }
          else {
            console.error('Error en el registro:', res);
            alert('Imagen no registrada');
          }
        }
        else {
          window.location.href = '/editar-hoteles';
        }
      }
      else {
        console.error('Error en el registro:', response);
        alert('Hotel no registrado');
      }
    }
  };

  return (
    <div className="registration-container" onLoad={Verificacion}>
      <h2 className="nombre">{hotelData["nombre"]}</h2>
      <form onSubmit={ActualizarHotel} className="registration-form">
        <label>
          Nombre:
          <input
            type="text"
            name="nombre"
            value={formData.nombre}
            onChange={handleChange}
            required
          />
        </label>
        <br />
        <label>
         Descripcion:
          <input
            type="text"
            name="descripcion"
            value={formData.descripcion}
            onChange={handleChange}
            required
          />
        </label>
        <br />
        <label>
          Email:
          <input
            type="text"
            name="email"
            value={formData.email}
            onChange={handleChange}
            required
          />
        </label>
        <br />
        <label>
         Imagen:
          <input
            type="file"
            name="imagen"
            onChange={handleChange}
          />
        </label>
        <br />
        <label>
          Cant_hab:
          <input
            type="text"
            name="cant_hab"
            value={formData.cant_hab}
            onChange={handleChange}
            required
          />
        </label>
        <br />
        <label>
         Amenities:
          <input
            type="text"
            name="amenities"
            value={formData.amenities}
            onChange={handleChange}
            placeholder="Ingrese las Amenities"
          />
        </label>
        <br />
        <button type="submit">Guardar Hotel</button>
      </form>
      <button className="botonBack" onClick={handleVolver}>
        🔙
      </button>
    </div>
  );
}

export default EditarHotel;