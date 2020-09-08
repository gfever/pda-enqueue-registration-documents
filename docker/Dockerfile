FROM golang:1.15.1-buster

# Set go bin which doesn't appear to be set already.
ENV GOBIN /go/bin
ENV GOROOT /usr/local/go

RUN apt-get update && apt-get install -y unzip

ARG PROTOC_ZIP=protoc-3.9.1-linux-x86_64.zip
RUN curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.9.1/$PROTOC_ZIP
RUN unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
RUN unzip -o $PROTOC_ZIP -d /usr/local 'include/*'
RUN rm -f $PROTOC_ZIP

WORKDIR /tmp
RUN go get -u google.golang.org/protobuf/cmd/protoc-gen-go && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go

RUN mkdir /go/src/pda-enqueue-registration-documents
COPY . /go/src/pda-enqueue-registration-documents
WORKDIR /go/src/pda-enqueue-registration-documents

RUN go get -u github.com/gin-gonic/gin

RUN go mod vendor
RUN chmod 777 /go/src/pda-enqueue-registration-documents/run.sh

ENTRYPOINT ["sh", "/go/src/pda-enqueue-registration-documents/run.sh"]
