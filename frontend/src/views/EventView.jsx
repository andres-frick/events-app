import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";


function EventView() {
    const id = useParams().id;
    const navigate = useNavigate();
    const [status, setStatus] = useState("")

    const [event, setEvent] = useState([]);

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
                EventStatus(data.event_date);
            }
            );

    }, []);

    const EventStatus = (date) => {
        const today = new Date();
        const eventDate = new Date(date);
        if (eventDate > today) {
            setStatus("Activo")
        } else {
            setStatus("Finalizado")
        }
    }

    return (
        <div>
            <center>
                <h2>Evento</h2>
                <h3>Titulo: {event.title}</h3>
                <p>Descripcion: {event.short_description}</p>
                <p>Detalle: {event.long_description}</p>
                <p>Fecha y Hora: {event.event_date}</p>
                <p>Organizador: {event.organizer}</p>
                <p>Estado: {status}</p>
                <button onClick={() => navigate("/app")}>Volver</button>
            </center>
        </div>


    )
}

export default EventView;