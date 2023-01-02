# Build application
# https://github.com/GoogleContainerTools/distroless/
FROM golang:1.19 as build

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go vet -v
RUN go test -v

RUN CGO_ENABLED=0 go build -o /go/bin/app

# Copy application into base image
FROM gcr.io/distroless/static-debian11

COPY --from=build /go/bin/app /

ENTRYPOINT ["/app"]