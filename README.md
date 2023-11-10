# Simple Go Web-Server with helm chart

## A simple kubernetes cluster locally with the ingress activated

`k3d cluster create -p "8081:80@loadbalancer"`

## Installation of the chart
`helm upgrade hello-world-go-web hello-world-go-web -i --set imageCredentials.username="${GHCR_IO_USERNAME}" --set imageCredentials.password="${GHCR_IO_PASSWORD}"`

## Simple test run
`curl localhost:8081/factorial`

## 100 Parallel calls to make the cpus burn
`for i in {1..100}; do curl "http://localhost:8081/factorial" & done; wait`
