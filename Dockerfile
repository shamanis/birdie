## build
FROM golang:1.19-alpine AS build
WORKDIR /opt/birdie
COPY . .

RUN go mod download
RUN go build -ldflags="-w -s" -o birdie cmd/birdie/main.go

# deploy
FROM scratch
WORKDIR /opt/birdie
COPY --from=build /opt/birdie/birdie /birdie
ENTRYPOINT ["/birdie"]
