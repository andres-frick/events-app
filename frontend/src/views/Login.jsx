import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";



function Login() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (token) {
      navigate("/app");
    }
  }, []);


  const handleSubmit = async (e) => {
    e.preventDefault();

    const data = {
      username: username,
      password: password,
    };

    const resp = await fetch("http://localhost:4000/user/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    })

    if (resp.status === 200) {
      const { hash, rol } = await resp.json();
      localStorage.setItem("token", hash);
      localStorage.setItem("rol", rol);
      navigate("/app");

    }

  }


  return (
    <div>
      <center>
        <h2>Inicio de Sesion</h2>
        <form onSubmit={handleSubmit}>
          <div>
            <input
              type="text"
              placeholder="Usuario"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
            />
          </div>
          <input
            type="password"
            placeholder="Contraseña"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
          <div>
            <button type="submit">Ingresar</button>
          </div>
        </form>
      </center>
    </div>
  )
}

export default Login;