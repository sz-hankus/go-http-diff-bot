FROM golang:alpine as builder

RUN mkdir /workdir
WORKDIR /workdir
COPY . .

RUN go mod download
# TODO: dodać GOARCH jeśli to będzie na raspi
RUN CGO_ENABLED=0 GOOS=linux go build .

FROM alpine:latest

COPY --from=builder /workdir/http_diff_bot ./
ENTRYPOINT [ "./http_diff_bot" ] 