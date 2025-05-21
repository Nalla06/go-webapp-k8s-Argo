---
title: "Deploying a Scalable Three-Tier Application on AWS with EKS, Argo CD, Jenkins, SonarQube, and Prometheus-Grafana Monitoring"
datePublished: Wed May 21 2025 09:13:46 GMT+0000 (Coordinated Universal Time)
cuid: cmaxq79sj001h09jydct5ewv1
slug: deploying-a-scalable-three-tier-application-on-aws-with-eks-argo-cd-jenkins-sonarqube-and-prometheus-grafana-monitoring
cover: https://cdn.hashnode.com/res/hashnode/image/upload/v1747313111841/265ba88c-2bdc-4751-8606-355a1aecf8fd.gif

---

### **A hands-on guide to deploying a full-stack 3-tier application on AWS using EKS, ArgoCD, Jenkins, Prometheus, and Grafana.**

In AWS, there are several ways to deploy a 3-tier architecture depending on your use case, team size, automation requirements, and scalability goals. Each method offers a unique blend of control, ease of use, and integration with DevOps tools. Below are the most common deployment strategies I explored before choosing the one for my project:

---

### 1️⃣ **Manual Deployment using EC2 Instances**

This is the most basic method — spinning up individual EC2 instances for:

* **Frontend:** Hosted with Nginx or Apache
    
* **Backend:** Hosted with Node.js, Spring Boot, etc.
    
* **Database:** Deployed on another EC2 or managed via Amazon RDS
    

📝 **Pros:** Full control, good for learning  
⚠️ **Cons:** Hard to scale, manual setup, not cloud-native

---

### 2️⃣ **Using AWS Elastic Beanstalk**

Elastic Beanstalk abstracts the infrastructure and lets you deploy:

* Backend and frontend apps in separate environments
    
* Databases using Amazon RDS integrations
    

📝 **Pros:** Simplifies deployment, minimal DevOps needed  
⚠️ **Cons:** Limited flexibility and customization

---

### 3️⃣ **Using Amazon ECS with Fargate**

Containers are deployed as services:

* Frontend and backend as Docker containers
    
* Backend services talk to an Amazon RDS instance
    

📝 **Pros:** No server management, scalable  
⚠️ **Cons:** Less control compared to EKS

---

### 4️⃣ **Using Amazon EKS (Elastic Kubernetes Service)**

My chosen method for this project.

EKS offers the full power of Kubernetes:

* **Frontend:** React/Vue app exposed via Ingress
    
* **Backend:** Microservices in pods behind a ClusterIP or LoadBalancer service
    
* **Database:** External RDS or in-cluster StatefulSet (dev)
    

📝 **Pros:** Scalable, production-ready, integrates well with CI/CD  
⚠️ **Cons:** Steeper learning curve, requires proper automation

---

### 5️⃣ **Serverless Approach**

Perfect for lightweight apps:

* **Frontend:** Static files on S3 + CloudFront
    
* **Backend:** AWS Lambda with API Gateway
    
* **Database:** DynamoDB or Aurora Serverless
    

📝 **Pros:** No server management, cost-efficient  
⚠️ **Cons:** Not ideal for complex or stateful apps

---

After careful consideration, I selected **Amazon EKS** for this project due to its flexibility, scalability, and robust support for DevOps tooling. This blog details my implementation journey.

## Project Architecture Overview

My project implements:

* Infrastructure as Code (IaC) using Terraform
    
* CI/CD pipeline with Jenkins
    
* Security scanning (SonarQube, OWASP, Trivy)
    
* Kubernetes deployment on AWS EKS
    
* GitOps with ArgoCD
    
* Monitoring with Prometheus & Grafana
    
* Three-tier application (React frontend, Node.js backend, MongoDB)
    
* Persistent storage for database
    
* Custom domain configuration
    

## Implementation Overview

Here's a summary of the steps I completed in this project:

1. Created an IAM user with necessary permissions for deployment and management
    
2. Used Terraform and AWS CLI to set up the Jenkins server on AWS EC2
    
3. Configured Jenkins with Docker, SonarQube, Terraform, Kubectl, AWS CLI, and Trivy
    
4. Deployed Amazon EKS cluster using eksctl commands
    
5. Configured AWS Application Load Balancer (ALB) for the EKS cluster
    
6. Created private Amazon ECR repositories for frontend and backend Docker images
    
