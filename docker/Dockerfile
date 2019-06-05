from golang:alpine
RUN apk add git
# RUN go get -u github.com/basraven/duckhunt && go build github.com/basraven/duckhunt
# CMD duckhunt

#TODO: Remove
RUN go get -u github.com/go-telegram-bot-api/telegram-bot-api && \
    go get -u github.com/lib/pq
CMD ["sleep", "1000000"]