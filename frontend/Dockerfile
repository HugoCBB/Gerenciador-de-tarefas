FROM node:20

WORKDIR /gerenciador-tarefas

COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 5173

CMD [ "npm", "run", "dev" ]