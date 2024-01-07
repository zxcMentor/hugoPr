FROM grafana/grafana

COPY datasources.yaml /etc/grafana/provisioning/datasources/
COPY dashboards.json /var/lib/grafana/dashboards/