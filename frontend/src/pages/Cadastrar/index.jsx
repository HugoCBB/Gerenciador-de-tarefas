import './Cadastrar.css';
import Navbar from "../../components/NavBar";
import { useState } from 'react';
import axios from 'axios';

const Cadastrar = () => {

    const [usuarios, setUsuarios] = useState([]);
    const [cadastraUsuario, setCadastrarUsuario] = useState({
        nome: '',
        email: '',
        senha: '',
    })


    const postCadastrarUsuario = async () => {
        try {
            const response = await axios.post("http://localhost:8000/api/user/cadastrar", cadastraUsuario, {
                withCredentials: true
            })
            setCadastrarUsuario(response.data);
            setUsuarios([...usuarios, cadastraUsuario]);
            console.log(usuarios);
        } catch (error) {
            console.error("Erro ao cadastrar usuario:", error);
        }
    }

    return(
        <>
        <section className='cadastrar'> 
            <Navbar />
            <section className="cadastrar__container">
                <form className="form-cadastrar" type="submit">
                    <input type="text" 
                    placeholder="Digite o seu nome" 
                    value={cadastraUsuario.nome}
                    onChange={(e) => setCadastrarUsuario({...cadastraUsuario, nome: e.target.value})}
                    />
                    
                    <input type="email" 
                    placeholder="Digite um email valido" 
                    value={cadastraUsuario.email}
                    onChange={(e) => setCadastrarUsuario({...cadastraUsuario, email: e.target.value})}
                    />
                    
                    <input type="password" 
                    placeholder="Digite uma senha"
                    value={cadastraUsuario.senha} 
                    onChange={(e) => setCadastrarUsuario({...cadastraUsuario, senha: e.target.value})}
                    />
                    
                    <input type="password" 
                    placeholder="Confirme a senha" />
                    <button type="submit" onClick={postCadastrarUsuario}>Cadastrar</button>
                </form>
            </section>
        </section>
        </>
    )
}

export default Cadastrar;