FROM node:alpine

WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH

COPY ./escape-room/package.json ./

COPY ./escape-room/package-lock.json ./

RUN npm install

COPY ./escape-room ./

ARG PORT

CMD ["npm", "run", "start", "--port", $PORT]