FROM alpine:3.7
LABEL maintainer="misskiwi@decentrandom.com"

RUN apk update && \
    apk upgrade && \
    apk --no-cache add curl jq file

VOLUME /randd
WORKDIR /randd
EXPOSE 26656 26657
ENTRYPOINT ["/usr/bin/wrapper.sh"]
CMD ["start"]
STOPSIGNAL SIGTERM

COPY wrapper.sh /usr/bin/wrapper.sh