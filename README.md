# GoFlow

> **Secure, observable, GitOps-driven cloud-native task processing platform built with Go and Kubernetes.**

GoFlow is a containerized task-processing application designed to demonstrate how a modern application can be **built, secured, deployed, monitored, and managed using DevSecOps and cloud-native practices**.

Unlike a basic CRUD or Docker deployment project, GoFlow combines **CI/CD security, Kubernetes hardening, GitOps, observability, AWS infrastructure, IAM least privilege, and cost monitoring** into a single end-to-end system.

---

## 🚀 Why GoFlow?

GoFlow demonstrates a practical production-style workflow:

**Code → Test → Security Scan → Container → GitOps Deployment → Kubernetes → Monitoring → Alerting**

It is useful for learning and demonstrating:

* Cloud Security and DevSecOps
* Kubernetes deployment and hardening
* Infrastructure as Code
* GitOps with Argo CD
* Application and infrastructure monitoring
* IAM least-privilege design
* Secure cloud database integration
* Automated security scanning and CI/CD

---

## 🏗️ Architecture

```text
                         GitHub Repository
                                │
                                ▼
                     GitHub Actions CI/CD
                     ┌────────────────────┐
                     │ Go Test / go vet   │
                     │ Docker Build       │
                     │ Trivy Scan         │
                     └─────────┬──────────┘
                               │
                               ▼
                      GitHub Container Registry
                               │
                               ▼
                           Argo CD
                         (GitOps Sync)
                               │
                               ▼
                    Kubernetes / Docker Desktop
                    ┌──────────────────────────┐
                    │                          │
                    │      GoFlow × 2          │
                    │          │               │
                    │          ▼               │
                    │      Prometheus          │
                    │          │               │
                    │          ▼               │
                    │       Grafana             │
                    │                          │
                    └──────────┬───────────────┘
                               │
                               ▼
                     AWS RDS PostgreSQL
                               │
                               ▼
                        AWS CloudWatch
                               │
                               ▼
                             SNS
                               │
                               ▼
                            Email
```

---

## ✨ Key Features

### Application

* Go-based task processing API
* Multiple workers for background task execution
* REST API
* Health/readiness endpoints
* Prometheus `/metrics` endpoint
* PostgreSQL database integration

### Kubernetes Security

* Runs as a non-root user
* Read-only root filesystem
* Drops Linux capabilities
* `allowPrivilegeEscalation: false`
* Seccomp `RuntimeDefault`
* Disabled automatic service-account token mounting
* CPU and memory requests/limits
* Readiness and liveness probes
* Kubernetes `NetworkPolicy` restricting application traffic

### DevSecOps CI/CD

* GitHub Actions
* Automated unit testing
* `go vet`
* Docker image build
* Trivy vulnerability scanning
* CI fails on HIGH/CRITICAL image vulnerabilities
* Multi-architecture container images
* GitHub Container Registry (GHCR)
* Immutable image SHA deployment

### GitOps

* Argo CD continuously tracks the Git repository
* Automated synchronization
* Automated pruning
* Self-healing Kubernetes deployments
* Kubernetes configuration maintained as code

### AWS Infrastructure

* Amazon RDS PostgreSQL
* Encryption at rest
* Restricted security group
* Automated backup retention
* Terraform-managed infrastructure
* AWS CloudWatch monitoring
* Amazon SNS email notifications

### IAM & Least Privilege

GoFlow separates administrative and monitoring access.

```text
Monitoring User
      │
      │ MFA
      ▼
GoFlowMonitoringRole
      │
      ▼
CloudWatch Read-Only Policy
```

The monitoring identity is not given unnecessary permissions such as:

* RDS modification/deletion
* IAM administration
* Security group modification
* EC2 administration
* CloudWatch configuration changes

### Monitoring & Alerting

#### Application Monitoring

**Prometheus + Grafana**

Monitors GoFlow application metrics such as:

* Queue length
* Task failures
* Task retries
* Worker utilization
* CPU usage
* Memory usage

#### AWS/RDS Monitoring

**CloudWatch**

Monitors:

* CPU utilization
* Free storage
* Database connections
* Freeable memory
* Read latency
* Write latency

Alerts are delivered using:

**CloudWatch → SNS → Email**

### Cost Protection

AWS cost monitoring is included to reduce the risk of unexpected charges:

