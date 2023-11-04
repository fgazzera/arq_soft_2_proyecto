import React, { useContext, useEffect, useState } from 'react';
import { AuthContext } from './login/auth';
import './estilo/ver_clientes.css';

const handleVolver = () => {
  window.history.back();
};

const VerClientes = () => {
  const [clientes, setClientes] = useState([]);
  const { isLoggedAdmin } = useContext(AuthContext);

  const getClientes = async () => {
    try {
      const request = await fetch("http://localhost:8090/admin/clientes");
      const response = await request.json();
      setClientes(response);
    } catch (error) {
      console.log("No se pudieron obtener los hoteles:", error);
    }
  };

  useEffect(() => {
    getClientes();
  }, []);

  const Verificacion = () => {
    if (!isLoggedAdmin) {
      window.location.href = '/login-admin';
    }
  };

  return (
    <body className="bodyinicio" onLoad={Verificacion}>
      <div className="containerIni">
        <div className="hotels-container">
          {clientes.length ? (
            clientes.map((cliente) => (
              <div className="hotel-card" key={cliente.id}>
                <div className="hotel-info">
                  <h4>{cliente.name}</h4>
                  <p>{cliente.last_name}</p>
                </div>
                <div className="hotel-info">
                  <p>{cliente.username}</p>
                  <p>{cliente.email}</p>
                </div>
              </div>
            ))
          ) : (
            <p>No hay clientes</p>
          )}
        </div>
      </div>
      <button className="botonBack" onClick={handleVolver}>
        🔙
      </button>
    </body>
  );
};

export default VerClientes;