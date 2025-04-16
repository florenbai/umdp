FROM harbor.maas.local/library/golang:1.20 as builder
WORKDIR /data
COPY ./ ./
RUN go env -w GO111MODULE="on" && go env -w GOPROXY="https://goproxy.cn" && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /data/umdpserver ./app/manage/cmd/main.go

FROM harbor.maas.local/library/ubuntu:22.10
WORKDIR /data
COPY --from=builder /data/umdpserver /data/umdpserver
RUN ls -al
RUN mkdir -p /data/log
ENV INST_NO=0
ENV BEGIN_PORT=8000
ENV SERVICE=umdp
ENV LOCATION=CN
ENV LANG=C.UTF-8
ENTRYPOINT exec /data/umdpserver -c /data/config/config.yaml