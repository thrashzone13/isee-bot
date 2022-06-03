# build stage
FROM golang:1.18-alpine AS builder
RUN apk add --no-cache gcc libc-dev git
ADD . /src/
RUN cd /src && CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o isee-bot

# final stage
FROM alpine:3
COPY --from=builder /src/isee-bot /isee-bot
ENTRYPOINT ["./isee-bot"]