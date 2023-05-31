FROM golang:latest

RUN go version
ENV GOPATH=/

WORKDIR /app

COPY . /app

RUN go mod download
RUN go build -o currencyapi ./cmd/main.go

CMD ["./currencyapi"]

