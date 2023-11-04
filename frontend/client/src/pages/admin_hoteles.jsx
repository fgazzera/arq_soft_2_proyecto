import React, { useContext } from 'react';
import { AuthContext } from './login/auth';
import { Link } from 'react-router-dom';
import './estilo/admin_hoteles.css';

const handleVolver = () => {
  window.history.back();
};

const AdminHotelesPage = () => {
  const { isLoggedAdmin } = useContext(AuthContext);
  
  const Verificacion = () => {
    if (!isLoggedAdmin) {
      window.location.href = '/login-admin';
    }
  };

  return (
    <div className="containerHotel" onLoad={Verificacion}>
      <div className="rectangulo">
        <h1 className="titulo">Hoteles🏨</h1>
        <div className="botones-container">
          <Link to="/agregar-hoteles" className="botonAH">
            Agregar Hoteles
          </Link>
          <Link to="/ver-hoteles" className="botonAH">
            Ver Hoteles
          </Link>
          <Link to="/editar-hoteles" className="botonAH">
            Editar Hoteles
          </Link>
          <Link to="/agregar-imagenes" className="botonAH">
            Agregar Imágenes
          </Link>
        </div>
        <button className="botonBack" onClick={handleVolver}>
        🔙
      </button>
      </div>
    </div>
  );
};

export default AdminHotelesPage;