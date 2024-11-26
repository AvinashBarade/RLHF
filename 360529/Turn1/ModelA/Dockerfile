FROM golang:1.20 AS build
WORKDIR /app
COPY go.mod go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /app/main

FROM scratch
COPY --from=build /app/main /app/main
EXPOSE 8080
CMD ["/app/main"]