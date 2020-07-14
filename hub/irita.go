package hub

import (
	"fmt"
	"math/rand"

	"relayer/core"
)

// IritaHubChain defines the Irita-Hub chain
type IritaHubChain struct {
	ChainID string
	RPCAddr string
	Client  interface{}
	Key     interface{}
}

// NewIritaHubChain constructs a new Irita-Hub chain
func NewIritaHubChain(
	chainID string,
	rpcAddr string,
	client interface{},
	key interface{},
) IritaHubChain {
	return IritaHubChain{
		ChainID: chainID,
		RPCAddr: rpcAddr,
		Client:  client,
		Key:     key,
	}
}

// GetChainID implements IritaHubChainI
func (ic IritaHubChain) GetChainID() string {
	return ic.ChainID
}

// SendInterchainRequest implements IritaHubChainI
func (ic IritaHubChain) SendInterchainRequest(request core.InterchainRequestI) (icRequestID string, err error) {
	return fmt.Sprintf("interchain-request-id-%d", rand.Int()), nil
}

// ResponseListener implements IritaHubChainI
func (ic IritaHubChain) ResponseListener(icRequestID string) (core.ResponseI, error) {
	response := core.ResponseAdaptor{
		StatusCode:  200,
		Result:      `{"code":"200","message":""}`,
		Output:      fmt.Sprintf(`{"value":"%d"}`, rand.Intn(10)),
		ICRequestID: icRequestID,
		ErrMsg:      "",
	}

	return response, nil
}
