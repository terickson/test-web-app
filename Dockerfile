FROM golang:1.20 AS builder
RUN mkdir -p /build
WORKDIR /build
COPY . /build
RUN CGO_ENABLED=0 go build
RUN go test
RUN mkdir /new_tmp
RUN chmod 777 /new_tmp

FROM scratch as deploy
COPY --from=builder /new_tmp /tmp
COPY --from=builder /build/test-web-app /test-web-app
COPY --from=builder /build/web /web
ENTRYPOINT ["/test-web-app"]
