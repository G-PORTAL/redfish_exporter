# Redfish Exporter

Redfish Exporter for Prometheus. This exporter is using [Redfish API](https://www.dmtf.org/standards/redfish) to
fetch metrics from servers. 

## Warning

This is a work in progress.

## Usage

### With Command Line

```
Usage of redfish_exporter:
  -config.file string
        Path to the configuration file (default "config.yml")
  -web.listen-address string
        Address to listen on (default "0.0.0.0:9096")
```

### With Docker

```shell
docker run -d -p 0.0.0.0:9096:9096 \
  -v "$(pwd)/config.example.yml:/etc/redfish_exporter/config.yml" \
  --name redfish-exporter gportal/redfish_exporter:latest
```

## Supports

Redfish Exporter is providing generic metrics fetched from Redfish endpoints.

### Tested with
* HPE ILO4
* HPE ILO5
* Dell iDRAC 8
* Dell iDRAC 9
* AsRockRack
* SuperMicro SuperServer
