
FROM golang:latest

WORKDIR /order


COPY . .

RUN go mod download
COPY *.go ./
EXPOSE 8080
RUN go build -o main ./cmd

CMD ["./main"]
 
