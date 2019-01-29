From golang:alpine as builder

MAINTAINER Shrivatsa Upadhye "ishrivatsa@gmail.com"

# Future proofing by installing ca-cert for https support
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates
COPY . $GOPATH/src/github.com/vmwarecloudadvocacy/user
WORKDIR $GOPATH/src/github.com/vmwarecloudadvocacy/user
ENV GO111MODULE=on
ENV CGO_ENABLED=0
RUN go build -o user .

FROM alpine
RUN mkdir app
#Copy the executable uilt from the previous image
COPY --from=builder /go/src/github.com/vmwarecloudadvocacy/user/user /app
WORKDIR /app
EXPOSE 80
CMD ["./user"]