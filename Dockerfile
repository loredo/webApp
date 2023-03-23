FROM golang:1.20.2-bullseye as builder
WORKDIR /webapp
COPY . ./
RUN go build -o webapp

FROM ubuntu:20.04
WORKDIR /webapp
COPY --from=builder /webapp/config/ config/
COPY --from=builder /webapp/webapp ./
CMD ["./webapp"]
EXPOSE 3000
