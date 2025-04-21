
## 📂 Estrutura do Projeto

### Backend
O backend está localizado na pasta `backend` e utiliza o framework [Gin](https://gin-gonic.com/) para criar APIs RESTful e o ORM [GORM](https://gorm.io/) para interagir com o banco de dados.

- **`cmd/server`**: Contém o ponto de entrada principal do backend (`main.go`).
- **`config`**: Configurações do sistema, como carregamento de variáveis de ambiente.
- **`controllers`**: Controladores responsáveis por gerenciar as requisições HTTP.
- **`database`**: Configuração e conexão com o banco de dados.
- **`middleware`**: Middlewares para autenticação e validação.
- **`models`**: Modelos que representam as tabelas do banco de dados.
- **`routers`**: Configuração das rotas da API.

### Frontend
O frontend está localizado na pasta `frontend` e utiliza o [Vite](https://vitejs.dev/) como ferramenta de build.

- **`src/components`**: Componentes reutilizáveis da interface, como `Card` e `Navbar`.
- **`src/pages`**: Páginas principais da aplicação, como `Home`, `Cadastrar`, `Login` e `Tarefas`.
- **`App.jsx`**: Componente principal que define a estrutura da aplicação.

---

## 🚀 Funcionalidades

### Backend
- Hot reload com [Air](https://github.com/cosmtrek/air).
- Cadastro de usuários com validação de campos obrigatórios.
- Hash de senhas utilizando `bcrypt`.
- Relacionamento entre usuários e tarefas.
- APIs RESTful para gerenciar usuários e tarefas.

### Frontend
- Interface para gerenciar tarefas e usuários.
- Integração com o backend para consumir APIs.
- Navegação entre páginas.

---

## 🛠️ Tecnologias Utilizadas

### Backend
- **Linguagem**: Go (Golang)
- **Framework**: Gin
- **ORM**: GORM
- **Banco de Dados**: PostgreSQL
- **Autenticação**: Bcrypt
- **Variáveis de Ambiente**: Godotenv
- **Hot Reload**: Air

### Frontend
- **Linguagem**: JavaScript (ES6+)
- **Framework**: React
- **Build Tool**: Vite

---

## ⚙️ Configuração do Ambiente

### Pré-requisitos
- [Go](https://golang.org/) (versão 1.19 ou superior)
- [Node.js](https://nodejs.org/) (versão 16 ou superior)
- [PostgreSQL](https://www.postgresql.org/)
- [Air](https://github.com/cosmtrek/air) (para hot reload no backend)
