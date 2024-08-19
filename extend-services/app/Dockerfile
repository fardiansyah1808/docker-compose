FROM golang:1.23-alpine

ENV APP_PORT=8080
ENV MODE=local

RUN apk --no-cache add curl

RUN mkdir app
COPY main.go app

EXPOSE ${APP_PORT}

# HEALTHCHECK --interval=5s --start-period=5s CMD curl -f http://localhost:8080/health
CMD go run app/main.go