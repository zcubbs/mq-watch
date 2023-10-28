# MQ Watch

`mq-watch` is a topic subscriber for MQTT brokers.

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
  host: <string>              # MQTT broker host
  port: <int>                 # MQTT broker port  
  username: <string>          # MQTT broker username
  password: <string>          # MQTT broker password
  tls: <bool>                 # Use TLS (true/false)
  caCert: <string>            # Path to CA certificate
  clientCert: <string>        # Path to client certificate
  clientKey: <string>         # Path to client key
  clientID: <string>          # Client ID
  topic: <string>             # Topic to subscribe to

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
