FROM golang:1.19-alpine AS build
WORKDIR /app/
COPY . .
RUN go install
RUN go build -o executable
EXPOSE 8080

FROM alpine:3.13.2
RUN apk add --no-cache tzdata
COPY --from=build /executable /executable
ENV PORT=80
EXPOSE 80
ENTRYPOINT ["/executable"]