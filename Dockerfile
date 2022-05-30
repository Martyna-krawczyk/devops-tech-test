FROM golang:1.16-alpine as base
WORKDIR /app
COPY . .

ARG version
ARG commit
ENV COMMIT ${commit}
ENV VERSION ${version}

ENV CGO_ENABLED 0
RUN go build -ldflags "-X main.version=$VERSION -X main.gitCommit=$COMMIT" -o ./_build

FROM alpine:latest as runtime
WORKDIR /app
COPY --from=base /app/_build ./
HEALTHCHECK --interval=5s --timeout=3s \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080 || exit 1
CMD ["./_build"]