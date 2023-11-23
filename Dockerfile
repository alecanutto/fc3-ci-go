FROM golang:1.21

WORKDIR /app

COPY . .

RUN go mod download

ARG FILE_TO_BUILD="cmd/main.go"

RUN go build -o main ${FILE_TO_BUILD}

CMD ["./main"]