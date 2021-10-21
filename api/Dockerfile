# Base image
ARG GO_IMAGE_VERSION=1.16-alpine
FROM golang:${GO_IMAGE_VERSION}}

# cd into /app
WORKDIR /app

# copy the module, and sum file then installing necessary dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# copy all of the files in the /api folder. .dockerignore will exclude unneeded file
COPY . .

# build the binary
RUN go build -o /escape-room-revamp-api

# run the binary
CMD [ "/escape-room-revamp-api" ]