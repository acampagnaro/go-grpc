FROM golang:1.14.3-alpine AS build
WORKDIR /src
COPY . .
RUN go build -o /src .
FROM scratch AS bin
COPY --from=build /src /