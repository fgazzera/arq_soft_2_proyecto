import React from 'react';
import './Congrats.css'; // Asegúrate de que el nombre del archivo CSS coincida

const Congrats = ({ isSuccessful }) => {
  return (
    <div className="congrats">
      <h1>Confirmacion {isSuccessful ? 'Exitosa' : 'Rechazada'}</h1>
      <p>
        {isSuccessful
          ? 'Gracias por tu reserva exitosa!'
          : 'Lo sentimos, la reserva ha sido rechazada.'}
      </p>
    </div>
  );
};

export default Congrats;

