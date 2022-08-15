FROM golang:1.16-alpine3.15
RUN mkdir "/app"

ADD . /app
WORKDIR /app

RUN go build -o main

EXPOSE 8080

CMD ["/app/main"]