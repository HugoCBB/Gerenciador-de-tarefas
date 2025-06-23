FROM golang:1.24

WORKDIR /gerenciador-tarefas

COPY go.mod go.sum ./

RUN go mod download

RUN go install github.com/air-verse/air@latest

COPY . .

EXPOSE 8000

CMD ["air"]