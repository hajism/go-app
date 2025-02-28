# go-app

# 2.A.1. Build Docker Image
```bash
docker build -t testing/welcome .
```

# 2.A.2. Run Docker image 
```bash
docker run -p 8000:8000 testing/welcome
```
# Generate Letsencrypt
```bash
helm repo add jetstack https://charts.jetstack.io
helm repo update
helm install cert-manager jetstack/cert-manager \
  --namespace cert-manager \
  --create-namespace \
  --version v1.13.2 \
  --set installCRDs=true
```

--
![Diagram] ([https://github.com/hajism/go-app/Architecture.png](https://github.com/hajism/go-app/blob/main/Architecture.png))
