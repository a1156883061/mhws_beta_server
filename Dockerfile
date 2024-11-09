FROM golang:1.23.3-alpine
WORKDIR /app
RUN go build main.go
CMD run . 0.0.0.0