FROM node:lts-alpine3.14

WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH

COPY ./escape-room/package.json ./

COPY ./escape-room/package-lock.json ./

RUN npm install --silent -y

EXPOSE 3000

COPY ./escape-room ./

ENTRYPOINT ["npm", "start"]