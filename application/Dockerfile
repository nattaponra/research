FROM golang:1.11 AS builder
RUN ls
# Magic line, notice in use that the lib name is different!
RUN apt-get update && apt-get install -y gcc-aarch64-linux-gnu
# Add your app and do what you need to for dependencies
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=1 CC=aarch64-linux-gnu-gcc GOOS=linux GOARCH=arm64 go build -o app .

# Final stage - pick any old arm64 image you want
 
FROM multiarch/ubuntu-core:arm64-bionic
RUN apt-get update & apt-get install dnsutils
WORKDIR /root/

COPY --from=builder /app .
CMD ["./app"]