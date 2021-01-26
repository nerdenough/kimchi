FROM golang:1.15
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -v -o app ./main.go

FROM alpine
WORKDIR /app
COPY --from=0 /go/src/app .
ENTRYPOINT ["./app"]
CMD ["run", "--logtostderr"]
