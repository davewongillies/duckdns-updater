# duckdns-updater

[![Docker Hub Pulls](https://img.shields.io/docker/pulls/davewongillies/duckdns-updater.svg)](https://hub.docker.com/r/davewongillies/duckdns-updater/)

A simple [Duck DNS](http://www.duckdns.org) update client

## Installation

```
$ go get github.com/davewongillies/duckdns-updater
```

Or pull the docker image:

```
$ docker pull davewongillies/duckdns-updater
```

## Usage

```
$ env DUCKDNS_TOKEN=<your_token> DUCKDNS_SUBDOMAIN=example duckdns-updater
2019/04/04 20:49:08 Updating example.duckdns.org
2019/04/04 20:49:08 Successfully updated example.duckdns.org
```

Or using docker:

```
$ docker run -e DUCKDNS_REFRESH_INTERVAL=60 -e DUCKDNS_TOKEN=<your_token> -e DUCKDNS_SUBDOMAIN=example -it --rm davewongillies/duckdns-updater:latest
```

Or `docker-compose`:

```yaml
---
version: '3'
services:
  duckdns:
    image: davewongillies/duckdns-updater
    container_name: duckdns
    environment:
      - DUCKDNS_REFRESH_INTERVAL=600
      - DUCKDNS_TOKEN=<your_token>
      - DUCKDNS_SUBDOMAIN=example
    restart: unless-stopped
```

## Variables

`duckdns-updater` takes the following environment variables:

|Environment Variable|Description|Example|
|--------------------|-----------|-------|
|`DUCKDNS_REFRESH_INTERVAL`|(Optional) An integer of seconds that duckdns-updater will update your DNS records again |`DUCKDNS_REFRESH_INTERVAL=60`|
|`DUCKDNS_SUBDOMAIN`|(Required) DuckDNS subdomain you want to update|`DUCKDNS_SUBDOMAIN=example`|
|`DUCKDNS_TOKEN`|(Required) Duck DNS auth token||
