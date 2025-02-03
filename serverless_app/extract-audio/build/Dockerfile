FROM golang:1.22.0-alpine AS build_image

WORKDIR /tmp/build-app

RUN apk add git

COPY . .

RUN go mod tidy && go build -o app .

FROM alpine:3.16.2

WORKDIR /app/
COPY --from=build_image /tmp/build-app/app ./
RUN apk add ffmpeg && \
    mkdir -p /tmp

CMD ["./app"]
VOLUME ["/tmp"]
