import React, { useState, useEffect } from 'react';
import './Tarefas.css';
import Card from '../../components/card';
import api from '../../services/Api';

export default function Tarefas () {
    const [tarefas, setTarefas] = useState([]);
    const [novaTarefa, setNovaTarefa] = useState({titulo:'', descricao:''})

    useEffect(() => {
        getTarefas();
    }, []);

    const getTarefas = async () => {
        try {
            const response = await api.get("/tarefas/", {
                withCredentials: true,
            });
            setTarefas(response.data);
        } catch (err) {
            console.error("Erro ao buscar tarefas:", err);
        
        }
    };

    const postTarefas = async () => {
        try {
            if (novaTarefa.descricao != "" && novaTarefa.titulo != "") {
                const response = await api.post("/tarefas/adicionar", novaTarefa, {
                withCredentials: true,
                });
                
                setTarefas([...tarefas, response.data]);
                setNovaTarefa({ titulo: '', descricao: '' });
            } else {
                alert("Preencha os campos corretamente");
            }
        } catch (err) {
            alert("Erro ao adicionar tarefa: ", err);
        }
    };

    const deleteTarefas = async (id) => {
        if (window.confirm("Tem certeza que deseja deletar essa tarefa?")) {
            try {
                await api.delete(`/tarefas/deletar/${id}`, {
                    withCredentials: true,
                })
                setTarefas(tarefas.filter(tarefa => tarefa.id !== id));
            } catch (err) {
                alert("Erro ao deletar tarefa: ",err)
                
            }
        }
    }
    return (
        <section className="tarefas">
            <div className="tarefas__container">
                <h1>Lista de Tarefas</h1>
                <a href='/'>Home</a>
            <div>
                <div className='tarefas__inputs'>
                    <input
                        type="text"
                        placeholder="Título"
                        value={novaTarefa.titulo}
                        onChange={(e) => setNovaTarefa({ ...novaTarefa, titulo: e.target.value })}
                    />
                    <input
                        type="text"
                        placeholder="Descrição"
                        value={novaTarefa.descricao}
                        onChange={(e) => setNovaTarefa({ ...novaTarefa, descricao: e.target.value })}
                    />
                    <button onClick={postTarefas}>Adicionar Tarefa</button>
                </div>
            </div>
                <div className="tarefas__cards">
                    {tarefas.map(tarefa => (
                        <Card
                            key={tarefa.id}
                            titulo={tarefa.titulo}
                            descricao={tarefa.descricao}
                            dataCriacao={tarefa.dataCriacao}
                            onDelete={() => deleteTarefas(tarefa.id)}
                        />
                    ))}
                    
                </div>
            </div>
        </section>
    );
};
