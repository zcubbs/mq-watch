# MQ Watch

`mq-watch` is a topic subscriber for MQTT brokers.

---
<p align="center">
</p>
<p align="center">
  <img width="750" src="docs/assets/v2.png">
</p>

---

## Installation

> Supported Platforms: `linux_amd64/linux_arm64`.

### From script

```bash
curl -sfL https://raw.githubusercontent.com/zcubbs/mq-watch/main/scripts/install.sh | sudo bash 
```

### From Binary

You can download the latest release from [here](https://github.com/zcubbs/mq-watch/releases)
```bash
mq-watch -config /path/to/config.yaml
```

### Using Docker

```bash
docker run -d \
    -p 8000:8000 \
    -v /path/to/config.yaml:/app/config.yaml \
    ghcr.io/zcubbs/mq-watch:latest
```

### Using Helm

```bash
helm install mq-watch oci://ghcr.io/zcubbs/mq-watch/mq-watch -f /path/to/values.yaml
```

see [values.yaml](charts/mq-watch/values.yaml) for the default values.

## Configuration

mq-watch is configured via a YAML file you can provide to the container/binary. The example configuration is located at [config.yaml](./examples/config.yaml). The following is an example configuration:

```yaml
mqtt:
  broker: "mqtt://127.0.0.1:1883" # or "mqtts://
  client_id: "my_client_id" # client id to use when connecting to the broker

tenants:
  - name: "tenant1" # tenant name
    topics: # list of topics to subscribe to
      - "tenant1/#" # topic filter
  - name: "tenant2"
    topics:
      - "tenant2/#"

database:
  dialect: "sqlite" # or "postgres"
  datasource: "mq-watch.db" # or "postgres://user:password@host:port/dbname?sslmode=disable"
  auto_migrate: true # automatically migrate the database schema on startup

server:
  port: 8000 # port to listen on

```

## Development

### Prerequisites

- [Go](https://golang.org/doc/install)
- [Task](https://taskfile.dev/#/installation)

### Run Locally

```bash
task run
```

## Roadmap

### Messages Over Time
- **Description**: Visualize the growth or decline in message traffic over time.
- **Visualization**: Line graph showing the trend of total messages per day/week/month.

### Average Messages per Tenant
- **Description**: Display the average number of messages sent by each tenant.
- **Visualization**: Bar chart comparing average messages across tenants.

### Most Active Hours
- **Description**: Identify peak periods of messaging activity.
- **Visualization**: Heatmap or bar chart showing message volume by hour or time of day.

### Message Size Statistics
- **Description**: Provide insights into the sizes of messages being sent.
- **Visualization**: Statistics display for average, minimum, maximum, and total message sizes.

### Failed Messages
- **Description**: Track the rate and count of message delivery failures.
- **Visualization**: Counter or percentage display indicating failure rates.

### System Health
- **Description**: Monitor the health of the messaging system's infrastructure.
- **Visualization**: Dashboard indicators for server CPU, memory usage, and error rates.

### Tenant Engagement
- **Description**: Gauge the engagement level of each tenant.
- **Visualization**: Metric based on the number of active days or message interactions.

### New Tenants
- **Description**: Keep track of the growth in the number of tenants.
- **Visualization**: Incremental counter for new tenants added within a date range.

### Response Time
- **Description**: Measure the responsiveness of the messaging system.
- **Visualization**: Display for the average response time for message processing.

### Message Type Breakdown
- **Description**: Differentiate between the various types of messages sent.
- **Visualization**: Pie chart or bar graph breaking down messages by type.

### Tenant Comparison
- **Description**: Compare messaging statistics between two tenants.
- **Visualization**: Side-by-side comparison of key messaging stats.


## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

mq-watch is licensed under the [MIT](./LICENSE) license.
