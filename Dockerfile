FROM golang:alpine as builder

# install binaries
RUN apk add protoc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN go install github.com/jackc/tern/v2@latest

# pull dep-s
COPY go.mod go.sum /go/src/leanchat/
WORKDIR /go/src/leanchat
RUN go mod download

# pull src-s
COPY . .

# build go
RUN ./build_protobuf.sh
RUN go build -o build/bin ./main.go

# final image
FROM alpine
RUN apk add --no-cache libc6-compat
COPY --from=builder /go/src/leanchat/build/bin /usr/bin/leanchat
COPY --from=builder /go/bin/tern /usr/bin/tern
COPY --from=builder /go/src/leanchat/sql_migrations/*.sql /usr/bin/sql_migrations/
COPY --from=builder /go/src/leanchat/entrypoint.sh /usr/bin/entrypoint.sh

# pre-flight
EXPOSE 3000
ENTRYPOINT /usr/bin/entrypoint.sh

