# apollo
VM provisioning scripts + application deployment system that focuses on one-app-per-VM. Suitable for self-hosting and hobby-scale application deployments.

## Features
* VM Provisioning (one-time)
* VM Security setup (one-time)
* VM Monitoring (continuous, to a central location across VMs)
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
* Logging (continuous, to a central location across VMs)
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
* Application Monitoring (continuous, to a central location across VMs)
