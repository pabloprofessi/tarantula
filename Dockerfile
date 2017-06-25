FROM golang:1.7-alpine 

RUN apk update
RUN apk upgrade
RUN apk add bash curl git wget

# Set environment variables.
ENV PATH $PATH:/app/bin
ENV GOPATH /app

RUN mkdir -p /app/src/github.com/tarantula
COPY . /app/src/github.com/tarantula

WORKDIR /app/src/github.com/tarantula

RUN go install

CMD ["tarantula"]
