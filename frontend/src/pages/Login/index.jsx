import './Login.css'
import Navbar from "../../components/NavBar";

const Login = () => {
    return (
        <>
        <section className='login'> 
            <Navbar/>   
            <section className="login__container">
                <form className="form-login">
                    <input type="email" placeholder="Digite seu email" />
                    <input type="password" placeholder="Digite sua senha" />
                    <button type="submit">Login</button>
                </form>
            </section>
        </section>
        </>
    );
}

export default Login;