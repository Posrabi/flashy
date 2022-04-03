FROM golang:1.17.6

RUN export GOBIN=$GOPATH/bin

RUN apt update && apt install -y protobuf-compiler python3 python3-pip

RUN BIN="/usr/local/bin" && \
  VERSION="1.3.1" && \
  BINARY_NAME="buf" && \
  curl -sSL \
  "https://github.com/bufbuild/buf/releases/download/v${VERSION}/${BINARY_NAME}-$(uname -s)-$(uname -m)" \
        -o "${BIN}/${BINARY_NAME}" && \
    chmod +x "${BIN}/${BINARY_NAME}"  

WORKDIR /go/src/github.com/Posrabi/flashy
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install \
  github.com/golang/protobuf/protoc-gen-go \
  google.golang.org/grpc/cmd/protoc-gen-go-grpc

RUN ./dockerfile-build.sh
