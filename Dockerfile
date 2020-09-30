FROM golang:1.14

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/zzlpeter/dawn-go
COPY . $GOPATH/src/github.com/zzlpeter/dawn-go
RUN go mod init & go mod vendor & swag init
RUN go build .

EXPOSE 8001
CMD ["./dawn-go"]
