FROM golang:latest as builder
WORKDIR /app
COPY ./src .
RUN GOOS=linux go build -ldflags="-w -s" -o artefact .

FROM debian:latest
WORKDIR /app
COPY --from=builder /app/artefact .
COPY ./data .
ENV TZ=America/Sao_Paulo
RUN apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/$TZ /etc/localtime && \
    echo $TZ > /etc/timezone
RUN apt-get update && apt-get install -y cron default-mysql-client

CMD ["/app/artefact"]