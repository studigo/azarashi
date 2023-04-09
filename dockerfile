# go をコンパイルする.
FROM golang:1.20 as builder
ARG CGO_ENABLED=0
COPY source/ /build/
WORKDIR /build
RUN apt update &&\
    apt install -y upx &&\
    go build -ldflags '-s -w' main.go &&\
    upx main

# シングルバイナリのコンテナを作成.
FROM scratch
EXPOSE 8080
EXPOSE 3306
COPY --from=builder /build/main /
CMD [ "/main" ]