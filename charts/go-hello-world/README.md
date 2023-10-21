
# hello-world

![Version: 0.0.1](https://img.shields.io/badge/Version-0.0.1-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.0.2](https://img.shields.io/badge/AppVersion-0.0.2-informational?style=flat-square)

## Description

\"Hello World\" application with Prometheus Metrics.

## Usage

### Add Helm repository

```shell
helm repo add hello-world https://braveokafor.github.io/go-hello-world/
helm repo update
```

### Install Helm chart

Using default configurations:

```bash
helm install --generate-name hello-world/hello-world
```

Customising configurations:

```bash
helm install --generate-name --set ingress.enabled=true, hello-world/hello-world
```

## Configuration

The following table lists the configurable parameters of the Hello-World chart and their default values.

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| deployment | object | `{"containerPort":5000,"image":{"pullPolicy":"IfNotPresent","repository":"braveokafor/go-hello-world","tag":"latest"},"livenessProbe":{"initialDelaySeconds":30,"path":"/healthz","periodSeconds":10},"replicaCount":1,"resources":{"limits":{"cpu":"128m","memory":"256Mi"},"requests":{"cpu":"128m","memory":"256Mi"}}}` | Deployment configurations |
| deployment.containerPort | int | `5000` | Port on which the container will run |
| deployment.image.pullPolicy | string | `"IfNotPresent"` | Pull policy of the image |
| deployment.image.repository | string | `"braveokafor/go-hello-world"` | Repository of the image |
| deployment.image.tag | string | `"latest"` | Tag of the image |
| deployment.livenessProbe | object | `{"initialDelaySeconds":30,"path":"/healthz","periodSeconds":10}` | Liveness probe configurations |
| deployment.replicaCount | int | `1` | Number of replicas for the deployment |
| deployment.resources | object | `{"limits":{"cpu":"128m","memory":"256Mi"},"requests":{"cpu":"128m","memory":"256Mi"}}` | Resource requests and limits |
| ingress | object | `{"annotations":{},"enabled":false,"hosts":[{"host":null,"paths":["/"]}],"ingressClassName":null,"tls":[]}` | Ingress configurations |
| ingress.annotations | object | `{}` | Additional annotations for the ingress |
| ingress.enabled | bool | `false` | Enable or disable ingress |
| ingress.ingressClassName | string | `nil` | Class name of the ingress |
| ingress.tls | list | `[]` | TLS configurations for the ingress |
| runtimeConfigs | object | `{"errorRate":0.15,"maxSleepMs":1050,"minSleepMs":0,"name":"Brave","serverPort":5000}` | Runtime configurations |
| runtimeConfigs.errorRate | float | `0.15` | Error rate as a float (e.g., 0.15 equals 15% errors) |
| runtimeConfigs.maxSleepMs | int | `1050` | Maximum sleep duration in milliseconds |
| runtimeConfigs.minSleepMs | int | `0` | Minimum sleep duration in milliseconds |
| runtimeConfigs.name | string | `"Brave"` | Name configuration |
| runtimeConfigs.serverPort | int | `5000` | Port on which the server will run |
| service | object | `{"annotations":{},"port":80,"type":"ClusterIP"}` | Service configurations |
| service.annotations | object | `{}` | Additional annotations for the service |
| service.port | int | `80` | Port on which the service will be exposed |
| service.type | string | `"ClusterIP"` | Type of the service (e.g., ClusterIP, NodePort, LoadBalancer) |

**Homepage:** <https://github.com/braveokafor/go-hello-world/>

## Source Code

* <https://github.com/braveokafor/go-hello-world/>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| Brave Okafor | <okaforbrave@gmail.com> | <https://www.linkedin.com/in/braveokafor/> |