7. Installed and configured ArgoCD for GitOps deployment
    
8. Integrated SonarQube for code quality analysis
    
9. Created Jenkins pipelines for backend and frontend deployment
    
10. Implemented monitoring with Helm, Prometheus, and Grafana
    
11. Deployed the three-tier application using ArgoCD
    
12. Configured DNS settings with custom subdomains
    
13. Implemented persistent volumes for database pods
    
14. Set up comprehensive monitoring dashboards in Grafana
    

## Prerequisites

* AWS account with appropriate permissions
    
* GitHub account
    
* Domain name (for proper TLS and subdomain configuration)
    
* Basic knowledge of Kubernetes, Docker, and AWS services
    

## Step 1: Setting Up IAM User

First, I created an IAM user with the necessary permissions to deploy and manage resources:

1. Log in to AWS Management Console and navigate to IAM service
    
2. Create a new user with programmatic access
    
3. Attach policies for:
    
    * AmazonECR-FullAccess
        
    * AmazonEKSClusterPolicy
        
    * AmazonRoute53FullAccess
        
    * EC2FullAccess
        
    * IAMFullAccess
        

## Step 2: Setting Up Infrastructure with Terraform

Next, I provisioned the Jenkins server using Infrastructure as Code:  
In this comprehensive guide, I'll walk you through implementing a complete DevSecOps pipeline for a three-tier application on AWS EKS with security scanning, monitoring, and GitOps deployment. This blog details the exact steps I took to successfully implement this project.

## Project Architecture Overview

## Implementation Overview

Here's a summary of the steps I completed in this project:

1. Created an IAM user with necessary permissions for deployment and management
    
2. Used Terraform and AWS CLI to set up the Jenkins server on AWS EC2
    
3. Configured Jenkins with Docker, SonarQube, Terraform, Kubectl, AWS CLI, and Trivy
    
4. Deployed Amazon EKS cluster using eksctl commands
    
5. Configured AWS Application Load Balancer (ALB) for the EKS cluster
    
6. Created private Amazon ECR repositories for frontend and backend Docker images
    
7. Installed and configured ArgoCD for GitOps deployment
    
8. Integrated SonarQube for code quality analysis
    
9. Created Jenkins pipelines for backend and frontend deployment
    
10. Implemented monitoring with Helm, Prometheus, and Grafana
    
11. Deployed the three-tier application using ArgoCD
    
12. Configured DNS settings with custom subdomains
    
13. Implemented persistent volumes for database pods
    
14. Set up comprehensive monitoring dashboards in Grafana
    

## Prerequisites

* AWS account with appropriate permissions
    
* GitHub account
    
* Domain name (for proper TLS and subdomain configuration)
    
* Basic knowledge of Kubernetes, Docker, and AWS services
    

## Step 1: Setting Up IAM User

First, I created an IAM user with the necessary permissions to deploy and manage resources:

1. Log in to AWS Management Console and navigate to IAM service
    
2. Create a new user with programmatic access
    
3. Attach policies for:
    
    * AmazonECR-FullAccess
        
    * AmazonEKSClusterPolicy
        
    * AmazonRoute53FullAccess
        
    * EC2FullAccess
        
    * IAMFullAccess
        

## Step 2: Setting Up Infrastructure with Terraform

Next, I provisioned the Jenkins server using Infrastructure as Code:

1. Clone the repository:
    

```bash
git clone https://github.com/Nalla06/End-to-End-3-tier-DevSecops-Project.git
cd End-to-End-3-tier-DevSecops-Project/terraform-jenkins
```

2. Initialize and apply Terraform:
    

```bash
terraform init
terraform apply -auto-approve
```

This created:

* EC2 instance for Jenkins
    
* Security groups with appropriate port access
    
* IAM roles with necessary permissions
    

3. Retrieve Jenkins initial admin password:
    

```bash
ssh -i your-key.pem ubuntu@<jenkins-server-ip>
sudo cat /var/lib/jenkins/secrets/initialAdminPassword
```

4. Access Jenkins at `http://<jenkins-server-ip>:8080` and complete setup:
    
    * Install suggested plugins
        
    * Create admin user
        
    * Configure Jenkins URL
        

## Step 3: Jenkins Server Configuration

I installed and configured several essential tools on the Jenkins server:

1. Install Docker:
    

```bash
sudo apt update
sudo apt install docker.io -y
sudo usermod -aG docker jenkins
sudo systemctl restart jenkins
```

