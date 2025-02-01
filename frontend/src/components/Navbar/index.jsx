import './Navbar.css';

const Navbar = () => {
    return (
        <header className="navbar">
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
    );
}


export default Navbar;