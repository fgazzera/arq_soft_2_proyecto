import React, { useContext, useEffect, useState, useCallback } from 'react';
import { AuthContext } from './login/auth';
import './estilo/reservas_cliente.css';

const handleVolver = () => {
  window.history.back();
};

const ReservasCliente = () => {
  const [reservations, setReservations] = useState(null);
  const [reservasFiltradas, setReservasFiltradas] = useState([]);
  const [hoteles, setHoteles] = useState([]);
  const { isLoggedCliente } = useContext(AuthContext);
  const [hotelFiltrado, setHotelFiltrado] = useState('');
  const [startDateFilter, setStartDateFilter] = useState('');
  const [endDateFilter, setEndDateFilter] = useState('');

  const getHoteles = useCallback(async () => {
    try {
      const hotelesArray = [];
      for (let i = 0; i < reservations.length; i++) {
        const reserva = reservations[i];
        const request = await fetch(`http://localhost:8090/cliente/hotel/${reserva.hotel_id}`);
        const response = await request.json();
        hotelesArray.push({ id: response.id, nombre: response.nombre });
      }
      const uniqueHotels = Array.from(new Set(hotelesArray.map((hotel) => hotel.id))).map((id) => {
        return hotelesArray.find((hotel) => hotel.id === id);
      });
      setHoteles(uniqueHotels);
    } catch (error) {
      console.log("No se pudieron obtener los hoteles:", error);
    }
  }, [reservations]);

  const getReservations = useCallback(async () => {
    if (isLoggedCliente) {
      const accountId = localStorage.getItem("id_cliente");
      try {
        const request = await fetch(`http://localhost:8090/cliente/reservas/${accountId}`);
        const response = await request.json();
        if (response) {
          setReservations(response);
          setReservasFiltradas(response);
        } else {
          setReservations([]);
          setReservasFiltradas([]);
        }
      } catch (error) {
        console.log("No se pudieron obtener las reservas:", error);
      }
    } else {
      window.location.href = '/';
    }
  }, [isLoggedCliente]);
  

  const getReservasFiltradas = useCallback(async () => {
    try {
      const startDateObj = new Date(startDateFilter);
      const endDateObj = new Date(endDateFilter);
      let reservasArray = [];
  
      if (startDateFilter && endDateFilter) {
        const request = await fetch(`http://localhost:8090/cliente/reservas-por-fecha/${startDateObj.getFullYear()}/${startDateObj.getMonth() + 1}/${startDateObj.getDate() + 1}/${endDateObj.getFullYear()}/${endDateObj.getMonth() + 1}/${endDateObj.getDate() + 1}`);
        const response = await request.json();
        reservasArray = response;
        reservasArray = reservasArray.filter((reserva) => hotelFiltrado === '' || hotelFiltrado === 0 || hotelFiltrado === reserva.hotel_id);
      } else {
        reservasArray = reservations.filter((reserva) => hotelFiltrado === '' || hotelFiltrado === 0 || hotelFiltrado === reserva.hotel_id);
      }
  
      setReservasFiltradas(reservasArray);
    } catch (error) {
      console.log("No se pudieron obtener las reservas:", error);
    }
  }, [reservations, hotelFiltrado, startDateFilter, endDateFilter]);

  useEffect(() => {
    getReservations();
  }, [getReservations]);

  useEffect(() => {
    getReservasFiltradas();
  }, [getReservasFiltradas]);

  useEffect(() => {
    getHoteles();
  }, [getHoteles]);

  const handleHotelFilterChange = (hotelId) => {
    setHotelFiltrado(hotelId);
  };

  const handleStartDateFilterChange = (event) => {
    setStartDateFilter(event.target.value);
    const selectedStartDateObj = new Date(event.target.value);
    const endDateObj = new Date(endDateFilter);
    if (selectedStartDateObj > endDateObj) {
      setEndDateFilter('');
      alert("Fechas no validas");
    }
  };

  const handleEndDateFilterChange = (event) => {
    setEndDateFilter(event.target.value);
    const startDateObj = new Date(startDateFilter);
    const selectedEndDateObj = new Date(event.target.value);
    if (startDateObj > selectedEndDateObj) {
      setEndDateFilter('');
      alert("Fechas no validas");
    }
  };

  return (
    <div className="reservations-container1">
      <div className="reservations-container2">
      <div className="filters-container">
        <div>
          <h6 htmlFor="hotelFilter">Filtro Reservas:</h6>
          <div id="hotelFilter" className="hotel-filter-container">
            <button className="hotel-filter-button" onClick={() => handleHotelFilterChange(0)}>
              Todos los hoteles
            </button>
            {hoteles.map((hotel) => (
              <button
                className="hotel-filter-button"
                key={hotel.id}
                value={hotel.id}
                onClick={() => handleHotelFilterChange(hotel.id)}
              >
                {hotel.nombre}
              </button>
            ))}
          </div>
        </div>
          <div className="contdeFechas">
            <div className="date-picker">
              <label htmlFor="startDateFilter">Fecha de inicio:</label>
              <input type="date" id="startDateFilter" value={startDateFilter} onChange={handleStartDateFilterChange} />
            </div>
            <div className="date-picker">
              <label htmlFor="endDateFilter">Fecha de fin:</label>
              <input type="date" id="endDateFilter" value={endDateFilter} onChange={handleEndDateFilterChange} />
            </div>
          </div>
        </div>
        <h4>Datos de tus reservas:</h4>
        <div className="scroll-container">
          {reservasFiltradas.length ? (
            reservasFiltradas.map((reservation) => {
              const hotel = hoteles.find((hotel) => hotel.id === reservation.hotel_id);
              const fechaInicio = `${reservation.dia_inicio}/${reservation.mes_inicio}/${reservation.anio_inicio}`;
              const fechaFin = `${reservation.dia_final}/${reservation.mes_final}/${reservation.anio_final}`;
              return (
                <div className="reservation-card" key={reservation.ID}>
                  <p>Hotel: {hotel ? hotel.nombre : 'Hotel desconocido'}</p>
                  <p>Fecha de llegada: {fechaInicio}</p>
                  <p>Fecha de fin: {fechaFin}</p>
                  <p>Gracias por elegirnos!</p>
                </div>
              );
            })
          ) : (
            <p>No tienes reservas</p>
          )}
        </div>
        <button className="botonBack" onClick={handleVolver}>
          🔙
        </button>
      </div>
    </div>
  );
};
export default ReservasCliente;



