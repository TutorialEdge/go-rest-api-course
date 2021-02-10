# Our builder image used to build the Go binary
FROM golang:1.15.7-alpine as builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go clean --modcache
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/server/main.go

# Our production image used to run our app
FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache git make musl-dev go
COPY --from=builder /app/main .
# Configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin
EXPOSE 8080
CMD ["./main"]