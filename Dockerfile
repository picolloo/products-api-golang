FROM golang:latest

LABEL maintainer = "Lucas Picollo <lucaspicolloo@gmail.com>"

WORKDIR /app/current

COPY go.mod ./
RUN go mod download

COPY . .

EXPOSE 3000

RUN GOOS=linux go build .