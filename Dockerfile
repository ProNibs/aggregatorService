# Using Alpine image as it's small compared to the rest
FROM golang:1.17.3-alpine3.14 as build

WORKDIR /app
# Download dependencies
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
# Copy rest of source code
# If dependencies don't change, docker build will be quicker
# as the above layers will be cached.
COPY . .
RUN go build -o main .


FROM golang:1.17.3-alpine3.14 as prod
# This is the default, but just to be explicit
ENV HTTP_PORT=7777
WORKDIR /app
COPY --from=build /app/main .
EXPOSE 7777
CMD [ "/app/main" ]