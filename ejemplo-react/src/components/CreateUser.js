import { React, useState } from 'react'
import '../css/CreateUser.css'
function CreateUser() {
    const [nombre, setnombre] = useState("")
    const [apellido, setapellido] = useState("")
    const [acepta, setacepta] = useState(false)

    const enviar = () => {
        var json = [
            nombre,
            apellido,
            acepta
        ]
        var datos = localStorage.getItem("usuarios")
        if (datos == null || datos === undefined) {
            localStorage.setItem("usuarios", JSON.stringify([json]))
        } else {
            datos = JSON.parse(datos)
            datos.push(json)
            console.log(datos)
            localStorage.setItem("usuarios", JSON.stringify(datos))
        }
        alert(JSON.stringify(json))
    }
    return (
        <div className="UserList">
            <br></br>
            <div className="ui segment container formulario form">

                <div className="field">
                    <label>First Name</label>
                    <input type="text" name="first-name" placeholder="First Name" onChange={e => setnombre(e.target.value)} />
                </div>
                <div className="field">
                    <label>Last Name</label>
                    <input type="text" name="last-name" placeholder="Last Name" onChange={e => setapellido(e.target.value)} />
                </div>
                <div className="field">
                    <div className="ui checkbox">
                        <input type="checkbox" tabIndex={0} onChange={e => setacepta(!acepta)} />
                        <label>I agree to the Terms and Conditions</label>
                    </div>
                </div>
                <button className="ui button" onClick={enviar} >Submit</button>
            </div>
        </div>
    )
}

export default CreateUser
