---
title: "Deploying a Go Web App in Kubernetes with Argo CD: A Step-by-Step Guide"
slug: deploying-a-go-web-app-in-kubernetes-with-argo-cd-a-step-by-step-guide
cover: https://cdn.hashnode.com/res/hashnode/image/upload/v1742997056793/196f1f98-d3a0-4319-850f-bb507bfe26b7.png

---

## Introduction

In this blog post, I will walk you through how I deployed my Go web application in a Kubernetes (K8s) cluster using Argo CD for Git Ops-based continuous deployment. By leveraging Kubernetes for container orchestration and Argo CD for automated deployments, I was able to achieve a streamlined and scalable deployment process.

## Prerequisites

Before we begin, ensure you have the following:

* A Go web application (basic HTTP server)
    
* Docker installed for containerization
    
* Kubernetes cluster (Minikube, Kind, or cloud-based like GKE, EKS, or AKS)
    
* kubectl installed and configured
    
* ArgoCD installed on your Kubernetes cluster
    
* GitHub repository for GitOps integration
    

## Step 1: Containerizing the Go Web App

First, let's create a `Dockerfile` to containerize our Go web application:

```dockerfile
# Use the official Golang image for building the app
FROM golang:1.20 AS builder

WORKDIR /app
COPY . .

RUN go mod tidy && go build -o app

# Use a minimal image for running the app
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/app .

CMD ["./app"]
```

Now, build and push the Docker image:

```sh
docker build -t your-dockerhub-username/go-web-app:latest .
docker push your-dockerhub-username/go-web-app:latest
```

## Step 2: Writing Kubernetes Manifests

### Deployment (`deployment.yaml`)

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: go-web-app
spec:
  replicas: 2
  selector:
    matchLabels:
      app: go-web-app
  template:
    metadata:
      labels:
        app: go-web-app
    spec:
      containers:
        - name: go-web-app
          image: your-dockerhub-username/go-web-app:latest
          ports:
            - containerPort: 8080
```

### Service (`service.yaml`)

```yaml
apiVersion: v1
kind: Service
metadata:
  name: go-web-app-service
spec:
  selector:
    app: go-web-app
  ports:
    - protocol: TCP
      port: 80
      targetPort: 8080
  type: LoadBalancer
```

Apply these manifests to Kubernetes:

```sh
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
kubectl apply -f ingress.yaml
```

## Step 3: Setting Up ArgoCD

Install ArgoCD:

```sh
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```

Login to ArgoCD CLI:

```sh
kubectl port-forward svc/argocd-server -n argocd 8080:443
argocd login localhost:8080
```

Set up an ArgoCD application for GitOps:

```yaml
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: go-web-app
  namespace: argocd
spec:
  destination:
    namespace: default
    server: https://kubernetes.default.svc
  source:
    repoURL: https://github.com/your-repo/go-k8s-app.git
    targetRevision: HEAD
    path: manifests
  syncPolicy:
    automated:
      prune: true
      selfHeal: true
```

Apply this configuration:

```sh
kubectl apply -f argocd-app.yaml
```

## Step 4: Continuous Deployment with GitOps

Once ArgoCD is monitoring your GitHub repository, any updates pushed to your Kubernetes manifests will trigger an automatic deployment to the cluster. This ensures a seamless CI/CD process.

## Conclusion

In this article, we covered:

* How to containerize a Go web app
    
* Deploying it to Kubernetes
    
* Using ArgoCD for GitOps-based continuous deployment
    

This setup improves scalability, reliability, and automation for your Go web applications in Kubernetes. Have any questions or improvements? Let me know in the comments! 🚀