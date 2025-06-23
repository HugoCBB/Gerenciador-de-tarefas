import './Login.css'
import Navbar from "../../components/NavBar";
import api from '../../services/Api';
import { useState } from 'react';


export default function Login () {

    const [form, setForm] = useState({
        email: '',
        senha: ''
    })


    const Login = async (e) => {
        e.preventDefault();
        try {
            const response = await api.post("/user/login", form, {
                withCredentials: true
            })
            
            localStorage.setItem("token", response.data.token)
            setLogin(response.data)
            console.log("Login realizado com sucesso")
        } catch (err) {
            alert(err.response?.data?.error)
        }

    }
    
    return (
        <>
        <section className='login'> 
            <Navbar/>   
            <section className="login__container">
                <form className="form-login" onSubmit={Login}>
                    <input 
                        type="email" 
                        placeholder="Digite seu email" 
                        value={form.nome}
                        onChange={(e) => setForm({...form, email:e.target.value})}    
                        />
                    <input
                        type="password" 
                        placeholder="Digite sua senha" 
                        value={form.senha}
                        onChange={(e) => setForm({...form, senha:e.target.value})}    

                    />
                    <button type="submit">Login</button>
                </form>
            </section>
        </section>
        </>
    );
}
