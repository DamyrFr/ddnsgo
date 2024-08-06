# DDNSGo

DDNSGo is a simple Dynamic DNS programm. It's build in Golang.

# Providers

For now I only support Cloudflare. If I have time for that, I will add somme others.

## Cloudflare

You can configure it using environement variables or by the following config file on 

```
provider: cloudflare
cloudflare:
  api_token: XXXXXXX # API token create it on you profile
  zone_id: BBBBBB # The zone ID on the main page of zone
  record_id: RRRRR # Get it bu API
  record_name: DDDD # Subdomain to use
```

Get the record_id by API (using curl) :

```
curl --request GET  --url https://api.cloudflare.com/client/v4/zones/YOUR_ZONE_ID/dns_records --header 'Content-Type: application/json' --header 'Authorization: Bearer YOUR_API_KEY' | jq
```
