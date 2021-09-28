##
## Build
##

FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY *.go ./

RUN go build -o /apollo-lyrics

##
## Deploy
##

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /apollo-lyrics /apollo-lyrics

EXPOSE 8080

USER nonroot:nonroot

HEALTHCHECK CMD curl --fail http://localhost:8080/ping || exit 1

ENTRYPOINT ["/apollo-lyrics"]
