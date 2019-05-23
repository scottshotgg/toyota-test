# toyota-test

This is an example of a simple, but well documented micro-service that includes a REST API entirely built from a Swagger.

## To run:
The easiest way to run the program is to run the following command from the root directory:
`GO111MODULE=on bin/scratch.sh`

In order to deploy this app onto your Kubernetes cluster, or to run in a Docker container locally, run
the following commands from the root folder:
1) `docker build -t toyota-test .`
2) `docker run -p 8080:8080 -it toyota-test:latest`
