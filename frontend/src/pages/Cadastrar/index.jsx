import './Cadastrar.css';
import Navbar from "../../components/Navbar";
import { useState } from 'react';
import api from '../../services/Api';

export default function Cadastrar () {

    const [form, setForm] = useState({
        nome: '',
        email: '',
        senha: '',
    })


    const Cadastro = async (e) => {
        e.preventDefault();
        try {
            const response = await api.post("/user/cadastrar", form, {
                withCredentials: true
            })
            setForm(response.data);
            alert(response.data.status)
        
        } catch (err) {
            alert(err.response?.data?.error)
        }
    }

    return(
        <>
        <section className='cadastrar'> 
            <Navbar />
            <section className="cadastrar__container">
                <form className="form-cadastrar"  onSubmit={Cadastro}>
                    
                    <input type="text" 
                        placeholder="Digite o seu nome" 
                        value={form.nome}
                        onChange={(e) => setForm({...form, nome: e.target.value})}
                    />
                    
                    <input type="email" 
                        placeholder="Digite um email valido" 
                        value={form.email}
                        onChange={(e) => setForm({...form, email: e.target.value})}
                    />
                    
                    <input type="password" 
                        placeholder="Digite uma senha"
                        value={form.senha} 
                        onChange={(e) => setForm({...form, senha: e.target.value})}
                    />
                    
                    <input 
                        type="password" 
                        placeholder="Confirme a senha" />
                    <button type="submit">Cadastrar</button>
                </form>
            </section>
        </section>
        </>
    )
}
