#build stage
FROM golang:1.22.4
RUN go install -v github.com/go-delve/delve/cmd/dlv@latest && go install -v golang.org/x/tools/gopls@latest

ENV ENV="/etc/profile"
COPY alias /tmp/
RUN cat /tmp/alias >> /etc/profile