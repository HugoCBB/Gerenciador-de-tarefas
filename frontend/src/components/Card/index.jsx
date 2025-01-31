import './Card.css';

const Card = ({ titulo, descricao, dataCriacao }) => {
    return (
        <section>
        <a href="/">Home</a>
        <div className="card">
            <div className="card-content">
                <h3>{titulo}</h3>
                <p>{descricao}</p>
                <p>{dataCriacao}</p>
            </div>
        </div>
        </section>
    );
}

export default Card;