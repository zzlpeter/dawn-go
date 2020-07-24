FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/zzlpeter/dawn-go
COPY . $GOPATH/src/github.com/zzlpeter/dawn-go
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./go-gin-example"]
