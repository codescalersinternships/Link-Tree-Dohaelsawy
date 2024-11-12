# Link Tree ~ Deployment Section

## Directory Structure:
```golang
 .
├── helm
├── kubernetes
└── README.md
```
## Deployment Options:
- 1- Deploy using Kubernetes files to learn how to apply kubectl commands and control single deployments
  - to apply all files once `` kubectl apply ./kubernetes/``
- 2- Deploy using Helm charts to make things easier
  - to deploy backend using helm `` helm install backend helm/backend-charts ``
  - to deploy frontend using helm `` helm install frontend helm/frontend-charts ``
 
>[!NOTE]
> Don't forget to customize your own env inside Kubernetes directory or change values for helm charts
