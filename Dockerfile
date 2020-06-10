FROM golang:latest AS builder
LABEL maintainer="oeid@cognitivescale.com"

WORKDIR /app/

COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o exe .

FROM alpine:latest
WORKDIR /
COPY --from=builder /app/exe /app/exe
ENTRYPOINT [ "/app/exe" ]
