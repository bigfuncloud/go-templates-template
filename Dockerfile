FROM debian:buster-slim

RUN apt-get update && apt-get install -y ca-certificates

WORKDIR /app
COPY out/ /app/
CMD ["./app"]
