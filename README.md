# BSN Service Relayer

BSN Service Relayer is intended for interchain service invocation between application chains and IRITA-HUB in BSN ecosystem.

The project is under development, with a mock implementation which is used to demonstrate the relayer working model and as a development reference.

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
relayer fabric
```

## Development

See [Spec](./spec/design.md) for interfaces.

Implement an application chain in `appchains` directory:

- Implement the `InterchainEventI` and `AppChainI`

- Implement appchain-specific commands
