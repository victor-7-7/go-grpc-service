FROM golang:1.22 as builder
RUN mkdir /app
ADD . /app
WORKDIR /app
#RUN CGO_ENABLED=0 goos=windows go build -o app cmd/server/main.go
RUN CGO_ENABLED=0 goos=darwin go build -o app cmd/server/main.go

FROM alpine:latest as production
COPY --from=builder /app .
CMD ["./app"]