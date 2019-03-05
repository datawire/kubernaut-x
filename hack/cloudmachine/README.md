# Cloud (Virtual) Machine

Manages a vanilla (most likely virtual) machine in the cloud. It is designed to be easily extended with additional 
providers as needed (e.g. DigitalOcean or Microsoft Azure).

# Purpose

Kubernetes clusters run on cloud compute infrastructure and the setup logic needs to be developed and tested. This tool
aides that work.

# Commands

The commands are relatively straightforward:

| Command                | Description                     |
| ---------------------- | ------------------------------- |
| `make start.aws`       | Start your AWS EC2 machine      |
| `make shutdown.aws`    | Stop your AWS EC2 machine       |
| `make start.gcloud`    | Start your Google Cloud machine |
| `make shutdown.gcloud` | Stop your Google Cloud machine  |
| `make ssh`             | SSH into the virtual machine    | 
