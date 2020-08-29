FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git 

RUN mkdir /gcounter

WORKDIR /gcounter

COPY . .

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -installsuffix cgo -o /go/bin/gcounter


FROM scratch

COPY --from=builder /go/bin/gcounter /go/bin/gcounter

ENTRYPOINT ["/go/bin/gcounter"]

EXPOSE 8080