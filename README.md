# duckdns-updater

A simple [Duck DNS](http://www.duckdns.org) update client

## Installation

```
$ go get github.com/davewongillies/duckdns-updater
```

## Usage

```
$ env DUCKDNS_TOKEN=692d9c18-7f63-4759-b78d-e873ee7e78fc DUCKDNS_SUBDOMAIN=example duckdns-updater
2019/04/04 20:49:08 Updating example.duckdns.org
2019/04/04 20:49:08 Successfully updated example.duckdns.org
```

Or using docker:

```
docker run -e DUCKDNS_REFRESH_INTERVAL=60 -e DUCKDNS_TOKEN=692d9c18-7f63-4759-b78d-e873ee7e78fc -e DUCKDNS_SUBDOMAIN=example -it --rm davewongillies/duckdns-updater:latest
```

## Variables

`duckdns` takes the following environment variables:

|Environment Variable|Description|Example|
|--------------------|-----------|-------|
|`DUCKDNS_REFRESH_INTERVAL`|(Optional) An integer of seconds that duckdns-updater will update your DNS records again |`DUCKDNS_REFRESH_INTERVAL=60`|
|`DUCKDNS_SUBDOMAIN`|(Required) DuckDNS subdomain you want to update|`DUCKDNS_SUBDOMAIN=example`|
|`DUCKDNS_TOKEN`|(Required) Duck DNS auth token|`DUCKDNS_TOKEN=692d9c18-7f63-4759-b78d-e873ee7e78fc`|
