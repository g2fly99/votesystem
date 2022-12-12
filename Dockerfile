FROM golang:1.18.3

#CN need proxy
#ENV GOPROXY=https://goproxy.io

RUN go env -w GO111MODULE=on
RUN go install github.com/beego/bee/v2@latest

ENV PATH $PATH:$GOPATH/bin
WORKDIR $GOPATH/src/

RUN git clone https://github.com/g2fly99/votesystem.git

WORKDIR $GOPATH/src/votesystem

# Expose the application on port 8080
EXPOSE 8080
CMD ["bee","run"]
