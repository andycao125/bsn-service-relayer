# BSN Service Relayer

BSN Service Relayer is intended for interchain service invocation between application chains and IRITA-HUB in BSN ecosystem.

The project is under development, with a mock implementation used to demonstrate the relayer working model and as a development reference.

## Install

Install the relayer using the following command:

```bash
make install
```

## Start

Get help information by:

```bash
relayer -h
```

Start the demo relayer:

```bash
relayer start ./config/config.yaml
```

## Development

See [Spec](./spec/design.md) for the application chain interfaces.

Add an application chain in `appchains` directory:

- Implementing the required interfaces
- Add to `appchains/factory.go` for routing
