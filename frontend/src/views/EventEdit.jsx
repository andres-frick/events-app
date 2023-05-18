import React, { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";

function EventEdit() {
    const id = useParams().id;
    const [event, setEvent] = useState([]);

    const navigate = useNavigate();

    useEffect(() => {
        const token = localStorage.getItem("token");
        const resp = fetch(`http://localhost:4000/event/${id}`, {
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
                setEvent(data);
            }
            );
    }, []);

    const updateEvent = async (e) => {
        e.preventDefault();

        const token = localStorage.getItem("token");
        const resp = await fetch(`http://localhost:4000/event/${id}`, {
            method: "PUT",
            headers: {
                Authorization: `Bearer ${token}`,
                "Content-Type": "application/json",
            },
            body: JSON.stringify(event),
        });
        if (resp.status === 200) {
            alert("Event actualizado");
            navigate("/app");
        } else {
            alert("No se pudo actualizar el Evento");
        }

    };



    return (
        <div>
            <center>
                <h2>Editar Evento</h2>
                <form onSubmit={updateEvent}>
                    <label>Titulo: </label>
                    <input type="text" value={event.title || ""} onChange={(e) => setEvent({ ...event, title: e.target.value })} /> <br />
                    <label>Descripcion: </label>
                    <input type="text" value={event.short_description || ""} onChange={(e) => setEvent({ ...event, short_description: e.target.value })} /> <br />
                    <label>Detalle: </label>
                    <textarea rows={3} cols={40} value={event.long_description || ""} onChange={(e) => setEvent({ ...event, long_description: e.target.value })} /><br />
                    <label>Fecha y hora: </label>
                    <input type="text" value={event.event_date || ""} onChange={(e) => setEvent({ ...event, event_date: e.target.value })} /><br />
                    <label>Organizador: </label>
                    <input type="text" value={event.organizer || ""} onChange={(e) => setEvent({ ...event, organizer: e.target.value })} /><br />
                    <label>Estado: </label>
                    <select value={event.state || "publicada"} onChange={(e) => setEvent({ ...event, state: e.target.value })}>
                        <option value="publicada">Publicada</option>
                        <option value="borrador">Borrador</option>
                    </select><br /><br />
                    <button type="submit">Guardar</button>
                    <button type="button" onClick={() => navigate("/app")}>Cancel</button>
                </form>
            </center>
        </div>
    )
}

export default EventEdit;