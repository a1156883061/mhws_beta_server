FROM golang:1.23.3-alpine
WORKDIR /app
COPY ./ ./
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build main.go
CMD go run . 127.0.0.1