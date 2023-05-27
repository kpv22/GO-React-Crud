import { useEffect, useState } from "react";

function App() {
  const [user, setUser] = useState({
    first_name: "",
    last_name: "",
    email: ""
  });
  const [listUsers, setlistUsers] = useState([])

useEffect(()=>{
  async function loadUsers(){
    const response = await fetch(import.meta.env.VITE_API + '/users');
    const data = await response.json()
    setlistUsers(data)
  }
  loadUsers()
},[listUsers])


  const handleChange = (e) => {
    setUser({
      ...user,
      [e.target.name]: e.target.value
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      // Realizar la solicitud fetch aquí usando los datos del usuario
      const response = await fetch(import.meta.env.VITE_API + '/users', {
        method: "POST",
        body: JSON.stringify(user),
        headers: {
          "Content-Type": "application/json"
        }
      });
      const data = await response.json();
  
      // Lógica adicional después de la respuesta del servidor
      console.log(data);
      setUser({
        first_name: "",
        last_name: "",
        email: ""
      });
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          name="first_name"
          placeholder="Coloca tu nombre"
          value={user.first_name}
          onChange={handleChange}
        />
        <input
          type="text"
          name="last_name"
          placeholder="Coloca tu apellido"
          value={user.last_name}
          onChange={handleChange}
        />
        <input
          type="email"
          name="email"
          placeholder="Coloca tu correo electrónico"
          value={user.email}
          onChange={handleChange}
        />
        <button type="submit">Guardar</button>
      </form>

      <ul>{listUsers.map(user => (
        <li key={user.ID}>{user.first_name}</li>
      ))}</ul>
    </div>
  );
}

export default App;