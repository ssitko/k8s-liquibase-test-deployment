# syntax=docker/dockerfile:1
FROM debian:latest

WORKDIR /app

COPY . .

CMD ["./kubernetes-gh-k8s"]

EXPOSE 8090