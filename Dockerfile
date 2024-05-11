FROM golang:latest

WORKDIR .

COPY . .

RUN go build -o myapp .

COPY ./docker/entrypoint.sh /var/www/html/docker/entrypoint.sh
RUN chmod +x /var/www/html/docker/entrypoint.sh

RUN apt-get update && apt-get install -y netcat-openbsd

CMD ["./docker/entrypoint.sh"]