package hub

import (
	iritaclient "github.com/bianjieai/irita-ent-sdk-go"
	iservice "github.com/bianjieai/irita-ent-sdk-go/modules/service"
	iritasdk "github.com/bianjieai/irita-ent-sdk-go/types"

	"relayer/core"
)

// IritaHubChain defines the Irita-Hub chain
type IritaHubChain struct {
	ChainID    string
	RPCAddr    string
	Client     iritaclient.IRITAClient
	Key        string
	Passphrase string
}

// NewIritaHubChain constructs a new Irita-Hub chain
func NewIritaHubChain(
	chainID string,
	rpcAddr string,
	key string,
	passphrase string,
) IritaHubChain {
	hub := IritaHubChain{
		ChainID:    chainID,
		RPCAddr:    rpcAddr,
		Key:        key,
		Passphrase: passphrase,
	}

	config := defaultClientConfig
	if len(rpcAddr) != 0 {
		config.NodeURI = rpcAddr
	}

	hub.Client = iritaclient.NewIRITAClient(config)

	return hub
}

// MakeIritaHubChain builds an Irita-Hub from the given config
func MakeIritaHubChain(config Config) IritaHubChain {
	return NewIritaHubChain(config.ChainID, config.NodeRPCAddr, config.Key, config.Passphrase)
}

// GetChainID implements IritaHubChainI
func (ic IritaHubChain) GetChainID() string {
	return ic.ChainID
}

// SendInterchainRequest implements IritaHubChainI
func (ic IritaHubChain) SendInterchainRequest(
	request core.InterchainRequestI,
	cb core.ResponseCallback,
) error {
	invokeServiceReq, err := ic.BuildServiceInvocationRequest(request)
	if err != nil {
		return err
	}

	reqCtxID, err := ic.Client.Service.InvokeService(invokeServiceReq, ic.BuildBaseTx())
	if err != nil {
		return err
	}

	return ic.ResponseListener(reqCtxID, cb)
}

// BuildServiceInvocationRequest builds the service invocation request from the given interchain request
func (ic IritaHubChain) BuildServiceInvocationRequest(
	request core.InterchainRequestI,
) (iservice.SvcInvocationRequest, error) {
	serviceFeeCap, err := iritasdk.ParseDecCoins(request.GetServiceFeeCap())
	if err != nil {
		return iservice.SvcInvocationRequest{}, err
	}

	return iservice.SvcInvocationRequest{
		ServiceName:   request.GetServiceName(),
		Providers:     []string{request.GetProvider()},
		Input:         request.GetInput(),
		Timeout:       int64(request.GetTimeout()),
		ServiceFeeCap: serviceFeeCap,
	}, nil
}

// ResponseListener gets and handles the response of the given request context ID by event subscription
func (ic IritaHubChain) ResponseListener(reqCtxID string, cb core.ResponseCallback) error {
	callbackWrapper := func(reqCtxID, requestID, response string) {
		resp := core.ResponseAdaptor{
			StatusCode:  200,
			Output:      response,
			ICRequestID: requestID,
		}

		cb(requestID, resp)
	}

	_, err := ic.Client.Service.SubscribeServiceResponse(reqCtxID, callbackWrapper)
	if err != nil {
		return err
	}

	return nil
}

// BuildBaseTx builds a base tx
func (ic IritaHubChain) BuildBaseTx() iritasdk.BaseTx {
	return iritasdk.BaseTx{
		From:     ic.Key,
		Password: ic.Passphrase,
	}
}
