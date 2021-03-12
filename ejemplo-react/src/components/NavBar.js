import { React, useState } from 'react'
import { Link } from 'react-router-dom';
import { Menu } from "semantic-ui-react";

import '../css/Nav.css'
const colores = ['blue', 'orange', 'red']
const opciones = ['Crear Usuario', 'Lista de productos', 'Ver usuarios']
const url = ['/formulario', "/listado2", "/listado", "/"]

function NavBar() {
    const [activo, setActivo] = useState(colores[0])
    return (
        <Menu inverted className="Nav">
            {colores.map((c, index) => (
                <Menu.Item as={Link} to={url[index]}
                    key={c}
                    name={opciones[index]}
                    active={activo === c}
                    color={c}
                    onClick={() => setActivo(c)}
                />
            ))}
        </Menu>

    )
}
export default NavBar
