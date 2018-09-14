FROM alpine:3.8

ADD kube-netcheck /usr/bin/kube-netcheck
ENTRYPOINT ["/usr/bin/kube-netcheck"]
