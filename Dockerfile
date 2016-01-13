FROM golang:latest
COPY . /go/src/dig
WORKDIR /go/src/dig/
RUN go build
EXPOSE 9001
CMD /go/src/dig/dig
