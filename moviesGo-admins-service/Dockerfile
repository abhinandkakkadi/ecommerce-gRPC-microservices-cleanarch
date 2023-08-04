
FROM golang:1.19-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./out/dist ./cmd


FROM alpine:3.14

COPY --from=build /app/out/dist /app/dist

WORKDIR /app

CMD ["/app/dist"]