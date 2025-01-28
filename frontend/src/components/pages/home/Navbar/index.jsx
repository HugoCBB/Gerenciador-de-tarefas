import "./Navbar.css"
const Navbar = () => {
    return (
        <header className="home__navbar">
        <nav>
            <ul>
                <li>
                    <a href="/">Home</a>
                </li>
                <li>
                    <a href="/tarefas">Tarefas</a>
                </li>
                <li>
                    <a href="/tarefas">Login</a>
                </li>
                <li>
                    <a href="/tarefas">Cadastrar</a>
                </li>
            </ul>
        </nav>
    </header>
    )
}

export default Navbar;