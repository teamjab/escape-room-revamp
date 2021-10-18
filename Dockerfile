FROM node:latest

WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH

COPY ./escape-room/package.json ./

COPY ./escape-room/package-lock.json ./

RUN npm install

COPY ./escape-room ./

ARG PORT

RUN npm run build

CMD ["npm", "run", "start"]