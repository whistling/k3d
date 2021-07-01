FROM golang:latest
ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/my-web
COPY . $GOPATH/src/my-web
#RUN go build -o app .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
#RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY --from=0 /go/src/my-web/app .
EXPOSE 8080
ENTRYPOINT ["./app"]