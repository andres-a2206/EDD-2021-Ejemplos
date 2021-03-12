import React from 'react'
import Tree from 'react-tree-graph';
import 'react-tree-graph/dist/style.css'
let data = {
    name: 'Parent',
    children: [{
        name: 'Child One',
    }, {
        name: 'Child Two'
    }]
};
let fileReader;
const handleFileRead=(e)=>{
    const content=fileReader.result;
    console.log(content)
}
const handleFileChosen=(file)=>{
    fileReader=new FileReader();
    fileReader.onloadend=handleFileRead;
    fileReader.readAsText(file)
}
function Arbol() {
    return (
        <input type='file'
            id='file'
            className='input-file'
            accept='.json'
            onChange={e=>handleFileChosen(e.target.files[0])}
        >
        </input>
    )
}

export default Arbol
