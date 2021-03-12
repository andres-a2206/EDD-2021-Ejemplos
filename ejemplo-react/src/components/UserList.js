import React, { useEffect, useState } from 'react'
import Tabla from './Table'
import '../css/Userlist.css'
function UserList() {
    const encabezados = ["id", "Nombre", "Apellido", "De acuerdo","Acciones"]
    const [listado, setlistado] = useState([
        ["Marvin", "Martinez", true],
        ["Yaiza", "Pineda", false]
    ])
    useEffect(() => {
        let data = localStorage.getItem('usuarios')
        if (data != null) {
            setlistado(JSON.parse(data))
        }
    }, [])
    return (
        <div className="Userlist">
            <Tabla data={listado}
                encabezados={encabezados}
            />
        </div>

    )
}

export default UserList
