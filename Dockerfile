FROM golang:latest

WORKDIR .

COPY . .

RUN go build -o myapp .

CMD ["./myapp"]