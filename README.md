# apollo
VM provisioning scripts + application deployment system that focuses on one-app-per-VM. Suitable for self-hosting and hobby-scale application deployments.

## Philosophy
* One-app-per-VM: Tools such as Dokku and CapRover focuses on having multiple apps per machine/VM. I think this makes their systems more complex, more noisy-neighbor and less fault-tolerant. 
* VM as the base: By not using higher-level abstractions such as Kubernetes, the system is easier to understand and easier to make low-level customizations (if needed). It especially allows self-hosted software that doesn't not provide a Kubernetes-native or even Docker-native deployment option to run on "bare-metal", without messing with either Kubernetes or Docker.
* Only take what you need: You can take only a subpart/submodule of the system and use just that. For example, you can take only VM Provisioning and Security Setup, Database Provisioning and Reverse Proxy Provisioning in order to use a VM just as a dedicated database host.

## Features
* VM Provisioning and Security Setup (one-time)
* VM Monitoring (continuous)
* Application Deployment (one-time) and Continuous Deployment (continuous)
  * Application types
    * Private non-Docker workload
    * Public Docker workload
    * Public non-Docker workload
  * Workload types
    * Web
    * Pre-deploy
    * Background
* Environment and Secret Injection (continuous)
* Logging (continuous)
* Database Provisioning
  * Types
    * MongoDB
    * Postgres
    * Redis
    * ElasticSearch
  * Tasks
    * Provisioning (one-time)
    * Auto injection (continuous)
    * Backup (continuous)
    * Restore (one-time)
* Reverse Proxy Provisioning
  * Types
    * Caddy
    * Nginx
  * Tasks
    * Provisioning (one-time)
* Application Monitoring (continuous)
