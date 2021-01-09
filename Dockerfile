FROM golang:latest

ADD . /app

WORKDIR /app

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

RUN chmod +x ./Golang_Rest-API_CRUD_echo

EXPOSE 8080

CMD ./Golang_Rest-API_CRUD_echo