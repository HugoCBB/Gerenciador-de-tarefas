import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './Tarefas.css';
import Card from '../../components/card';

const Tarefas = () => {
    const [tarefas, setTarefas] = useState([]);

    useEffect(() => {
        getTarefas();
    }, []);

    const getTarefas = async () => {
        try {
            const response = await axios.get("http://localhost:8000/tarefa", {
                withCredentials: true,
            });
            setTarefas(response.data);
            console.log(response.data);
        } catch (err) {
            console.error("Erro ao buscar tarefas:", err);
        }
    };

    return (
        <section className="tarefas">
            <div className="tarefas__container">
                <h1>Lista de Tarefas</h1>
                <div className="tarefas__cards">
                    {tarefas.map(tarefa => (
                        <Card
                            key={tarefa.id}
                            titulo={tarefa.titulo}
                            descricao={tarefa.descricao}
                            dataCriacao={tarefa.data_criacao}
                        />
                    ))}
                </div>
            </div>
        </section>
    );
};

export default Tarefas;