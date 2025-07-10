ARG GO_VERSION=1.24.2
ARG MODULE_NAME

FROM golang:${GO_VERSION}-alpine AS build

ARG MODULE_NAME

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/${MODULE_NAME}/main.go ./main.go
COPY pkg/ ./pkg/

RUN echo "Building cmd/${MODULE_NAME}/main.go"

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /app/app ./main.go

FROM gcr.io/distroless/static:nonroot

WORKDIR /

COPY --from=build /app/app .

EXPOSE 8080

ENTRYPOINT ["/app"]
