import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

function EventList() {
    const [events, setEvents] = useState([]);
    const [show, setShow] = useState("none");
    const [filters, setFilters] = useState("");

    const navigate = useNavigate();


    useEffect(() => {
        getEvents({});

    }, []);

    const subscribe = async (id) => {
        const token = localStorage.getItem("token");
        const resp = await fetch(`http://localhost:4000/event/${id}/register`, {
            method: "POST",
            headers: {
                Authorization: `Bearer ${token}`,
            },
        });
        if (resp.status === 201) {
            alert("Te inscribiste al Evento");
        } else {
            alert("No se pudo inscribir al Evento, el evento Finalizo");
        }
    };

    const edit = async (id) => {
        navigate(`/app/event/${id}`);
    }

    const remove = async (id) => {
        const token = localStorage.getItem("token");
        const resp = await fetch(`http://localhost:4000/event/${id}`, {
            method: "DELETE",
            headers: {
                Authorization: `Bearer ${token}`,
            },
        });
        if (resp.status === 204) {
            alert("Evento eliminado");
            setEvents(events.filter((event) => event.ID !== id));
        } else {
            alert("No se pudo elminar el Evento");
        }
    };

    const viewEvent = async (id) => {
        navigate(`/event/${id}`);
    }

    const logOut = () => {
        localStorage.removeItem("token");
        localStorage.removeItem("rol");
        navigate("/");
    };

    const selFilters = async (e) => {
        e.preventDefault();
        getEvents(filters);


    };

    const getEvents = async (filters) => {
        const token = localStorage.getItem("token");

        if (!token) {
            navigate("/");
        }

        const query = Object.entries(filters)
            .filter((entry) => entry[1] !== "")
            .map((entry) => `${entry[0]}=${entry[1]}`)
            .join("&");

        
        const rol = localStorage.getItem("rol");

        const resp = fetch(`http://localhost:4000/event?${query}`, {
            headers: {
                Authorization: `Bearer ${token}`,
            },
        })
            .then((resp) => {
                if (resp.status === 200) {
                    return resp.json();
                }
            })
            .then((data) => {
                setEvents(data);
                setShow(rol === "admin" ? "block" : "none");
            }
            );
    }


    return (
        <div>
            <center>
                <h1>Eventos</h1>
                <form onSubmit={selFilters}>
                    <label>Fecha desde: </label>
                    <input type="datetime-local" id="start" onChange={(e) => setFilters({ ...filters, start: e.target.value + ":00Z"})} />
                    <label>Fecha hasta: </label>
                    <input type="datetime-local" id="end" onChange={(e) => setFilters({ ...filters, end: e.target.value + ":00Z"})} />
                    <label>Titulo: </label>
                    <input type="text" id="title" onChange={(e) => setFilters({ ...filters, title: e.target.value })} />
                    <label>Estado: </label>
                    <select id="state" onChange={(e) => setFilters({ ...filters, state: e.target.value })}>
                        <option value="">Todos</option>
                        <option value="publicada">Publicado</option>
                        <option value="borrador">Borrador</option>
                    </select>
                    <button type="submit">Filtrar</button>
                    <button onClick={() => setFilters({})}>Limpiar</button>
                </form>
            </center>

            <br />
            <ul>
                {events.map((event) => (
                    <li key={event.ID}>
                        <h3>Titulo: {event.title}</h3>
                        <p>Fecha de evento: {event.event_date}</p>
                        <p>Organizador: {event.organizer}</p>
                        <button onClick={() => subscribe(event.ID)}>Inscribirse</button>
                        <button onClick={() => viewEvent(event.ID)}>Detalle</button>
                        <div style={{ display: show }}>
                            <button onClick={() => edit(event.ID)}>Editar</button>
                            <button onClick={() => remove(event.ID)}>Borrar</button>
                        </div>
                    </li>
                ))}
            </ul>
            <center>
                <button onClick={() => { navigate('/event/registrations') }}>Tus Eventos</button>
                <div style={{ display: show }}>
                    <button onClick={() => navigate("/event/new")}>Nuevo Evento</button>
                </div>

                <div>
                    <button onClick={() => logOut()}>Salir</button>
                </div>
            </center>
        </div>
    )
}

export default EventList;