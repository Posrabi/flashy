ARG BASE_IMAGE
FROM $BASE_IMAGE as build-env
FROM alpine:3.12.0 as final
RUN apk --no-cache add \
  ca-certificates build-base libffi-dev openssl-dev \
  python3 python3-dev py3-pip
WORKDIR /root/
COPY --from=build-env /go/src/github.com/Posrabi/flashy/build/users .
CMD ["./users", "server"]
