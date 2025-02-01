import './Cadastrar.css';
import Navbar from "../../components/NavBar";

const Cadastrar = () => {
    return(
        <>
        <section className='cadastrar'> 
            <Navbar />
            <section className="cadastrar__container">
                <form className="form-cadastrar">
                    <input type="text" placeholder="Digite o seu nome" />
                    <input type="email" placeholder="Digite um email valido" />
                    <input type="password" placeholder="Digite uma senha" />
                    <input type="password" placeholder="Confirme a senha" />
                    <button type="submit">Cadastrar</button>
                </form>
            </section>
        </section>
        </>
    )
}

export default Cadastrar;