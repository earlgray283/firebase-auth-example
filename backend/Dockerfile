FROM golang:1.18 AS build

WORKDIR /work

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN \
    GOOS=linux \
    go build -o backend-app

FROM ubuntu:22.04

WORKDIR /work

COPY --from=build /work/backend-app .
COPY --from=build /work/credentials.json .

ENTRYPOINT [ "/work/backend-app" ]