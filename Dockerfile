FROM golang:1.22-alpine AS builder

RUN apk add --no-cache git gcc libc-dev

RUN go install github.com/gobuffalo/cli/cmd/buffalo@v0.18.14

ENV GO111MODULE=on
ENV GOPROXY https://proxy.golang.org

WORKDIR /src/endpoints-under-test

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN /go/bin/buffalo build --static -o /bin/app

FROM alpine
RUN apk add --no-cache bash
RUN apk add --no-cache ca-certificates

WORKDIR /bin/

COPY --from=builder /bin/app .

ENV GO_ENV=production
ENV PORT=8001
ENV ADDR=0.0.0.0

EXPOSE 8001

CMD exec /bin/app
