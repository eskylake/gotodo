FROM golang:1.21 as base

# Dev
FROM base as dev

WORKDIR /app/api
# RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
RUN go install github.com/cosmtrek/air@latest
ENV CGO_ENABLED=0

COPY go*.* ./
RUN go mod download && go mod tidy && go mod verify && go mod vendor

COPY . .

CMD ["tail", "-f", "/dev/null"]

# Prod
FROM base as built

WORKDIR /go/app/api
COPY . .

ENV CGO_ENABLED=0

RUN go get -d -v ./...
RUN go build -o /tmp/api-server ./*.go

FROM busybox

COPY --from=built /tmp/api-server /usr/bin/api-server
CMD ["api-server", "start"]