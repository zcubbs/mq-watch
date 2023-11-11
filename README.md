# MQ Watch

`mq-watch` is a topic subscriber for MQTT brokers.

---
<p align="center">
</p>
<p align="center">
  <img width="550" src="docs/assets/current.png">
</p>

---

## Installation

> Supported Platforms: `linux_amd64/linux_arm64`.

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

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

mq-watch is licensed under the [MIT](./LICENSE) license.
