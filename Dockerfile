FROM golang:1.23-alpine

WORKDIR /app

COPY . /app/

RUN go mod tidy && go build -o binary

EXPOSE 8080

ENTRYPOINT [ "/app/binary" ]