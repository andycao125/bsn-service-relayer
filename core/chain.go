package core

// ChainI defines the basic chain interface
type ChainI interface {
	GetChainID() string // chain ID getter
}

// IritaHubChainI defines the interface to interact with Irita-Hub
type IritaHubChainI interface {
	ChainI

	// send the interchain request and handle the response with the given callback
	SendInterchainRequest(request InterchainRequestI, cb ResponseCallback) error
}

// AppChainI defines the interface to interact with the application chain
type AppChainI interface {
	ChainI

	// listen to the interchain events with an event handler
	InterchainEventListener(cb InterchainEventHandler) error

	// send the response to the application chain
	SendResponse(invocationID string, response ResponseI) error

	// iService market interface
	IServiceMarketI
}

// InterchainEventI abstracts the event which is triggered for an interchain service invocation
type InterchainEventI interface {
	GetInvocationID() string // invocation ID getter
	GetServiceName() string  // service name getter
	GetProvider() string     // provider getter
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

// KeyManager defines the key management interface
type KeyManager interface {
	Add(name, passphrase string) (addr string, mnemonic string, err error)
	Delete(name string) error
	Get(name string) (addr string, err error)
	Import(name, passphrase, keystore string) error
	Export(name, passphrase, encryptPassphrase string) (keystore string, err error)
}

// IServiceMarketI defines the interface for the iService market on the application chain
type IServiceMarketI interface {
	// add a service binding to the iService market
	AddServiceBinding(serviceName, schemas, provider, serviceFee string, qos uint64) error

	// get the service fee by the given service name and provider
	GetServiceFee(serviceName string, provider string) (fee string, err error)
}

// InterchainEventHandler defines the interchain event callback interface
type InterchainEventHandler func(icEvent InterchainEventI)

// ResponseCallback defines the response callback interface
type ResponseCallback func(icRequestID string, response ResponseI)
