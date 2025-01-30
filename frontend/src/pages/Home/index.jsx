
import "./Home.css"
const Home = () => {
    return (
        <section className="home">
            <header className="home__navbar">
                <nav>
                    <ul>
                        <li>
                            <a href="/">Home</a>
                        </li>
                        <li>
                            <a href="/login">Login</a>
                        </li>
                        <li>
                            <a href="/cadastrar">Cadastro</a>
                        </li>
                        <li>
                            <a href="/tarefas">Tarefas</a>
                        </li>
                    </ul>
                </nav>
            </header>
            <div className="home__container">
                <h1>Gerenciador de tarefas</h1>
                <p><strong>Organize sua vida e aumente sua produtividade</strong></p>
            </div>

        </section>
    )
}

export default Home;