2. Install SonarQube via Docker:
    

```bash
docker run -d --name sonarqube -p 9000:9000 sonarqube:lts
```

3. Install Terraform:
    

```bash
wget -O- https://apt.releases.hashicorp.com/gpg | sudo gpg --dearmor -o /usr/share/keyrings/hashicorp-archive-keyring.gpg
echo "deb [signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] https://apt.releases.hashicorp.com $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/hashicorp.list
sudo apt update && sudo apt install terraform -y
```

4. Install kubectl:
    

```bash
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
```

5. Install AWS CLI:
    

```bash
curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
sudo ./aws/install
```

6. Install Trivy:
    

```bash
sudo apt-get install wget apt-transport-https gnupg lsb-release -y
wget -qO - https://aquasecurity.github.io/trivy-repo/deb/public.key | sudo apt-key add -
echo deb https://aquasecurity.github.io/trivy-repo/deb $(lsb_release -sc) main | sudo tee -a /etc/apt/sources.list.d/trivy.list
sudo apt-get update
sudo apt-get install trivy -y
```

## Step 4: Creating AWS EKS Cluster

For my EKS cluster, I used eksctl instead of Terraform to have more hands-on control:

1. Install eksctl on the Jenkins server:
    

```bash
curl --silent --location "https://github.com/weaveworks/eksctl/releases/latest/download/eksctl_$(uname -s)_amd64.tar.gz" | tar xz -C /tmp
sudo mv /tmp/eksctl /usr/local/bin
```

2. Create a cluster configuration file `eks-cluster.yaml`:
    

```bash
yamlapiVersion: eksctl.io/v1alpha5
kind: ClusterConfig
metadata:
  name: my-eks-cluster
  region: us-east-1
  version: "1.27"
nodeGroups:
  - name: ng-1
    instanceType: t3.medium
    desiredCapacity: 2
    minSize: 2
    maxSize: 4
    volumeSize: 20
```

3. Create the EKS cluster:
    

```bash
eksctl create cluster -f eks-cluster.yaml
```

4. Configure kubectl to work with your EKS cluster:
    

```bash
aws eks update-kubeconfig --name my-eks-cluster --region us-east-1
kubectl get nodes
```

## Step 5: Load Balancer Configuration

I configured the AWS Application Load Balancer (ALB) for the EKS cluster:

1. Install AWS Load Balancer Controller:
    

```bash
curl -O https://raw.githubusercontent.com/kubernetes-sigs/aws-load-balancer-controller/main/docs/install/iam_policy.json
aws iam create-policy --policy-name AWSLoadBalancerControllerIAMPolicy --policy-document file://iam_policy.json
```

2. Create IAM service account:
    

```bash
eksctl create iamserviceaccount \
  --cluster=my-eks-cluster \
  --namespace=kube-system \
  --name=aws-load-balancer-controller \
  --attach-policy-arn=arn:aws:iam::<AWS_ACCOUNT_ID>:policy/AWSLoadBalancerControllerIAMPolicy \
  --override-existing-serviceaccounts \
  --approve
```

3. Install AWS Load Balancer Controller using Helm:
    

```bash
helm repo add eks https://aws.github.io/eks-charts
helm repo update
helm install aws-load-balancer-controller eks/aws-load-balancer-controller \
  -n kube-system \
  --set clusterName=my-eks-cluster \
  --set serviceAccount.create=false \
  --set serviceAccount.name=aws-load-balancer-controller
```

## Step 6: Creating Amazon ECR Repositories

I created private ECR repositories for both frontend and backend Docker images:

1. Create frontend repository:
    

```bash
aws ecr create-repository --repository-name three-tier-frontend --image-scanning-configuration scanOnPush=true
```

2. Create backend repository:
    

```bash
aws ecr create-repository --repository-name three-tier-backend --image-scanning-configuration scanOnPush=true
```

3. Log in to ECR:
    

```bash
aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin <AWS_ACCOUNT_ID>.dkr.ecr.us-east-1.amazonaws.com
```

## Step 7: Setting Up ArgoCD for GitOps

I installed ArgoCD for GitOps-based deployment:

1. Create ArgoCD namespace:
    

```bash
kubectl create namespace argocd
```

2. Install ArgoCD:
    

```bash
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```

3. Expose ArgoCD with a Load Balancer:
    

