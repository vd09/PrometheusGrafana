# How to Monitor Golang with Prometheus (Counter - Gauge - Histogram - Summary)

## Command to Run code
docker-compose up --build

## curl to increment histogram count
curl localhost:8080/devices
curl localhost:8081/metrics

## Grafana Dashboard link
http://localhost:3000/d/TqLbieBIk/my-app?orgId=1
user: admin
password: devops123

You can find tutorial [here](https://antonputra.com/monitoring/monitor-golang-with-prometheus/).
