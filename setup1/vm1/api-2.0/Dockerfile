FROM node:18

RUN mkdir app

COPY package*.json ./app/

WORKDIR /app

COPY . .

RUN npm install

CMD ["node","app.js"]