FROM golang:1.18-alpine
WORKDIR /zk
COPY bin/zk-socket-server .

CMD ["/zk/zk-socket-server", "-c", "/zk/config/config.yaml"]
