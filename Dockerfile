From golang:alpine as builder

MAINTAINER Shrivatsa Upadhye "ishrivatsa@gmail.com"

# Future proofing by installing ca-cert for https support
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates
COPY . $GOPATH/src/github.com/vmwarecloudadvocacy/user
WORKDIR $GOPATH/src/github.com/vmwarecloudadvocacy/user
ENV GO111MODULE=on
ENV CGO_ENABLED=0
RUN go build -o bin/user ./cmd/users

FROM bitnami/minideb:stretch
RUN apt-get update && apt-get install -y gnupg2
RUN apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 4B7C549A058F8B6B
RUN echo "deb http://repo.mongodb.org/apt/debian stretch/mongodb-org/4.2 main" | tee /etc/apt/sources.list.d/mongodb-org-4.2.list
RUN apt update && apt -y upgrade
RUN apt-get install -y mongodb-org
# needed for redis-cli ; the server is not used
RUN install_packages redis-server
#RUN install_packages mongodb-clients
RUN mkdir app
#Copy the executable from the previous image
COPY --from=builder /go/src/github.com/vmwarecloudadvocacy/user/bin/user /app
COPY entrypoint/docker-entrypoint.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/docker-entrypoint.sh
RUN ln -s usr/local/bin/docker-entrypoint.sh /app
WORKDIR /app
EXPOSE 80
EXPOSE 8081
ENTRYPOINT ["docker-entrypoint.sh"]
