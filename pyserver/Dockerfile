FROM ubuntu:14.04

RUN apt-get update

RUN  apt-get install -y python

COPY . /app/

WORKDIR /app/

CMD ["./run_server.sh"]