FROM golang:1.16.5-alpine3.14 AS build
WORKDIR src/
COPY . .
RUN go build -o /out/main .
FROM scratch AS bin
COPY --from=build /out/main /
