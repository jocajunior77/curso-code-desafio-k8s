FROM golang:alpine AS builder

WORKDIR /src/app
RUN apk add git
COPY app.go .

RUN go get -d -v \
 && go build app.go \
 && chmod +x app

FROM scratch

COPY --from=builder /src/app/app /bin/app

ENTRYPOINT ["/bin/app"]