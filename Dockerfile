FROM golang:1.9 as builder

WORKDIR /go/src/github.com/berttejeda/kernel164/go389

# Prerequisites
RUN apt-get update && \
    apt-get install \
      libpam0g-dev

COPY . .

RUN go get github.com/berttejeda/go389

RUN go get gopkg.in/yaml.v1
 
RUN go get github.com/msteinert/pam

RUN CGO_ENABLED=0 GOOS=linux go build -i -tags "yaml pam" -a -installsuffix cgo -o app .

FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /go/src/github.com/berttejeda/kernel164/go389/app .

CMD ["./app"]