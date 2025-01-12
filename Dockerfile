FROM golang:1.23.4-alpine

WORKDIR  /app

RUN apk --no-cache add git

COPY product-api/go.mod product-api/go.sum  ./


RUN go mod download

COPY product-api/ .

EXPOSE 8080

ENTRYPOINT ["go" , "run" , "main.go"]
