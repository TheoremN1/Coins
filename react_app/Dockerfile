FROM node:18.13.0-alpine as build_dev
RUN npm install -g pnpm

WORKDIR /app

COPY package.json .
COPY pnpm-lock.yaml .

RUN pnpm install

COPY . .
COPY .env.dev .env

RUN npm run build