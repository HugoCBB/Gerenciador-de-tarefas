import Home from "./pages/home"
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Tarefas from "./pages/Tarefas";
import Login from "./pages/Login";
import Cadastrar from "./pages/Cadastrar";

function App() {

  return (
    <>
      <Router>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/tarefas" element={<Tarefas />} />
          <Route path="/login" element={<Login />} />
          <Route path="/cadastrar" element={<Cadastrar />} />
        </Routes>
    </Router>
    </>
  )
}

export default App
