FROM golang:1.6

ADD . /go/src/hume

RUN go install hume

ENTRYPOINT /go/bin/hume

EXPOSE 8000