```bash
kubectl patch svc argocd-server -n argocd -p '{"spec": {"type": "LoadBalancer"}}'
```

4. Get the ArgoCD initial admin password:
    

```bash
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
```

5. Access ArgoCD UI:
    

```bash
kubectl get svc argocd-server -n argocd
```

## Step 8: Setting Up SonarQube Integration

I configured SonarQube for code quality analysis:

1. Installed the SonarQube scanner plugin in Jenkins:
    
    * Navigate to Manage Jenkins &gt; Manage Plugins
        
    * Install SonarQube Scanner plugin and restart Jenkins
        
2. Configure SonarQube in Jenkins:
    
    * Go to Manage Jenkins &gt; Configure System
        
    * Add SonarQube server with URL and authentication token
        
    * Save the configuration
        
3. Create a SonarQube Webhook for Jenkins:
    
    * In SonarQube: Administration &gt; Configuration &gt; Webhooks
        
    * Add a new webhook pointing to Jenkins URL
        

## Step 9: Setting Up Jenkins Pipelines

I created Jenkins pipelines for both backend and frontend deployment:

1. Create a new pipeline job in Jenkins:
    
    * Click "New Item" &gt; Select "Pipeline"
        
    * Configure GitHub repository in "Pipeline" section
        
    * Use the Jenkinsfile from the repository
        
2. Configure credentials in Jenkins:
    
    * Add AWS credentials
        
    * Add GitHub credentials
        
    * Add Docker Hub credentials (if needed)
        
    * Add ECR credentials
        
3. Backend Jenkinsfile includes:
    
    * Code checkout
        
    * SonarQube analysis
        
    * Docker image build
        
    * Trivy security scan
        
    * ECR image push
        
    * ArgoCD application update
        
4. Frontend Jenkinsfile follows similar pattern:
    
    * Code checkout
        
    * SonarQube analysis
        
    * Docker image build
        
    * Trivy security scan
        
    * ECR image push
        
    * ArgoCD application update
        

## Step 10: Implementing Monitoring with Prometheus and Grafana

I set up comprehensive monitoring using Helm, Prometheus, and Grafana:

1. Add Helm repository for Prometheus:
    

```bash
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
```

2. Create monitoring namespace:
    

```bash
kubectl create namespace monitoring
```

3. Install Prometheus stack with Grafana:
    

```bash
helm install prometheus prometheus-community/kube-prometheus-stack -n monitoring
```

4. Expose Grafana using LoadBalancer:
    

```bash
kubectl patch svc prometheus-grafana -n monitoring -p '{"spec": {"type": "LoadBalancer"}}'
```

5. Get Grafana access details:
    

```bash
kubectl get svc prometheus-grafana -n monitoring
```

6. Access Grafana (default credentials: admin/prom-operator) and import dashboards:
    
    * Kubernetes Cluster Overview (ID: 10856)
        
    * Node Exporter (ID: 1860)
        
    * Pods Overview (ID: 6417)
        

## Step 11: ArgoCD Application Deployment

I used ArgoCD to deploy the three-tier application:

1. Create a Git repository for Kubernetes manifests:
    
    * Added manifests for MongoDB, backend, frontend, and ingress
        
2. Define ArgoCD Application for database:
    

```bash
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: Database
  namespace: argocd
spec:
  project: default
  source:
    repoURL: https://github.com/Nalla06/kubernetes-manifests.git
    targetRevision: HEAD
    path: Database
  destination:
    server: https://kubernetes.default.svc
    namespace: three-tier-app
  syncPolicy:
    automated:
      prune: true
      selfHeal: true
```

3. Create similar ArgoCD Applications for ingress, backend and frontend components
    
4. Apply the Applications to ArgoCD:
    

```bash
kubectl apply -f argocd-applications/ -n argocd
```

5. Verify synchronization in ArgoCD UI or with:
    

```bash
kubectl get applications -n argocd
```

## Step 12: DNS Configuration

I configured DNS settings to make the application accessible via custom subdomains:

1. Create Route53 hosted zone (if not already existing):
    

```bash
aws route53 create-hosted-zone --name yourdomain.com --caller-reference $(date +%s)
```

2. Note the nameservers and update your domain registrar
    
3. Create A records for your application:
    

