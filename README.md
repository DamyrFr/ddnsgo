![GitHub Workflow Status](https://github.com/DamyrFr/ddnsgo/actions/workflows/go.yml/badge.svg)

# DDNSGo

DDNSGo is a simple Dynamic DNS program built in Go. It automatically updates your DNS records when your IP address changes.

## Features

- Lightweight and efficient
- Supports Cloudflare as a DNS provider
- Configurable via environment variables or YAML file
- Easy to set up and use

## Supported Providers

Currently, DDNSGo supports the following DNS providers:

- Cloudflare

We plan to add support for more providers in the future. Contributions are welcome!

## Configuration

### Cloudflare

You can configure DDNSGo using environment variables or a YAML config file.

#### YAML Configuration

Create a file named `ddnsgo.yaml` with the following structure:

```yaml
provider: cloudflare
cloudflare:
  api_token: XXXXXXX # API token (create in your Cloudflare profile)
  zone_id: BBBBBB    # The zone ID (found on the main page of your zone)
  record_id: RRRRR   # Get it via API (see below)
  record_name: DDDD  # Subdomain to update
```

#### Environment Variables

Alternatively, you can use the following environment variables:

```
DDNSGO_PROVIDER=cloudflare
DDNSGO_CLOUDFLARE_API_TOKEN=XXXXXXX
DDNSGO_CLOUDFLARE_ZONE_ID=BBBBBB
DDNSGO_CLOUDFLARE_RECORD_ID=RRRRR
DDNSGO_CLOUDFLARE_RECORD_NAME=DDDD
```

#### Getting the Record ID

To get the `record_id`, use the following curl command:

```bash
curl --request GET \
  --url https://api.cloudflare.com/client/v4/zones/YOUR_ZONE_ID/dns_records \
  --header 'Content-Type: application/json' \
  --header 'Authorization: Bearer YOUR_API_TOKEN' | jq
```

Replace `YOUR_ZONE_ID` and `YOUR_API_TOKEN` with your actual values.

## Installation

### Building from Source

To build DDNSGo from source, use the following command:

```bash
GOARCH=arm64 GOOS=linux CGO_ENABLED=0 go build .
```

Adjust `GOARCH` and `GOOS` as needed for your target system.

### Systemd Service

To run DDNSGo as a systemd service, create a file named `/etc/systemd/system/ddnsgo.service` with the following content:

```ini
[Unit]
Description=Dynamic DNS Client
After=network.target

[Service]
ExecStart=/usr/local/bin/ddnsgo
Restart=always
RestartSec=60
User=nobody
Group=nogroup
Environment="DDNSGO_CONFIG=/etc/ddnsgo.yaml"

[Install]
WantedBy=multi-user.target
```

Then enable and start the service:

```bash
sudo systemctl enable ddnsgo
sudo systemctl start ddnsgo
```

## Usage

Once configured and running, DDNSGo will periodically check your public IP address and update your DNS record if it has changed.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

GNU General Public License v3.0
