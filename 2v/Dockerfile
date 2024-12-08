FROM golang:1.23-alpine AS builder


WORKDIR /app


RUN apk --no-cache add bash git make gcc gettext musl-dev

COPY ["go.mod","go.sum","./"]

RUN go mod download

COPY . . 
#build of events
RUN go build -o ./events ./cmd/events/main.go
#build of notifications
RUN go build -o ./notifications ./cmd/notifications/main.go
#build of players
RUN go build -o ./players ./cmd/players/main.go
#build of rest api
RUN go build -o ./rest ./cmd/httpHandler/main.go


FROM alpine

WORKDIR /

COPY --from=builder /app/events /events

COPY --from=builder /app/notifications /notifications

COPY --from=builder /app/players /players

COPY --from=builder /app/rest /rest

CMD ["/rest"]

