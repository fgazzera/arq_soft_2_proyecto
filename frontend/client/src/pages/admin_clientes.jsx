import React, { useContext } from 'react';
import { AuthContext } from './login/auth';
import { Link } from 'react-router-dom';
import './estilo/admin_clientes.css';

const handleVolver = () => {
  window.history.back();
};

const AdminClientesPage = () => {
  const { isLoggedAdmin } = useContext(AuthContext);
  
  const Verificacion = () => {
    if (!isLoggedAdmin) {
      window.location.href = '/login-admin';
    }
  };

  return (
    <div className="container" onLoad={Verificacion}>
      <div className= "rectangulo1">
      <h1 className="titulo">Clientes👥</h1>
      <div className="botones-container">
      <Link to="/ver-reservas" className="botonAC">
          Ver Reservas
        </Link>
        <Link to="/ver-clientes" className="botonAC">
          Ver Clientes
        </Link>
        </div>
        <button className="botonBack" onClick={handleVolver}>
        🔙
      </button>
      </div>
    </div>
  );
};

export default AdminClientesPage;
