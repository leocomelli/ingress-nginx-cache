# Ingress Nginx Cache

It is a simple way to implements a request cache using Ingress Nginx.

First of all, add the `proxy_cache_path` directive within the top-level `http { ... }` block of the generated `nginx.conf` file. When using [Helm Chart](https://github.com/kubernetes/ingress-nginx/tree/main/charts/ingress-nginx), use the following values:

```yaml
controller:
  config:
    http-snippet: |
      proxy_cache_path /tmp/nginx_cache levels=1:2 keys_zone=app-cache:32m use_temp_path=off max_size=4g inactive=24h;
```

Afterward, we can proceed to apply the Kubernetes resources to the cluster.

```bash
kubectl apply -f kubernetes.yaml
```

Adjust the configutation according to your environment.

### Testing

We can make some GET requests to make sure that the cache is working well. Just so you know, the cache takes about a minute to expire

```bash
curl https://ingress-cache.leocomelli.dev/ -i
curl https://ingress-cache.leocomelli.dev/public -i
curl -H 'Authorization: s3cr3t' https://ingress-cache.leocomelli.dev/private -i
```

Check the log entries:

```bash
 kubectl logs -f deploy/ingress-nginx-cache-app -n ingress-nginx-cache
```
