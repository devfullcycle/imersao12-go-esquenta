FROM golang:latest

WORKDIR /go/app

RUN apt-get update && apt-get install -y librdkafka-dev

CMD ["tail", "-f", "/dev/null"]