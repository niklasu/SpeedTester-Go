FROM golang:1.20 AS build-stage

WORKDIR /app

COPY go.mod ./
COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /speedtester

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /speedtester /speedtester

USER nonroot:nonroot

ENTRYPOINT ["/speedtester"]