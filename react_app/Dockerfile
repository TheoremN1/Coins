FROM node:18.13.0-alpine

WORKDIR /app

COPY . .

RUN npm i

EXPOSE 80

CMD ["npm", "run", "dev"]