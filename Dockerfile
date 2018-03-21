FROM alpine:3.7

ADD kube-netcheck /usr/bin/kube-netcheck
ENTRYPOINT ["/usr/bin/kube-netcheck"]
