FROM golang:latest as build
WORKDIR /app
COPY . .
RUN go build main.go

FROM azer0s/qpmd:latest
COPY --from=build /app/main /app/main