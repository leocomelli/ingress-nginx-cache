FROM golang:1.21.3 as build

WORKDIR /app
COPY . .

ENV CGO_ENABLED 0
RUN go build -o ingress-nginx-cache main.go

FROM scratch

COPY --from=build /app/ingress-nginx-cache /
EXPOSE 8080

ENTRYPOINT [ "/ingress-nginx-cache" ]
