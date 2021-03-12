import React from 'react'
const axios = require('axios')

function ViewUser(props) {
    const usuarios = localStorage.getItem('usuarios')
    let id = props.match.params.id
    const usuario = JSON.parse(usuarios)[id]
    const mandardatos = () => {
        let json = { msg: ["EDD", "2021"] }
        axios.post('http://localhost:4000/agregar', json)
            .then(function (response) {
                console.log(response);
            })
            .catch(function (error) {
                console.log(error);
            });
    }
    return (
        <div className="ui raised very padded text container segment">
            <h2 className="ui header">{usuario[0]} </h2>
            <h3 className="ui ">Enviar mensaje</h3>
            <button className="ui inverted red button" onClick={mandardatos}>Standard</button>
        </div>
    )
}

export default ViewUser
