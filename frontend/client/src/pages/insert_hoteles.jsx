import React, { useContext, useEffect, useState } from 'react';
import { AuthContext } from './login/auth';
import './estilo/insert_hoteles.css'

const handleVolver = () => {
  window.history.back();
};

function InsertHotel() {
  const [Email, setEmail] = useState({});
  const [Nombre, setNombre] = useState({});
  const { isLoggedAdmin } = useContext(AuthContext);
  const [imagen, setImagen] = useState('');
  
  const Verificacion = () => {
    if (!isLoggedAdmin) {
      window.location.href = '/login-admin';
    }
  };

  const [formData, setFormData] = useState({
    nombre: '',
    descripcion: '',
    email: '',
    cant_hab: '',
    amenities: ''
  });

  const handleChange = (event) => {
    const { name, value, files } = event.target;
    if (name === "imagen") {
      setImagen(files[0]);
      console.log(imagen)
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

  const RegisterHotel = async () => {
    if (formData.email === Email.email) {
      alert('El email ya pertenece a un hotel');
    }
    else if (formData.nombre === Nombre.nombre) {
      alert('El nombre no esta disponible');
    }
    else
    {
      try {
        const request = await fetch('http://localhost:8090/admin/hotel', {
        method: 'POST',
        headers: {
        'Content-Type': 'application/json'
        },
        body: JSON.stringify(formData)
        })
        
        const response = await request.json();
        const hotelId = response.id;

        if (response) {
          insertImagen(hotelId);
        }
      } catch (error) {
        console.error('Error en el registro:', error);
        alert('Error en el registro del hotel');
      }
    }
  };

  const insertImagen = async (hotelId) => {
    const formDataWithImagen = new FormData();
    formDataWithImagen.append("imagen", imagen);
    
    const req = await fetch(`http://localhost:8090/admin/hotel/${hotelId}/add-imagen`, {
      method: 'POST',
      body: formDataWithImagen
    });

    const res = await req.json()
    if (res) {
      window.location.href = '/ver-hoteles';
    } else {
      console.error('Error en el registro:', res);
      alert('Error en el registro de la imagen');
    }
  };

  return (
    <div className="registration-container" onLoad={Verificacion}>
      <h2>Registro De Hoteles</h2>
      <form onSubmit={RegisterHotel} className="registration-form" encType="multipart/form-data">
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
            required
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
        <button type="submit">Registrar Hotel</button>
      </form>
      <button className="botonBack" onClick={handleVolver}>
        🔙
      </button>
    </div>
  );
}

export default InsertHotel;