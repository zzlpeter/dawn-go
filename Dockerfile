FROM golang:1.14

ARG app_env
ENV APP_ENV $app_env

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/zzlpeter/dawn-go
COPY . $GOPATH/src/github.com/zzlpeter/dawn-go
RUN go get -u github.com/swaggo/swag/cmd/swag

RUN go build .

EXPOSE 8001
CMD ["./dawn-go"]
