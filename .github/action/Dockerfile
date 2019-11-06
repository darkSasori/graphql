FROM golang:1.13 AS linter

RUN go get golang.org/x/lint/golint

FROM alpine:3.10.3

LABEL name="Golang Lint"
LABEL maintainer="Lineu Felipe"
LABEL version="0.0.1"

LABEL com.github.actions.name="Golang Lint"
LABEL com.github.actions.description="Run lint"
LABEL com.github.actions.icon="box"
LABEL com.github.actions.color="blue"

COPY --from=linter /go/bin/golint /usr/local/bin/golint
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
