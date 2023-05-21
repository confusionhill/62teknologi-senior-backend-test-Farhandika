FROM golang:1.18-alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN  go build -o build/http-rest application/rest/*.go

EXPOSE 50051

CMD [ "./build/http-rest" ]