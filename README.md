# golang-ethtool

Provides packages for interacting with the Linux kernel's ethtool ioctls.
Supports parsing transceiver's EEPROM data according to the standards:
* [SFF-8472](https://members.snia.org/document/dl/25916) rev 12.3
* [SFF-8636](https://members.snia.org/document/dl/26418) rev 4.9
* SFF-8463

## Overview
* `eeprom/eeprom.go` provides a unified interface for different EEPROM types.
* `eeprom/sff8472/eeprom.go` provides the SFF-8472 implementation
* `eeprom/sff8636/eeprom.go` provides the SFF-8636 implementation, which is also used for decoding SFF8463 eeproms.

## Usage
### Included basic example
A minimal example is included:
```bash
go get ./...
cd example
go build
```
It allows for parsing the EEPROM contents of a given <interface> and dumps them to STDOUT: `./example --interface swp42`

### Transceiver exporter
[Prometheus exporter](https://github.com/wobcom/transceiver-exporter) based on this package.

## Authors
* @fluepke
* @vidister
