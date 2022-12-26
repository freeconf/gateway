#########################
## G O  B U I L D
#########################
FROM golang:1.18 as build

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN apt-get update && apt-get install -y make

WORKDIR /app
COPY . .

RUN make fc-gateway

#########################
## T A R G E T
#########################
FROM alpine:latest

WORKDIR /app

COPY ./yang /app/yang
COPY ./cmd/fc-gateway/startup.json /app/startup.json
COPY --from=build /app/fc-gateway .

EXPOSE 8080
ENTRYPOINT ["/app/fc-gateway", "ypath", "/app/ypath", "-startup", "startup.json"]
