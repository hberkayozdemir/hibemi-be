FROM alpine:3.14.2

WORKDIR /app

COPY .config .config
COPY dists/* .

ENTRYPOINT [ "./service-template" ]