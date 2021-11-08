FROM golang:latest as build

WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .
RUN go build -o main .


FROM golang:latest as prod
WORKDIR /app
COPY --from=build /app/main .
CMD [ "/app/main" ]