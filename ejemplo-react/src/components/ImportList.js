import React, { useEffect, useState } from 'react'
import Mosaico from './Mosaico'
import '../css/ImportList.css'
const axios = require('axios')

function ImportList() {
    const [productos, setproductos] = useState([])
    const [loading, setloading] = useState(false)
    useEffect(() => {
        async function obtener() {
            if (productos.length === 0) {
                const data = await axios.get('https://gorest.co.in/public-api/products');
                console.log(data.data.data)
                setproductos(data.data.data)
                setloading(true)
            }
        }
        obtener()
    });
    if (loading === false) {
        return (
            <div className="ui segment carga">
                <div className="ui active dimmer">
                    <div className="ui text loader">Loading</div>
                </div>
                <p />
            </div>
        )
    } else {
        return (
            <div className="ImportList">
                <br></br>
                <Mosaico productos={productos} />
            </div>
        )
    }
}

export default ImportList
