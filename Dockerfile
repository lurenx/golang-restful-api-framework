FROM golang:alpine

RUN apk add build-base

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .

RUN CGO_ENABLED=1 go build -o app cmd/main.go

WORKDIR /dist

COPY config/setup.yml .

RUN cp /build/app .

EXPOSE 8888

ENTRYPOINT ["/dist/app","-f","/dist/setup.yml"]
