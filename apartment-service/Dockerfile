FROM golang:1.15-alpine3.13
WORKDIR apartment-service
RUN apk update \
    && apk add socat \
    && apk add build-base
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o apartment-service
EXPOSE 8080
CMD ["./apartment-service"]