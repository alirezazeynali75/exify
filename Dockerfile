FROM golang:1.22-bookworm AS installer

WORKDIR /app

COPY ./go.mod .
COPY ./go.sum .

RUN go mod download
COPY . .

FROM installer as builder

RUN make build-cli
RUN make build


FROM debian:bookworm-slim as final


RUN apt-get update && \
    apt-get install -y ca-certificates procps curl net-tools && \
    apt-get clean

COPY --from=builder /app/docker-entrypoint.sh ./entrypoint.sh
COPY --from=builder /app/bin/exify ./exify
COPY --from=builder /app/bin/exify-cli ./exify-cli

EXPOSE 3000
ENTRYPOINT [ "/entrypoint.sh" ]
CMD [ ]
