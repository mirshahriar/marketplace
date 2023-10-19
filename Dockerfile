ARG DOCKER_BUILD_IMAGE=golang:1.20.4
ARG DOCKER_BASE_IMAGE=alpine:3.18

FROM ${DOCKER_BUILD_IMAGE} AS build

RUN apt-get update -yq

WORKDIR /mirshahriar

COPY ./ ./

RUN make build

FROM ${DOCKER_BASE_IMAGE} AS final

RUN apk update && apk add ca-certificates

COPY --from=build /mirshahriar/binary/marketplace /bin/marketplace

WORKDIR /marketplace

EXPOSE 8080
ENTRYPOINT ["run"]
