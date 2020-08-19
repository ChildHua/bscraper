FROM golang:1.14.7-alpine

WORKDIR /web/app

RUN go env -w GO111MODULE=on \
	&& go env -w  GOPROXY=https://goproxy.cn,direct
