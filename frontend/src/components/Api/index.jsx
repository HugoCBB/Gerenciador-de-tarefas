import { useEffect } from "react";
import axios from "axios";

const Api = () => {
    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await axios.get("http://localhost:8000", {
                withCredentials: true, // Envia o cookie de sessão
                headers: {
                    "Content-Type": "application/json",
                },
            });
                console.log(response.data);
            } catch (error) {
            console.error("Erro ao buscar dados:", error);
            }
        };
    
        fetchData();
    }, []);

    return (
        <div>
            <h1>Api</h1>
        </div>
    );
};

export default Api;