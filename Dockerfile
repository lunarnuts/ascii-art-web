FROM golang:1.15

LABEL "authors"="lunarnuts"
LABEL "description"="ascii-art-web"

WORKDIR /ascii-art-web
COPY . .

EXPOSE 8080

RUN go build -o main .
CMD ["./main"]