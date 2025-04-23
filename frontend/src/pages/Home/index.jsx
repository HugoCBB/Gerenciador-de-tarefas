

import Navbar from "../../components/NavBar";
import "./Home.css"

export default function Home () {
    return (
        <section className="home">
            <Navbar/>
            <div className="home__container">
                <h1>Gerenciador de tarefas</h1>
                <p><strong>Organize sua vida e aumente sua produtividade</strong></p>
            </div>
        </section>
    )
}
