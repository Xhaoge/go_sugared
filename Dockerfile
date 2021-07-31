FROM golang:latest

WORKDIR $GOPATH/src/github.com/EDDYCJY/go_sugared
COPY . $GOPATH/src/github.com/EDDYCJY/go_sugared
RUN go build .

EXPOSE 8090
ENTRYPOINT ["./go_sugared"]