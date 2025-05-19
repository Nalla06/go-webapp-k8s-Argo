---
title: "Deploying Microservices on AWS: A DevOps Journey with Kubernetes and Rancher"
datePublished: Mon May 19 2025 12:56:51 GMT+0000 (Coordinated Universal Time)
cuid: cmav3ag2o000409l7etz5bgn8
slug: deploying-microservices-on-aws-a-devops-journey-with-kubernetes-and-rancher
cover: https://cdn.hashnode.com/res/hashnode/image/upload/v1747596892031/cc2898e2-4401-4bab-99d5-e9031856a500.png

---

## Introduction

In this blog post, I share my journey of building a production-ready microservices application on AWS.Throughout this project, I transformed a simple Spring Pet Clinic application into a strong, scalable, and maintainable microservices setup, complete with proper CI/CD pipelines for development, testing, staging, and production environments.

**GitHub Repository**: [Microservices-deploy-EKS-Rancher](https://github.com/Nalla06/Microservices-deploy-EKS-Rancher.git)

Let's dive into how I implemented this end-to-end DevOps project, with each step carefully documented and proven with real outputs.

### Project Overview

This project implements the Spring PetClinic application with a microservices architecture deployed on AWS using:

* Docker for containerization
    
* Kubernetes for orchestration
    
* Jenkins for CI/CD pipelines
    
* Rancher for Kubernetes management
    
* Prometheus & Grafana for monitoring
    
* AWS services including EC2, ECR, EKS, and Route53
    

The journey is divided into key phases:

1. Setting up the local development environment
    
2. Building a CI pipeline
    
3. Creating QA automation environments
    
4. Establishing staging and production environments
    
5. Implementing monitoring and observability
    

## Phase 1: Local Development Environment Setup

### Setting Up the Development Server

I began by establishing a consistent development environment on an Amazon Linux 2 EC2 instance. This involved installing essential tools including Docker, Docker Compose, Java 11, and Git to ensure all team members work with identical configurations.

Later, I enhanced this process by creating Terraform files to automate the development environment setup. This approach made the environment creation reproducible, consistent, and significantly faster to deploy. The Terraform configuration provisions the necessary EC2 instance with all required software pre-installed, saving valuable development time.

### Repository Setup and Branching Strategy

For the project foundation, I cloned the base PetClinic microservices application from the Clarusway repository. Then, I established a proper Git workflow with three main branches:

* `main`: Contains production-ready code
    
* `dev`: Houses ongoing development code
    
* `release`: Holds staging/pre-production code
    

This branching strategy follows GitFlow principles, allowing for organized feature development, integration, and controlled releases. For each task in the project, I created feature branches (like feature/msp-4) that were later merged into the dev branch after successful completion.

### Maven Build Validation

After setting up the repository structure, I validated the Maven build process for the application. This step confirmed that the application could be properly built, tested, and packaged on the development environment. I verified all Maven phases including test, package, and install to ensure the build process worked flawlessly.

### Packaging Script Development

To streamline the build process, I created a dedicated script for packaging the application with the Maven wrapper. This script simplified the build process, making it more accessible to developers who might not be familiar with Maven commands. It also ensured consistency in how the application was packaged across different environments.

### Dockerizing the Application

A critical step in modernizing the application was containerizing each microservice. I created Dockerfiles for each component of the PetClinic application, including:

* Admin Server
    
* API Gateway
    
* Config Server
    
* Discovery Server
    
* Customers Service
    
* Vets Service
    
* Visits Service
    
* Hystrix Dashboard
    

Each Dockerfile was tailored to the specific microservice's requirements while maintaining a consistent container structure. Following this, I developed a build script to automate the image building process for all microservices, streamlining the containerization workflow.

### Local Deployment with Docker Compose

To test the entire application locally before moving to more complex environments, I created a Docker Compose configuration. This setup allowed me to verify the application's functionality and inter-service communication on my development machine. Additionally, I wrote a script to test the deployment locally, making it easy to validate changes quickly during development.

## Phase 2: CI Server and Testing Setup

### Jenkins Server Configuration

For Continuous Integration and Continuous Deployment, I set up a dedicated Jenkins server on an EC2 instance. This involved installing Jenkins and configuring all necessary plugins for GitHub integration, Docker integration, AWS integration, and Pipeline support. The Jenkins server became the backbone of our automated build and deployment processes.

I configured the Jenkins server specifically for the PetClinic project, setting up necessary credentials, pipeline libraries, and build environments. This configuration allowed for automated builds, tests, and deployments across all environments.

### Implementing Unit Tests

To ensure code quality, I implemented unit tests for critical components of the application. These tests covered key functionality across different microservices, verifying that individual services behaved as expected. I focused on writing meaningful tests that would catch potential issues early in the development process.

Additionally, I updated the project's POM files to include code coverage reporting with JaCoCo. This allowed me to track test quality and identify areas of the codebase that needed additional test coverage. I could generate code coverage reports manually during development and automatically during CI builds.

### CI Pipeline Implementation

Using Jenkins Pipeline as Code, I created a comprehensive CI pipeline that runs on all development branches. This pipeline automatically:

* Builds the application
    
* Runs unit tests
    
* Generates code coverage reports
    
* Reports build status back to GitHub
    

The CI pipeline ensures that all code changes are automatically validated, maintaining code quality throughout the development process. It runs on all feature branches, bugfix branches, and the dev branch, providing quick feedback to developers.

### Selenium Tests for QA Automation

For end-to-end testing, I implemented Selenium tests that verify the application's functionality through the UI. These tests simulate user interactions with the PetClinic application, ensuring that the application works correctly from the user's perspective.

I created three Selenium jobs for QA automation tests and verified that they run successfully against the local environment. These tests were later integrated into the nightly build pipeline to ensure continuous validation of the application's UI functionality.

## Phase 3: QA Automation Environment

### Setting Up AWS ECR Registry

For storing our Docker images, I set up an AWS ECR (Elastic Container Registry). I created repositories for each microservice in the PetClinic application, providing a secure and reliable location to store container images. This was done manually at first using a Jenkins job, establishing the foundation for our container image management.

### Kubernetes QA Environment Creation

Next, I established a QA automation environment using Kubernetes. This environment provided a platform for automated testing that closely resembled the production environment. The Kubernetes setup involved:

1. Creating the necessary cluster infrastructure using Ansible playbooks.
    
2. Configuring networking and security groups
    
3. Setting up the Kubernetes control plane and worker nodes
    

This environment was used to validate changes in an environment similar to production before those changes were promoted to higher environments.

### Preparing Kubernetes YAML Files

To deploy the application to Kubernetes, I prepared comprehensive Kubernetes YAML files for all PetClinic microservices. These files defined:

* Deployments for each microservice
    
* Service definitions for inter-service communication
    
* ConfigMaps for configuration
    
* Persistent volume claims for stateful services
    
* Ingress rules for external access
    

While raw YAML files were created for full customization, **Helm charts** wrapped these manifests to enable parameterized, versioned deployments, which greatly facilitated environment-specific configurations.

### QA Automation Pipeline

With the QA environment in place, I created a QA automation pipeline on the dev branch for nightly builds. This pipeline:

1. Builds all microservices
    
2. Runs unit tests
    
3. Builds and pushes Docker images to ECR
    
4. Deploys the application to the QA Kubernetes environment using Helm to manage releases efficiently
    
5. Runs Selenium tests against the deployed application
    

This nightly build process ensured continuous validation of the application in a production-like environment, catching integration issues early.

## Phase 4: QA Environment for Release Validation

### Creating a Permanent QA Infrastructure

For release validation, I created a permanent QA infrastructure using eksctl to manage a Kubernetes cluster on AWS. This environment was more stable than the automation environment and was used for final validation before releasing to production.

The cluster was set up with:

* High availability across multiple availability zones
    
* Proper node sizing for performance testing
    
* Integration with AWS services for logging and monitoring
    

### Build Scripts for QA Environment

I prepared specialized build scripts for the QA environment that handled:

* Versioning of container images
    
* Tagging strategies for release candidates
    
* Configuration management for the QA environment
    

These scripts ensured consistent and repeatable builds for the QA environment, facilitating reliable testing of release candidates.

### Manual Deployment Pipeline

Initially, I built and deployed the application to the QA environment manually using Jenkins jobs. This allowed me to validate the deployment process and troubleshoot any issues before automating the process fully. These manual jobs served as templates for the automated pipeline.

### Weekly QA Pipeline

Once the manual process was validated, I created an automated QA pipeline for the release branch to support weekly builds. This pipeline:

1. Builds all microservices from the release branch
    
2. Versions and tags images appropriately
    
3. Deploys to the QA environment
    
4. Runs a full suite of tests
    
5. Generates reports for release readiness
    

This weekly build process validated release candidates thoroughly before they were promoted to staging.

## Phase 5: Staging and Production Setup

### High-Availability RKE Kubernetes Cluster

For staging and production environments, I prepared a high-availability Rancher Kubernetes Engine (RKE) cluster on AWS EC2. This involved:

* Setting up EC2 instances across multiple availability zones
    
* Configuring a resilient etcd cluster
    
* Establishing redundant control plane nodes
    
* Creating autoscaling worker node groups
    

This HA cluster provided the reliability and scalability required for production workloads.

### Rancher Installation and Configuration

On top of the RKE Kubernetes cluster, I installed Rancher for simplified Kubernetes management. Rancher provided:

* A user-friendly interface for cluster management
    
* Simplified workload deployment
    
* Integrated monitoring and logging
    
* Multi-cluster management capabilities
    
* Role-based access control
    

Rancher significantly improved the operational aspects of managing the Kubernetes clusters.

### Nexus Server for Artifact Management

I set up and configured a Nexus server for managing artifacts in the pipeline. This server:

* Stored Maven artifacts
    
* Cached dependencies to speed up builds
    
* Provided a private Docker registry for internal images
    
* Managed release versioning
    

Nexus improved build times and provided reliable artifact storage across the pipeline.

### Staging Pipeline Creation

With the infrastructure in place, I prepared a comprehensive staging pipeline on the Jenkins server. This pipeline:

1. Built from the release branch
    
2. Published artifacts to Nexus
    
3. Built and tagged Docker images
    
4. Deployed to the staging environment on the RKE cluster
    
5. Ran smoke tests to validate the deployment
    

The staging environment mirrored production exactly, allowing for final validation before release.

## Phase 6: Production Deployment

### Production Pipeline Implementation

The culmination of the project was the production pipeline on Jenkins. This pipeline:

1. Built from the main branch
    
2. Used approved artifacts from Nexus
    
3. Deployed to the production environment on the RKE cluster
    
4. Included approval gates for controlled releases
    
5. Implemented blue-green deployment for zero-downtime updates
    

The Jenkins production pipeline leveraged Helm to deploy microservices onto the RKE cluster, facilitating smooth updates, rollbacks, and maintaining deployment consistency.

### Domain Name and TLS Configuration

For the production environment, I set up a proper domain name and TLS configuration using AWS Route 53 and Let's Encrypt. This included:

* Configuring DNS records in Route 53
    
* Setting up certificate management
    
* Implementing TLS termination in Kubernetes
    
* Configuring redirect rules for HTTP to HTTPS
    

These steps ensured that the production application was secure and professionally presented to users.

### Monitoring Implementation

To complete the production setup, I implemented comprehensive monitoring using Prometheus and Grafana. This monitoring stack provided:

* Real-time metrics for all microservices
    
* Resource utilization tracking
    
* Custom dashboards for key performance indicators
    
* Alerting for critical issues
    

This monitoring setup enabled proactive issue detection and resolution, improving the overall reliability of the application.

## Conclusion

Throughout this project, I successfully implemented a complete DevOps lifecycle for the microservices application on AWS. By following these steps, I transformed a simple application into a production-ready, containerized, and orchestrated microservices architecture with proper CI/CD pipelines across all environments.

The key takeaways from this project include:

1. The importance of consistent development environments
    
2. The value of comprehensive test automation
    
3. The benefits of container orchestration with Kubernetes
    
4. The power of pipeline automation for reliable releases
    
5. The necessity of proper monitoring for production applications