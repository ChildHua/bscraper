FROM golang:1.14.7-alpine

WORKDIR /web/app

COPY . .

RUN go env -w GO111MODULE=on \
	&& go env -w  GOPROXY=https://goproxy.cn,direct

CMD go run main.go