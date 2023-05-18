import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

function NewEvent() {
    const [event, setEvent] = useState([]);
    const navigate = useNavigate();
    
    useEffect(() => {
        const token = localStorage.getItem("token");
        if (!token) {
            navigate("/");
        }
    }, []);

    const createEvent = async (e) => {
        e.preventDefault();
        const token = localStorage.getItem("token");

        const resp = await fetch(`http://localhost:4000/event`, {
            method: "POST",
            headers: {
                Authorization: `Bearer ${token}`,
                "Content-Type": "application/json",
            },
            body: JSON.stringify(event),
        });
        if (resp.status === 201) {
            alert("Event created");
            navigate("/app");
        } else {
            alert("Could not create event");
        }

    };


    return (
        <div>
            <center>
                <h2>Nuevo Evento</h2>
                <form onSubmit={createEvent}>
                    <input type="text" placeholder="Titulo" onChange={(e) => setEvent({ ...event, title: e.target.value })} /> <br />
                    <input type="text" placeholder="Description" onChange={(e) => setEvent({ ...event, short_description: e.target.value })} /> <br />
                    <textarea rows={3} cols={40} placeholder="Detalle" onChange={(e) => setEvent({ ...event, long_description: e.target.value })} /><br />
                    <input type="text" placeholder="Fecha y hora" onChange={(e) => setEvent({ ...event, event_date: e.target.value })} /><br />
                    <input type="text" placeholder="Organizador" onChange={(e) => setEvent({ ...event, organizer: e.target.value })} /><br />
                    <input type="text" placeholder="Lugar" onChange={(e) => setEvent({ ...event, location: e.target.value })} /><br />
                    <select value={event.state} onChange={(e) => setEvent({ ...event, state: e.target.value })}>
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

export default NewEvent;