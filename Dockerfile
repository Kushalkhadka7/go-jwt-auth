FROM golang:alpine AS builder

# Enable go modules.
ENV GO111MODULE=on

WORKDIR /app

COPY ["go.mod", "go.sum", "./"]
RUN go mod download

COPY . ./

RUN go build -v . .

# Development stage.
FROM builder AS dev

COPY --from=builder /app .

# Get Reflex for live reload in dev env.
RUN go get github.com/cespare/reflex

EXPOSE 8080
EXPOSE 6000
EXPOSE 5000
CMD ["Make","run-dev"]

# Testing stage.
FROM builder AS test

COPY --from=builder /app .

CMD ["cd cmd", "go","test -v"]

# Production stage.
FROM builder AS prod

COPY --from=builder /app/cmd .

EXPOSE 8080
EXPOSE 6000
EXPOSE 5000
CMD ["./cmd"]
