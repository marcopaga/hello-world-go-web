# Simple Go Web-Server with helm chart

# Local testing with k3d
K3d creates a fully functional local kubernets cluster. Simply install https://k3d.io/ and you are good to go with `k3d cluster create`.

## Expose the ingress

If you want to use the ingress to expose services, which this example does, you need to activate it. Detailed info: https://k3d.io/v5.4.6/usage/exposing_services/

`k3d cluster create -p "8081:80@loadbalancer"`

## Installation of the helm chart
`helm upgrade hello-world-go-web hello-world-go-web -i --set imageCredentials.username="${GHCR_IO_USERNAME}" --set imageCredentials.password="${GHCR_IO_PASSWORD}"`

## Simple test run
`curl localhost:8081/factorial`

## 100 Parallel calls to make the cpus burn
`for i in {1..100}; do curl "http://localhost:8081/factorial" & done; wait`
