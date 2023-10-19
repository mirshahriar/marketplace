ARG DOCKER_BUILD_IMAGE=golang:1.20.4
ARG DOCKER_BASE_IMAGE=alpine:3.18

FROM ${DOCKER_BUILD_IMAGE} AS build

RUN apt-get update -yq

WORKDIR /github.com/mirshahriar/marketplace

COPY ./ ./

RUN make build

FROM ${DOCKER_BASE_IMAGE} AS final

RUN apk update && apk add ca-certificates
RUN apk add --no-cache tzdata

COPY --from=build /github.com/mirshahriar/marketplace/binary/github.com/mirshahriar/marketplace /bin/github.com/mirshahriar/marketplace

WORKDIR /github.com/mirshahriar/marketplace

EXPOSE 8080
ENTRYPOINT ["run"]
