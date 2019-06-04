FROM golang:1.12.5-alpine AS gobuild-base
RUN apk add --no-cache \
    git

FROM gobuild-base AS base
WORKDIR /go/src/github.com/ehotinger/scratchfile
COPY . .
RUN go build ./... && mv scratchfile /usr/bin/scratchfile

FROM scratch
COPY --from=base /usr/bin/scratchfile /usr/bin/scratchfile
ENTRYPOINT [ "scratchfile" ]
CMD [ ]
