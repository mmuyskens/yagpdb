FROM golang:stretch as builder

WORKDIR $GOPATH/src


# Uncomment during development
# COPY . github.com/mmuyskens/yagpdb

RUN CGO_ENABLED=0 GOOS=linux go get -v \
  github.com/mmuyskens/yagpdb/cmd/yagpdb

FROM alpine:latest

WORKDIR /app
VOLUME /app/soundboard \
  /app/cert
EXPOSE 80 443

# We need the X.509 certificates for client TLS to work.
RUN apk --no-cache add ca-certificates

# Add ffmpeg for soundboard support
RUN apk --no-cache add ffmpeg

# Handle templates for plugins automatically
COPY --from=builder /go/src/github.com/mmuyskens/yagpdb/*/assets/*.html templates/plugins/

COPY --from=builder /go/src/github.com/mmuyskens/yagpdb/cmd/yagpdb/templates templates/
COPY --from=builder /go/src/github.com/mmuyskens/yagpdb/cmd/yagpdb/posts posts/
COPY --from=builder /go/src/github.com/mmuyskens/yagpdb/cmd/yagpdb/static static/

COPY --from=builder /go/bin/yagpdb .

# add extra flags here when running YAGPDB
# Set "-exthttps=true" if using a TLS-enabled proxy such as
# jrcs/letsencrypt-nginx-proxy-companion
# Set "-https=false" do disable https
ENV extra_flags ""

# `exec` allows us to receive shutdown signals.
CMD exec /app/yagpdb -all -pa $extra_flags

