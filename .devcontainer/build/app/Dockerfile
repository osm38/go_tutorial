#build stage
FROM golang:1.22.3-alpine3.19
RUN apk update && apk add --no-cache alpine-sdk build-base tzdata && go install -v github.com/go-delve/delve/cmd/dlv@latest && go install -v golang.org/x/tools/gopls@latest

ENV ENV="/etc/profile"
COPY alias /tmp/
RUN cat /tmp/alias >> /etc/profile