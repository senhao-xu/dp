# dp

一个解决docker镜像无法下载的问题，使用[dockerproxy](dockerproxy.com)进行代理。

## QuickStart

```bash
wget https://github.com/senhao-xu/dp/releases/download/Latest/dp /usr/local/bin/ && chmod -R 777 /usr/local/bin/dp
dp [registry.k8s.io/ingress-nginx/controller:v1.4.0]
```
## SupportRepo
- "ghcr.io"
- "gcr.io"
- "k8s.gcr.io"
- "registry.k8s.io"
- "quay.io"
- "mcr.microsoft.com"
