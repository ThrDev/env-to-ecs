FROM golang:1.10-alpine AS builder

RUN apk add --no-cache curl git
RUN curl https://glide.sh/get | sh

WORKDIR /go

# install deps
COPY ./src/glide.lock src/
COPY ./src/glide.yaml src/
RUN cd src && glide install

# Copy src
COPY ./src/7factor.io src/7factor.io

# build binary
RUN go install 7factor.io/...

FROM scratch

COPY --from=builder /go/bin/cmd /go/bin/cmd

# execute binary
ENTRYPOINT ["/go/bin/cmd"]