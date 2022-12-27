FROM golang:1.19-alpine3.17 as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /content-server

ENV HTTP_PORT=8080

EXPOSE $HTTP_PORT

CMD [ "/content-server" ]