import React from 'react'
import Navbar from './components/NavBar'
import { BrowserRouter as Router, Route } from 'react-router-dom'
import UserList from './components/UserList'
import CreateUser from './components/CreateUser'
import ImportList from './components/ImportList'
import ViewUser from './components/ViewUser'
import Arbol from './components/Arbol'
import './App.css'

function App() {
  return (
        <Router>
          <Navbar />
          <Route path="/listado" component={UserList} />
          <Route path="/formulario" component={CreateUser} />
          <Route path="/listado2" component={ImportList} />
          <Route path="/view/:id" component={ViewUser} />
          <Route path="/arbol" component={Arbol} />
        </Router>
  )
}

export default App
