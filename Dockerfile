FROM golang:1.18-alpine AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /main

FROM build-stage AS run-test-stage
RUN go test -v ./...

FROM alpine AS build-release-stage
WORKDIR /
COPY --from=build-stage /main /main
EXPOSE 8080
ENTRYPOINT ["/main"]