FROM golang:1.10-alpine3.8 as builder
ARG URL_BRANCH
RUN apk update && \
    apk add git && \
    apk add --update bash && \
    apk add openssh
    
RUN go get github.com/golang/dep/cmd/dep
WORKDIR $GOPATH/src/github.com/Ankr-network/dccn-cli/
COPY . $GOPATH/src/github.com/Ankr-network/dccn-cli/

RUN CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go build -a \
    -installsuffix cgo \
    -ldflags="-w -s -X github.com/Ankr-network/dccn-cli/commands.clientURL=${URL_BRANCH}" \
    -o /go/bin/akrctl \
    $GOPATH/src/github.com/Ankr-network/dccn-cli/cmd/akrctl/main.go

FROM scratch
COPY --from=builder /go/bin/akrctl /go/bin/akrctl
ENTRYPOINT ["/go/bin/akrctl"]
