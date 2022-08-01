# Build
FROM golang:alpine AS build

RUN apk add --no-cache -U build-base git make ffmpeg-dev

RUN mkdir -p /src

WORKDIR /src

# Copy Makefile
COPY Makefile ./

# Install deps
RUN make deps

# Copy go.mod and go.sum and install and cache dependencies
COPY go.mod .
COPY go.sum .

# Copy static assets
COPY ./static/* ./static/

# Copy templates

COPY ./templates/* ./templates/

# Copy sources
COPY *.go ./
COPY ./app/*.go ./app/
COPY ./importers/*.go ./importers/
COPY ./media/*.go ./media/
COPY ./utils/*.go ./utils/

# Version/Commit (there there is no .git in Docker build context)
# NOTE: This is fairly low down in the Dockerfile instructions so
#       we don't break the Docker build cache just be changing
#       unrelated files that actually haven't changed but caused the
#       COMMIT value to change.
ARG VERSION="0.0.0"
ARG COMMIT="HEAD"

# Build server binary
RUN make server VERSION=$VERSION COMMIT=$COMMIT

# Runtime
FROM alpine:latest

RUN apk --no-cache -U add su-exec shadow ca-certificates tzdata ffmpeg

ENV PUID=1000
ENV PGID=1000

RUN addgroup -g "${PGID}" tube && \
    adduser -D -H -G tube -h /var/empty -u "${PUID}" tube && \
    mkdir -p /data && chown -R tube:tube /data

VOLUME /data

WORKDIR /

# force cgo resolver
ENV GODEBUG=netdns=cgo

COPY --from=build /src/tube /usr/local/bin/tube

COPY .dockerfiles/entrypoint.sh /init
COPY .dockerfiles/config.json /

ENTRYPOINT ["/init"]
CMD ["tube"]
