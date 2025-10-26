# Используем Node.js для сборки
FROM node:18-alpine as builde

WORKDIR /app

# Копируем package.json и package-lock.json для установки зависимостей
COPY vue/package*.json ./

RUN npm install

copy ./vue .
