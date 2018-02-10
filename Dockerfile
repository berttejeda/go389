FROM golang:1.9 as builder

WORKDIR /go/src/github.com/berttejeda/go389

# Prerequisites
RUN apt-get update && \
    apt-get install \
      libpam0g-dev

COPY . .

RUN git config --global http.sslVerify false

ENV GOROOT /usr/local/go
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

RUN go get -d -v -insecure

RUN go get -insecure github.com/berttejeda/go389

RUN go get -insecure gopkg.in/yaml.v1
 
RUN go get -insecure github.com/msteinert/pam

RUN CGO_ENABLED=0 GOOS=linux go build -tags "yaml	 netgo" -a -installsuffix cgo

FROM alpine:edge 

RUN apk add --no-cache bash

WORKDIR /root/

COPY --from=builder /go/src/github.com/berttejeda/go389/scripts/entrypoint.sh .

RUN chmod +x entrypoint.sh && mkdir conf

COPY --from=builder /go/src/github.com/berttejeda/go389/conf/app.yml ./conf

COPY --from=builder /go/src/github.com/berttejeda/go389/conf/db.yml ./conf

COPY --from=builder /go/src/github.com/berttejeda/go389/go389 .

CMD ["./entrypoint.sh"]