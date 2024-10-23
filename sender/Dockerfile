# syntax=docker/dockerfile:1
FROM docker.io/golang:1.17-alpine
WORKDIR /app
COPY . ./
RUN go build -o /sender

FROM docker.io/alpine:latest
COPY --from=0 /sender /sender
CMD ["/sender"]
