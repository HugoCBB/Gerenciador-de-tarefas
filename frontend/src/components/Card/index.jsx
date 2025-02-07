import './Card.css';

const Card = ({ titulo, descricao, dataCriacao, onDelete }) => {
    return (
        <section>
        <div className="card">
            <div className="card-content">
                <h3>{titulo}</h3>
                <p>{descricao}</p>
                <p>{dataCriacao}</p>
                <span className="delete-icon" onClick={onDelete}>🗑️</span>
            </div>
        </div>
        </section>
    );
}

export default Card;