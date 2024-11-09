FROM golang:1.23.3-alpine
WORKDIR /app
COPY ./ ./
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build main.go
CMD run . 0.0.0.0