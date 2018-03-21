# kube-netcheck

Tiny script to check Kubernetes network health.

Kube-netcheck checks:
- dns by resolving service name
- pods from one node can connect to pods on all other nodes (full mesh)

kube-netcheck runs on every node as daemonset. Every pod runs own http server. Every pod tries to reach Kubernetes service backed by all pods in the daemon set.

## Quick start

```
kubectl apply -f https://raw.githubusercontent.com/giantswarm/kube-netcheck/master/examples/quickstart.yaml
```
```
sleep 60; kubectl get pods -l app=kube-netcheck -n kube-system
```

If restart number is not 0 and increasing over time than Kubernes network has issues.

## Links

- [k8s-netchecker-server](https://github.com/Mirantis/k8s-netchecker-server) is more advanced network checker.
