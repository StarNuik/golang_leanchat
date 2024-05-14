FROM golang:alpine as builder

RUN apk add protoc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

COPY go.mod go.sum /go/src/leanchat/
WORKDIR /go/src/leanchat
RUN go mod download

COPY . .

RUN ./build_protobuf.sh
RUN go build -o build/bin ./main.go


FROM alpine
RUN apk add --no-cache libc6-compat
COPY --from=builder /go/src/leanchat/build/bin /usr/bin/leanchat
EXPOSE 3000

ENTRYPOINT /usr/bin/leanchat serve 3000

