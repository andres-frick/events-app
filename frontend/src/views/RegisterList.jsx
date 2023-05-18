import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";

function RegisterList() {
    const [events, setEvents] = useState([]);
    const [status, setStatus] = useState("");

    const navigate = useNavigate();

    useEffect(() => {
        const token = localStorage.getItem("token");

        if (!token) {
            navigate("/");
        }


        const resp = fetch(`http://localhost:4000/event/registrations?status=${status}`, {
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
            }
            );
    }, [status]);

    return (
        <div>
            <h2>Tu inscripcion a Eventos</h2>
            <center>
                <select onChange={(e) => setStatus(e.target.value)}>
                    <option value="">Todos</option>
                    <option value="active">Activos</option>
                    <option value="completed">Completados</option>
                </select>
                <table>
                    <thead>
                        <tr>
                            <th>Titulo</th>
                            <th>Description</th>
                            <th>Fecha y hora</th>
                            <th>Organizador</th>
                        </tr>
                    </thead>
                    <tbody>
                        {events.map((event) => (
                            <tr key={event.ID}>
                                <td>{event.title}</td>
                                <td>{event.short_description}</td>
                                <td>{event.event_date}</td>
                                <td>{event.organizer}</td>
                            </tr>
                        ))}
                    </tbody>
                </table><br /><br />
            </center>
            <button onClick={() => navigate("/app")}>Volver</button>
        </div>
    );

}

export default RegisterList;