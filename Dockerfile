FROM ubuntu:14.04

RUN apt-get update

RUN  apt-get install -y wget git

RUN wget https://storage.googleapis.com/golang/go1.7.4.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.7.4.linux-amd64.tar.gz

# Set environment variables.
ENV PATH $PATH:/usr/local/go/bin
ENV GOPATH /app/tarantula
ENV PATH $PATH:/app/tarantula/bin

COPY . /app/tarantula

WORKDIR /app/tarantula/src/github.com/pprofessi/server

RUN go get
RUN go install

CMD ["server"]