* AWS Budget
* CloudWatch billing alarm
* Monthly cost threshold monitoring

---

## 🔐 Security Highlights

GoFlow follows a defense-in-depth approach:

```text
GitHub
  ↓
CI Tests
  ↓
Trivy Security Scan
  ↓
Immutable Container Image
  ↓
Argo CD GitOps
  ↓
Hardened Kubernetes Pod
  ↓
NetworkPolicy
  ↓
Encrypted RDS
  ↓
IAM Least Privilege
  ↓
CloudWatch + Alerting
```

Secrets such as database credentials are kept outside source code and injected through Kubernetes Secrets.

---

## 🛠️ Tech Stack

| Category               | Technologies              |
| ---------------------- | ------------------------- |
| Language               | Go                        |
| API Framework          | Gin                       |
| Database               | PostgreSQL                |
| Containerization       | Docker                    |
| Orchestration          | Kubernetes                |
| Monitoring             | Prometheus, Grafana       |
| GitOps                 | Argo CD                   |
| Infrastructure as Code | Terraform                 |
| Cloud                  | AWS                       |
| Database Service       | Amazon RDS                |
| Cloud Monitoring       | Amazon CloudWatch         |
| Notifications          | Amazon SNS                |
| Identity & Access      | AWS IAM                   |
| CI/CD                  | GitHub Actions            |
| Container Registry     | GitHub Container Registry |
| Security Scanning      | Trivy                     |
| Version Control        | Git, GitHub               |

---

## 📂 Project Structure

```text
GoFlow/
├── cmd/
├── internal/
├── k8s/
│   ├── deployment.yaml
│   ├── service.yaml
│   ├── configmap.yaml
│   ├── secret.yaml
│   └── network-policy.yaml
├── terraform/
│   ├── main.tf
│   ├── provider.tf
│   ├── variables.tf
│   ├── outputs.tf
│   ├── iam.tf
│   └── cloudwatch.tf
├── .github/
│   └── workflows/
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── README.md
```

---

## ⚙️ Deployment Flow

```text
1. Developer pushes code
2. GitHub Actions runs tests and go vet
3. Docker image is built
4. Trivy scans the image
5. Image is pushed to GHCR
6. Kubernetes manifests are managed through Git
7. Argo CD detects repository changes
8. Argo CD synchronizes Kubernetes
9. GoFlow runs with 2 replicas
10. Prometheus collects metrics
11. Grafana visualizes application health
12. AWS CloudWatch monitors RDS
13. SNS sends infrastructure alerts
```

---

## 💡 What Makes GoFlow Different?

GoFlow is not just a Go application or Kubernetes deployment.

It brings together multiple real-world engineering practices in one project:

* **DevSecOps:** security scanning integrated into CI/CD
* **GitOps:** deployments controlled through Argo CD
* **Cloud Security:** IAM least privilege, MFA, encrypted RDS
* **Kubernetes Security:** hardened containers and NetworkPolicies
* **Observability:** Prometheus + Grafana for application monitoring
* **Cloud Monitoring:** CloudWatch for AWS infrastructure
* **Automated Alerting:** SNS-based notifications
* **Infrastructure as Code:** Terraform
* **Cost Awareness:** AWS budget and billing monitoring
* **Immutable Deployments:** version-pinned container images

This makes GoFlow a practical demonstration of how **application development, cloud infrastructure, security, deployment, and observability can work together as one system**.

---

## 🎯 Project Outcome

GoFlow demonstrates an end-to-end **cloud-native DevSecOps and cloud-security workflow** using Kubernetes and AWS without relying on managed Kubernetes services.

It provides hands-on experience with:

**Go + Docker + Kubernetes + Terraform + AWS + IAM + GitHub Actions + Argo CD + Prometheus + Grafana + Trivy**

---

## 🔮 Future Scope

Possible future improvements include:

* AWS EKS migration
* Private RDS networking
* External secret management
* Advanced SIEM/SOAR integration
* Automated incident response
* Horizontal Pod Autoscaling
* More advanced distributed tracing
* Production-grade multi-environment deployments

---

## 📌 Status

**Project Status: Completed ✅**

GoFlow currently provides a complete development-to-deployment pipeline with **security, GitOps, Kubernetes orchestration, observability, AWS infrastructure monitoring, IAM controls, and cost protection**.
