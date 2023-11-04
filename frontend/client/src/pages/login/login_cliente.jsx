import React, { useContext, useState } from 'react'
import { AuthContext } from './auth';
import { Link } from 'react-router-dom';
import '../estilo/login_cliente.css';

const ClienteLogin = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const { loginCliente } = useContext(AuthContext);

  const handleLoginCliente = () => {
    fetch(`http://localhost:8090/cliente/email/${email}`)
    .then(response => response.json())
    .then(data => {
      if (email === data.cliente.email && password === data.cliente.password) {
        loginCliente(data.token, data.cliente.id);
        window.location.href = '/';
      } else {
        alert('Credenciales incorrectas');
      }
    })
    .catch(error => {
      console.error('Error al obtener los datos del cliente:', error);
    });
  };

  const handleVolver = () => {
    window.location.href = 'http://localhost:3000/';
  };

  return (
 <body className="bodylogclient">
    <div className="contLogClie1">
    <div className="contLogClien2">
      <h1 className="title">Bienvenido Cliente</h1>
       <div className="form-container">
        <input
          type="text"
          placeholder="Correo electrónico"
          className="inputLcli"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
        <input
          type="password"
          placeholder="Contraseña"
          className="inputLcli"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <div className="button-container">
          <button className="buttonClient" onClick={handleLoginCliente}>
            Iniciar Sesión
          </button>
          <Link to="/register" className="buttonClient">
            Registrarse
          </Link>
          <div className="home-container">
            <button className="home" onClick={handleVolver}>
              Home
            </button>
          </div>
        </div>
      </div>
    </div>
    </div>
  </body>
  );
};

export default ClienteLogin;
