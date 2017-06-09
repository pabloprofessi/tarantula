FROM golang:1.7-alpine 

RUN apk update
RUN apk upgrade
RUN apk add bash curl git wget

# Set environment variables.
ENV PATH $PATH:/usr/local/go/bin:/app/tarantula/bin
ENV GOPATH /app/tarantula

COPY . /app/tarantula

WORKDIR /app/tarantula/src/github.com/pprofessi/server

RUN go get
RUN go install

CMD ["server"]
