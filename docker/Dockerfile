FROM alpine:3.6
MAINTAINER Matt Snoby <matt.snoby@icloud.com>



RUN     apk --update add --no-cache ca-certificates bash vim openssl
COPY    webhook     /usr/local/bin/webhook
COPY    spark-pivot /usr/local/bin/spark-pivot
RUN     chmod 755 /usr/local/bin/spark-pivot && chmod 755 /usr/local/bin/webhook

#
EXPOSE      9000

COPY hooks.json /etc/webhook/hooks.json

VOLUME "/etc/webhook/"

ENTRYPOINT  ["/usr/local/bin/webhook"]
CMD ["-verbose", "-hooks=/etc/webhook/hooks.json", "-hotreload"]
