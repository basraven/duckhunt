from golang:alpine
RUN apk add git build-base

#PROD:
# RUN go get -u github.com/basraven/duckhunt && go build github.com/basraven/duckhunt
# CMD duckhunt

#DEV:
RUN go get -u github.com/go-telegram-bot-api/telegram-bot-api && \
    go get -u github.com/lib/pq

# DEBUGGING:
RUN go get -u github.com/go-delve/delve/cmd/dlv

COPY ./entrypoint.sh /
ENTRYPOINT /entrypoint.sh