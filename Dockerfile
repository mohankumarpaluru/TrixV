# Build
FROM prologic/go-builder:latest AS build

# Runtime
FROM golang:alpine

RUN apk --no-cache -U add git build-base ffmpeg ffmpeg-dev

RUN go install github.com/mutschler/mt@latest

COPY --from=build /src/tube /tube

ENTRYPOINT ["/tube"]
CMD [""]
