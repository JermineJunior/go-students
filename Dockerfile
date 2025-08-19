FROM golang:1.24 
WORKDIR /app

COPY go.* ./
COPY . .

RUN go mod download
COPY . .

RUN go build -o bin/app  .

EXPOSE 8080

CMD ["./bin/app"]