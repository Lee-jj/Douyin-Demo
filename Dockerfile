FROM golang:1.19-alpine
WORKDIR /apps
COPY . .

RUN go env -w GOPROXY=https://goproxy.cn,direct \
        && go mod tidy \
        && go build -o Douyin-Demo main.go \
        && sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
        && apk update --no-cache \
        && apk add ffmpeg  

EXPOSE 8080

CMD ["/apps/Douyin-Demo"]