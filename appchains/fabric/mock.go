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
	ChainID        string
	ChannelID      string
	PeerRPCAddrs   []string
	OrdererRPCAddr string
	Client         interface{}
	Key            interface{}
}

// NewFabricChain constructs a new FabricChain instance
func NewFabricChain(
	chainID string,
	channelID string,
	peerRPCAddrs []string,
	ordererRPCAddr string,
	client interface{},
	key interface{},
) FabricChain {
	return FabricChain{
		ChainID:        chainID,
		ChannelID:      channelID,
		PeerRPCAddrs:   peerRPCAddrs,
		OrdererRPCAddr: ordererRPCAddr,
		Client:         client,
		Key:            key,
	}
}

// GetChainID implements AppChainI
func (fc FabricChain) GetChainID() string {
	return fc.ChainID
}

// InterchainEventListener implements AppChainI
func (fc FabricChain) InterchainEventListener(cb func(icEvent core.InterchainEventI)) error {
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
