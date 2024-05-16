FROM golang:1.22 as builder
RUN  mkdir app
COPY go.mod ./
COPY go.sum ./
WORKDIR /app
CMD [ "app" ]