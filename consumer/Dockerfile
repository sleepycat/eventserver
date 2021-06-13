# COMPILE
FROM golang:alpine AS builder
ENV USER=consumer
ENV UID=10001
# always use go modules
ENV GO111MODULE=on
# build static standalone binary
ENV CGO_ENABLED=0

WORKDIR /app

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

COPY . .

RUN go build -o consumer

# BUILD
FROM scratch

ENV USER=consumer

WORKDIR /usr/local/bin
# Import the user and group files from step 1
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder --chown=${USER}:${USER} /app/consumer /usr/local/bin/consumer

USER ${USER}:${USER}

ENTRYPOINT ["consumer"]
