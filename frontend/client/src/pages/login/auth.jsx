import React, { createContext, useState, useEffect, useCallback } from 'react';

export const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
  const [isLoggedCliente, setIsLoggedCliente] = useState(localStorage.getItem('auth') === 'true');
  const [isLoggedAdmin, setIsLoggedAdmin] = useState(localStorage.getItem('auth') === 'true');

  const loginAdmin = (newToken, id) => {
    setIsLoggedAdmin(true);
    localStorage.setItem('token', newToken);
    localStorage.setItem('id_admin', id);
    localStorage.setItem('auth', true);
  };

  const loginCliente = (newToken, id) => {
    setIsLoggedCliente(true);
    localStorage.setItem('token', newToken);
    localStorage.setItem('id_cliente', id);
    localStorage.setItem('auth', true);
  };

  const logout = useCallback(async () => {
    setIsLoggedCliente(false);
    setIsLoggedAdmin(false);
    localStorage.removeItem('token');
    localStorage.removeItem('id_admin');
    localStorage.removeItem('id_cliente');
    localStorage.setItem('auth', false);
  }, []);

  const comprobarLogin = useCallback(async () => {
    if (isLoggedCliente) {
      const accountId = localStorage.getItem("id_cliente");
      try {
        const request = await fetch(`http://localhost:8090/cliente/${accountId}`);
        const response = await request.json();
        if (!response) {
          logout();
        }
      } catch (error) {
        console.log("No se pudieron obtener los datos del cliente:", error);
        logout();
      }
    }
    
    if (isLoggedAdmin) {
      const accountId = localStorage.getItem("id_admin");
      try {
        const request = await fetch(`http://localhost:8090/admin/${accountId}`);
        const response = await request.json();
        if (!response) {
          logout();
        }
      } catch (error) {
        console.log("No se pudieron obtener los datos del admin:", error);
        logout();
      }
    }
  }, [logout, isLoggedAdmin, isLoggedCliente]);

  useEffect(() => {
    comprobarLogin();
  }, [comprobarLogin]);

  const propiedades = {
    isLoggedCliente,
    isLoggedAdmin,
    loginCliente,
    loginAdmin,
    logout,
  };

  return (
    <AuthContext.Provider value={propiedades}>
      {children}
    </AuthContext.Provider>
  );
};
