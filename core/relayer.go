package core

import (
	log "github.com/sirupsen/logrus"
)

// Relayer is a relayer implementing the interchain service invocation between the application chain and Irita-Hub
type Relayer struct {
	IritaHubChain IritaHubChainI
	AppChain      AppChainI
	Logger        *log.Logger
}

// NewRelayer constructs a new Relayer instance
func NewRelayer(iritaHubChain IritaHubChainI, appChain AppChainI, logger *log.Logger) Relayer {
	return Relayer{
		IritaHubChain: iritaHubChain,
		AppChain:      appChain,
		Logger:        logger,
	}
}

// BuildInterchainRequest builds an interchain request from an interchain event
func (r Relayer) BuildInterchainRequest(icEvent InterchainEventI) (InterchainRequestI, error) {
	serviceFee, err := r.AppChain.GetServiceFee(icEvent.GetServiceName(), icEvent.GetProvider())
	if err != nil {
		return nil, err
	}

	return InterchainRequestAdaptor{
		ServiceName:   icEvent.GetServiceName(),
		Provider:      icEvent.GetProvider(),
		Input:         icEvent.GetInput(),
		Timeout:       icEvent.GetTimeout(),
		ServiceFeeCap: serviceFee,
	}, nil
}

// HandleInterchainEvent handles the interchain event
func (r Relayer) HandleInterchainEvent(icEvent InterchainEventI) {
	r.Logger.Printf("got the interchain event on %s: %+v", r.AppChain.GetChainID(), icEvent)

	icRequest, err := r.BuildInterchainRequest(icEvent)
	if err != nil {
		r.Logger.Printf(
			"failed to build the interchain request: %s",
			err,
		)

		return
	}

	callback := func(icRequestID string, response ResponseI) {
		r.Logger.Printf(
			"got the response of the interchain request %s on %s: %+v",
			icRequestID,
			r.IritaHubChain.GetChainID(),
			response,
		)

		err := r.AppChain.SendResponse(icEvent.GetInvocationID(), response)
		if err != nil {
			r.Logger.Printf(
				"failed to send the response to %s: %s",
				r.AppChain.GetChainID(),
				err,
			)
		} else {
			r.Logger.Printf(
				"response sent to %s successfully",
				r.AppChain.GetChainID(),
			)
		}
	}

	err = r.IritaHubChain.SendInterchainRequest(icRequest, callback)
	if err != nil {
		r.Logger.Printf(
			"failed to send the interchain request %+v to %s: %s",
			icRequest,
			r.IritaHubChain.GetChainID(),
			err,
		)
	}
}

// Start starts the relayer process
func (r Relayer) Start() {
	err := r.AppChain.InterchainEventListener(r.HandleInterchainEvent)
	if err != nil {
		r.Logger.Printf(
			"failed to listen to the interchain event on %s",
			r.AppChain.GetChainID(),
		)
	}
}
