FROM golang:latest

WORKDIR /go/src

COPY ./src .

RUN go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]