FROM golang:latest as builder
WORKDIR /go/src/awesomeProject/cicd
COPY . .
RUN go build -a -o helloworld Helloworld.go

FROM scratch
COPY --from=builder /go/src/awesomeProject/cicd/helloworld /helloworld
ENTRYPOINT ["/helloworld"]