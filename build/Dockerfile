FROM golang:1.19 AS build

WORKDIR /app

COPY ./go.mod ./
COPY ./cmd/main.go ./

RUN go build -o /server


FROM golang:1.19-alpine 

WORKDIR /
COPY --from=build /server /server

EXPOSE 8000

USER nonroot:nonroot

ENTRYPOINT [ "/server" ]

