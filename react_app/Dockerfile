FROM node:18.13.0-alpine

WORKDIR /app

COPY . .

RUN npm i

EXPOSE ${REACT_APP_PORT}

CMD ["npm", "run", "dev"]