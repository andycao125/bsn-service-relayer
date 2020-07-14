package core

// ChainI defines the basic chain interface
type ChainI interface {
	GetChainID() string // chain ID getter
}

// IritaHubChainI defines the interface to interact with Irita-Hub
type IritaHubChainI interface {
	ChainI
	SendInterchainRequest(request InterchainRequestI) (icRequestID string, err error)
	ResponseListener(icRequestID string) (ResponseI, error)
}

// AppChainI defines the interface to interact with the application chain
type AppChainI interface {
	ChainI
	InterchainEventListener(cb func(icEvent InterchainEventI)) error
	SendResponse(invocationID string, response ResponseI) error
}

// InterchainEventI abstracts the event which is triggered for an interchain service invocation
type InterchainEventI interface {
	GetInvocationID() string // invocation ID getter
	GetServiceName() string  // service name getter
	GetInput() string        // request input getter
	GetTimeout() uint64      // request timeout getter
}

// InterchainRequestI is an interface for the interchain request
type InterchainRequestI interface {
	GetServiceName() string   // service name getter
	GetProvider() string      // provider getter
	GetInput() string         // request input getter
	GetTimeout() uint64       // request timeout getter
	GetServiceFeeCap() string // service fee cap getter
}

// ResponseI exposes the response related interfaces
type ResponseI interface {
	GetInterchainRequestID() string // interchain request ID getter
	GetResponse() string            // service response getter
	GetError() string               // error msg getter
}
