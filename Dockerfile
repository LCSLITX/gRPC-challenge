FROM golang:alpine
RUN apk add --no-cache git
RUN apk add --update make
WORKDIR $GOPATH/src/github.com/lucassauro/klever-challenge/
COPY . .
RUN go mod tidy
RUN go build -o ./bin/ ./src/server
ENTRYPOINT [ "./bin/server" ]