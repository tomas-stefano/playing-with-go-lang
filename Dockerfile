FROM golang:1.14-alpine AS build

WORKDIR /src/
COPY main.go go.* /src/
RUN CGO_ENABLED=0 go build -o /bin/service-api

FROM scratch
COPY --from=build /bin/service-api /bin/service-api
ENTRYPOINT ["/bin/service-api"]
