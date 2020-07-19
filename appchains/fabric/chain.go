package fabric

import (
	"fmt"
	"time"

	"relayer/core"
)

// InterchainEvent is the Fabric event which represents an interchain service invocation to Irita-Hub
type InterchainEvent struct {
	InvocationID string
	ServiceName  string
	Provider     string
	Input        string
	Timeout      uint64
}

// GetInvocationID implements InterchainEventI
func (e InterchainEvent) GetInvocationID() string {
	return e.InvocationID
}

// GetServiceName implements InterchainEventI
func (e InterchainEvent) GetServiceName() string {
	return e.ServiceName
}

// GetProvider implements InterchainEventI
func (e InterchainEvent) GetProvider() string {
	return e.Provider
}

// GetInput implements InterchainEventI
func (e InterchainEvent) GetInput() string {
	return e.Input
}

// GetTimeout implements InterchainEventI
func (e InterchainEvent) GetTimeout() uint64 {
	return e.Timeout
}

// FabricChain defines the Fabric chain
type FabricChain struct {
	ChannelID      string
	PeerRPCAddrs   []string
	OrdererRPCAddr string
	Client         interface{}
	Key            string
	Passphrase     string
}

// NewFabricChain constructs a new FabricChain instance
func NewFabricChain(
	channelID string,
	peerRPCAddrs []string,
	ordererRPCAddr string,
	key string,
	passphrase string,
) FabricChain {
	fabric := FabricChain{
		ChannelID:      channelID,
		PeerRPCAddrs:   peerRPCAddrs,
		OrdererRPCAddr: ordererRPCAddr,
		Key:            key,
	}

	fabric.Client = nil

	return fabric
}

// MakeFabricChain builds a fabric chain from the given config
func MakeFabricChain(config Config) FabricChain {
	return NewFabricChain(
		config.ChannelID,
		config.PeerRPCAddrs,
		config.OrdererRPCAddr,
		config.Key,
		config.Passphrase,
	)
}

// GetChainID implements AppChainI
func (fc FabricChain) GetChainID() string {
	return fc.ChannelID
}

// InterchainEventListener implements AppChainI
func (fc FabricChain) InterchainEventListener(cb core.InterchainEventHandler) error {
	i := 0

	for {
		icEvent := InterchainEvent{
			InvocationID: fmt.Sprintf("invocation%d", i+1),
			ServiceName:  "exchange_rate",
			Input:        fmt.Sprintf(`{"name":"CNY-USD"}`),
			Timeout:      uint64(50),
		}

		cb(icEvent)

		time.Sleep(10 * time.Second)
		i++
	}
}

// SendResponse implements AppChainI
func (fc FabricChain) SendResponse(invocationID string, response core.ResponseI) error {
	return nil
}

// AddServiceBinding implements AppChainI
func (fc FabricChain) AddServiceBinding(serviceName, schemas, provider, serviceFee string, qos uint64) error {
	return nil
}

// GetServiceFee implements AppChainI
func (fc FabricChain) GetServiceFee(serviceName string, provider string) (fee string, err error) {
	return "100point", nil
}
