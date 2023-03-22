FROM golang:1.20.2-bullseye as builder
WORKDIR /hello
COPY . ./
RUN go build -o hellodocker

FROM ubuntu:20.04
WORKDIR /hello
COPY --from=builder /hello/config/ config/
COPY --from=builder /hello/hellodocker ./
CMD ["./hellodocker"]
EXPOSE 3000