```bash
INGRESS_LB=$(kubectl get svc -n ingress-nginx ingress-nginx-controller -o jsonpath='{.status.loadBalancer.ingress[0].hostname}')
LB_IP=$(dig +short $INGRESS_LB | head -n1)

cat > dns-records.json <<EOF
{
  "Changes": [
    {
      "Action": "CREATE",
      "ResourceRecordSet": {
        "Name": "app.yourdomain.com",
        "Type": "A",
        "TTL": 300,
        "ResourceRecords": [
          {
            "Value": "$LB_IP"
          }
        ]
      }
    },
    {
      "Action": "CREATE",
      "ResourceRecordSet": {
        "Name": "api.yourdomain.com",
        "Type": "A",
        "TTL": 300,
        "ResourceRecords": [
          {
            "Value": "$LB_IP"
          }
        ]
      }
    }
  ]
}
EOF

aws route53 change-resource-record-sets --hosted-zone-id YOUR_ZONE_ID --change-batch file://dns-records.json
```

## Step 13: Data Persistence Implementation

I implemented persistent storage for the MongoDB database to ensure data persistence:

1. Create StorageClass for EBS volumes:
    

```bash
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: ebs-sc
provisioner: ebs.csi.aws.com
volumeBindingMode: WaitForFirstConsumer
```

2. Apply the StorageClass:
    

```bash
kubectl apply -f storage-class.yaml
```

3. Create PersistentVolumeClaim for MongoDB:
    

```bash
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: mongodb-pvc
  namespace: three-tier-app
spec:
  accessModes:
    - ReadWriteOnce
  storageClassName: ebs-sc
  resources:
    requests:
      storage: 10Gi
```

4. Apply the PVC:
    

```bash
kubectl apply -f mongodb-pvc.yaml
```

5. Update MongoDB deployment to use the PVC:
    

```bash
apiVersion: apps/v1
kind: Deployment
metadata:
  name: mongodb
  namespace: three-tier-app
spec:
  selector:
    matchLabels:
      app: mongodb
  template:
    metadata:
      labels:
        app: mongodb
    spec:
      containers:
      - name: mongodb
        image: mongo:latest
        ports:
        - containerPort: 27017
        volumeMounts:
        - name: mongodb-data
          mountPath: /data/db
      volumes:
      - name: mongodb-data
        persistentVolumeClaim:
          claimName: mongodb-pvc
```

## Step 14: Validation and Monitoring

After deploying all components, I validated the setup and configured monitoring:

1. Verify all pods are running:
    

```bash
kubectl get pods -n three-tier-app
```

2. Check ingress and services:
    

```bash
kubectl get ingress -n three-tier-app
kubectl get svc -n three-tier-app
```

3. Access application through configured domains:
    
    * Frontend: [https://app.yourdomain.com](https://app.yourdomain.com)
        
    * Backend API: [https://api.yourdomain.com/api/health](https://api.yourdomain.com/api/health)
        
4. Set up Grafana dashboards for monitoring:
    
    * In Grafana UI, import dashboard ID 6417 for Pod monitoring
        
    * Import dashboard ID 1860 for Node Exporter monitoring
        
    * Create custom dashboard for application metrics
        
5. Configure alerts in Prometheus:
    
    * Set up alerts for high CPU/memory usage
        
    * Configure notification channels (email, Slack)
        

## Conclusion

In this project, I've successfully implemented a complete DevSecOps pipeline for a three-tier application on AWS EKS. The key achievements include:

1. Infrastructure as Code using Terraform for reproducible deployment
    
2. Secure CI/CD pipeline with Jenkins, incorporating security scanning
    
3. Private ECR repositories for container images
    
4. GitOps-based deployment with ArgoCD
    
5. Comprehensive monitoring with Prometheus and Grafana
    
6. Persistent storage for database reliability
    
7. Custom domain configuration with Route53
    

This implementation demonstrates a modern approach to cloud-native application deployment with security built into every stage of the process. The architecture ensures scalability, maintainability, and follows industry best practices for Kubernetes-based applications.

By implementing this project, I've gained hands-on experience with cloud-native technologies and DevSecOps practices that are highly valuable in today's containerized application landscape.

## GitHub Repository Reference

For the complete source code, configuration files, and implementation details of this project, please visit my GitHub repository:

[ht](https://github.com/Nalla06/End-to-End-3-tier-DevSecops-Project.git)[tps://github.com/Nalla06/End-to-End-3-tier-DevSecops-Project](https://github.com/Nalla06/End-to-End-3-tier-DevSecops-Project.git)