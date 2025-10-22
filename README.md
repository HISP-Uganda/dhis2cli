# dhis2cli

A CLI tool for interacting with DHIS2 servers. It provides commands for common tasks like working with metadata, tracker, data values, users, scheduling, and more.

## Features
- Authenticate and connect to a DHIS2 instance
- Manage and query:
    - Analytics
    - Apps
    - Data, data values, and data entry
    - Datastore
    - GeoJSON
    - Maintenance operations
    - Metadata
    - Organisation units
    - Scheduling
    - SMS
    - Tracker
    - Users

## Prerequisites
- Go 1.24+
- Access to a DHIS2 server (URL, username, password or token)

// ... existing code ...
## Installation
- From source:
    - Clone this repository
    - Build: `go build -o dhis2cli ./...`
    - Or install: `go install ./...`

- Using the install script:
    - `sh install.sh`

- Using make (if available):
    - `make build`
      // ... existing code ...

## Configuration
- Create a config file based on config.yaml.sample and set:
    - baseUrl
    - authentication (username/password or token)
- You can also use environment variables or flags to override config values.

## Usage
- Show help:
    - `dhis2cli --help`
    - `dhis2cli <command> --help`

- Examples:
    - Use a config file:
        - `dhis2cli --config ~/.config/dhis2cli/config.yaml <command>`
        - `dhis2cli -c ~/.config/dhis2cli/config.yaml <command>`

## Development
- Run locally: `go run ./main.go --help`
- Lint/test: `go test ./...`
- Release builds may use goreleaser.

## Contributing
- Open issues and pull requests.
- Please include steps to reproduce and relevant logs when reporting bugs.
