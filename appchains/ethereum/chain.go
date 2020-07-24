package ethereum

import (
	"time"

	"relayer/core"
)

// ServiceRequestEvent is the Ethereum event which represents an interchain service invocation to Irita-Hub
type iServiceRequestEvent struct {
	// TODO
}

// GetInvocationID implements InterchainEventI
func (e iServiceRequestEvent) GetInvocationID() string {
	return "invocation"
}

// GetServiceName implements InterchainEventI
func (e iServiceRequestEvent) GetServiceName() string {
	return "service-name"
}

// GetProvider implements InterchainEventI
func (e iServiceRequestEvent) GetProvider() string {
	return "provider"
}

// GetInput implements InterchainEventI
func (e iServiceRequestEvent) GetInput() string {
	return "input"
}

// GetTimeout implements InterchainEventI
func (e iServiceRequestEvent) GetTimeout() uint64 {
	return 0
}

// EthChain defines the Ethereum chain
type EthChain struct {
	ChainID                string
	NodeRPCAddr            string
	IServiceCoreContract   string
	IServiceMarketContract string
	Client                 interface{}
	Key                    string
	Passphrase             string
}

// NewEthChain constructs a new EthChain instance
func NewEthChain(
	chainID string,
	nodeRPCAddr string,
	iServiceCoreContract string,
	iServiceMarketContract string,
	key string,
	passphrase string,
) EthChain {
	eth := EthChain{
		ChainID:                chainID,
		NodeRPCAddr:            nodeRPCAddr,
		IServiceCoreContract:   iServiceCoreContract,
		IServiceMarketContract: iServiceMarketContract,
		Key:                    key,
		Passphrase:             passphrase,
	}

	eth.Client = nil

	return eth
}

// MakeEthChain builds an Ethereum chain from the given config
func MakeEthChain(config Config) EthChain {
	return NewEthChain(
		config.ChainID,
		config.NodeRPCAddr,
		config.IServiceCoreContract,
		config.IServiceMarketContract,
		config.Key,
		config.Passphrase,
	)
}

// GetChainID implements AppChainI
func (ec EthChain) GetChainID() string {
	return ec.ChainID
}

// InterchainEventListener implements AppChainI
func (ec EthChain) InterchainEventListener(cb core.InterchainEventHandler) error {
	i := 0

	for {
		icEvent := iServiceRequestEvent{}

		cb(icEvent)

		time.Sleep(10 * time.Second)
		i++
	}
}

// SendResponse implements AppChainI
func (ec EthChain) SendResponse(invocationID string, response core.ResponseI) error {
	return nil
}

// AddServiceBinding implements AppChainI
func (ec EthChain) AddServiceBinding(serviceName, schemas, provider, serviceFee string, qos uint64) error {
	return nil
}

// GetServiceFee implements AppChainI
func (ec EthChain) GetServiceFee(serviceName string, provider string) (fee string, err error) {
	return "100point", nil
}
