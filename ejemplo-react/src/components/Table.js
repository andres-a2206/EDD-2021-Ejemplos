import React from 'react'
import Fila from './Fila'
function Table(props) {
    return (
        <>
            <br />
            <div className="ui segment container ">
                <table className="ui celled table segment content">
                    <thead>
                        <tr>
                            {props.encabezados.map((dato) => (
                                <th>{dato}</th>
                            ))}
                        </tr>
                    </thead>
                    <tbody>
                        {props.data.map((dato, index) => (
                            <Fila
                                index={index}
                                data={dato}
                            />
                        ))}
                    </tbody>
                </table>
            </div>
        </>
    )
}

export default Table
