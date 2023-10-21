### Go Hello World

[![Build Status][badge_build_status]][link_build_status]
[![Release Status][badge_release_status]][link_build_status]
[![Repo size][badge_repo_size]][link_repo]
[![Image size][badge_size_latest]][link_docker_hub]
[![Docker Pulls][badge_docker_pulls]][link_docker_hub]

### Overview

Introducing the `braveokafor/go-hello-world` Docker image.  
This Go application extends beyond the conventional greeting, integrating Prometheus metrics to provide insights into the operational dynamics of HTTP servers under various conditions.

![Grafana Dashboard](dashboards/dashboard.gif)

### Features
- **HTTP Greetings:** The application issues a plain text greeting at the root endpoint.
- **Prometheus Metrics:** With Prometheus metrics integration, it highlights the total count of HTTP requests, their duration, and the number of in-flight requests.
- **Configurable Delays & Errors:** Users have the flexibility to introduce artificial delays and errors into the responses for experimental purposes.
- **Environment-Driven Configuration:** The application’s behaviour is customisable through both environment variables and command-line flags.

### Endpoints Available

- `/` - Greeting message endpoint.
- `/metrics` - Endpoint for accessing Prometheus metrics.
- `/healthz` - Endpoint providing the application’s health status.

### Docker-Compose Deployment & Grafana Dashboards

A [`docker-compose.yaml`](https://github.com/braveokafor/go-hello-world/blob/main/docker-compose.yaml) file is available in the project repository to facilitate the deployment of the application alongside the necessary monitoring tools.  
This configuration allows for an integrated environment where the application’s metrics can be easily visualised and analysed.

Example Grafana dashboards are also provided in the repository, offering a starting point for monitoring the application’s performance and behaviour.


### Helm Chart

A [helm chart](https://github.com/braveokafor/go-hello-world/blob/main/charts/go-hello-world) is available in the project repository for deploying the application on Kubernetes.  
A [README](https://github.com/braveokafor/go-hello-world/blob/main/charts/go-hello-world/README.md) is also available in `charts/go-hello-world/`.


## Installation
Download the appropriate binary for your operating system from the [GitHub Releases page](https://github.com/braveokafor/go-hello-world/releases).

### Usage
Run the `go-hello-world`` binary with desired flags, or set the environment variables before starting the utility.

## Options

| Flag	        | Environment Variable | Description	                                     | Default Value | 
|---------------|----------------------|-----------------------------------------------------|---------------|
| `-port`	    | `SERVER_PORT`	       | Port on which the server runs	                     | `5000`        | 
| `-min-sleep`	| `MIN_SLEEP_MS`       | Min sleep duration in milliseconds to simulate work | `0`           |
| `-max-sleep`	| `MAX_SLEEP_MS`	   | Max sleep duration in milliseconds to simulate work | `0`           |
| `-error-rate`	| `ERROR_RATE`	       | Error simulation rate	                             | `0.0`         |
| `-name`	    | `NAME`	           | Name to be used in greeting	                     | `Brave`       |

## Examples:

### CLI:
```sh
# With flags
go-hello-world --error-rate 0.15 --max-sleep 1050 --min-sleep 0 --name Brave --port 5000

# With environment variables
export SERVER_PORT="5000"
export MAX_SLEEP_MS="1050"
export MIN_SLEEP_MS="0"
export ERROR_RATE="0.15"
export NAME="Brave"
go-hello-world
```

### Docker
```sh
# With flags
docker run braveokafor/go-hello-world:latest --error-rate 0.15 --max-sleep 1050 --min-sleep 0 --name Brave --port 5000

# With environment variables
docker run -p 5000:5000 -e SERVER_PORT="5000" braveokafor/go-hello-world:latest
```

### Contribution and Support

[![Issues][badge_issues]][link_issues]
[![Issues][badge_pulls]][link_pulls]

While primarily a personal project for learning and usage, suggestions and identified issues are welcomed.  
Maintenance may not be continuous, but feedback will be reviewed and addressed as time allows.

### Contact & Repository
For additional information, suggestions, or enquiries, please visit the project on [GitHub](https://github.com/braveokafor/go-hello-world) or connect with me on [LinkedIn](https://www.linkedin.com/in/braveokafor/).

### Licence

The project is open-sourced under the MIT Licence, allowing for use, modification, and distribution of the code under the licence’s terms and conditions.

For detailed understanding, refer to the [LICENCE](https://github.com/braveokafor/go-hello-world/blob/main/LICENSE) file in the project repository.


[link_issues]:https://github.com/braveokafor/go-hello-world/issues
[link_pulls]:https://github.com/braveokafor/go-hello-world/pulls
[link_build_status]:https://github.com/braveokafor/go-hello-world/actions/workflows/go.yaml
[link_build_status]:https://github.com/braveokafor/go-hello-world/actions/workflows/release.yaml
[link_docker_hub]:https://hub.docker.com/r/braveokafor/go-hello-world
[link_repo]:https://github.com/braveokafor/go-hello-world

[badge_issues]:https://img.shields.io/github/issues-raw/braveokafor/go-hello-world?style=flat-square&logo=GitHub
[badge_pulls]:https://img.shields.io/github/issues-pr/braveokafor/go-hello-world?style=flat-square&logo=GitHub
[badge_build_status]:https://img.shields.io/github/actions/workflow/status/braveokafor/go-hello-world/go-ci.yaml?style=flat-square&logo=GitHub&label=build
[badge_release_status]:https://img.shields.io/github/actions/workflow/status/braveokafor/go-hello-world/go-release.yaml?style=flat-square&logo=GitHub&label=release
[badge_size_latest]:https://img.shields.io/docker/image-size/braveokafor/go-hello-world/latest?style=flat-square&logo=Docker
[badge_docker_pulls]:https://img.shields.io/docker/pulls/braveokafor/go-hello-world?style=flat-square&logo=Docker
[badge_repo_size]:https://img.shields.io/github/repo-size/braveokafor/go-hello-world?style=flat-square&logo=GitHub