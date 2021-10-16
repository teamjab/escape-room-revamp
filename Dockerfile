FROM node:lts-alpine3.14

WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH

COPY ./escape-room/package.json ./

COPY ./escape-room/package-lock.json ./

RUN npm install --silent -y

COPY ./escape-room ./

ARG PORT

ENTRYPOINT ["npm", "start", "--port", $PORT]