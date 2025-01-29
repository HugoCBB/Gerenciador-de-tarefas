import Home from "./pages/home"
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import Tarefas from "./pages/Tarefas";
import Login from "./pages/Login";
import Cadastro from "./pages/Cadastro";

function App() {

  return (
    <>
      <Switch>
        <Route path="/" component={Home} exact />
        <Route path="/tarefas" component={Tarefas}  />
        <Route path="/login" component={}  />
        <Route path="/cadastro" component={Tarefas}  />
      </Switch>
      
    </>
  )
}

export default App
