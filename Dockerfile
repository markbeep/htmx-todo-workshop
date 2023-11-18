FROM golang:1.21.3-alpine3.18
WORKDIR /app

RUN go install github.com/a-h/templ/cmd/templ@latest
RUN go install github.com/cosmtrek/air@latest

COPY static/main.css static/main.css
COPY go.mod go.sum .air.toml ./
RUN go mod download

CMD air
