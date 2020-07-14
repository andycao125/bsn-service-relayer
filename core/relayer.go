package core

import (
	"log"
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
	return InterchainRequestAdaptor{
		ServiceName:   icEvent.GetServiceName(),
		Provider:      "mock-provider",
		Input:         icEvent.GetInput(),
		Timeout:       icEvent.GetTimeout(),
		ServiceFeeCap: "mock-service-fee-cap",
	}, nil
}

// HandleInterchainEvent handles the interchain event
func (r Relayer) HandleInterchainEvent(icEvent InterchainEventI) {
	response := ResponseAdaptor{}

	r.Logger.Printf("got the interchain event on %s: %v", r.AppChain.GetChainID(), icEvent)

	icRequest, err := r.BuildInterchainRequest(icEvent)
	if err != nil {
		r.Logger.Printf(
			"failed to build the interchain request from the interchain event: %s",
			err,
		)

		response.ErrMsg = "invalid interchain event"
	} else {
		icRequestID, err := r.IritaHubChain.SendInterchainRequest(icRequest)
		if err != nil {
			r.Logger.Printf(
				"failed to send the interchain request %v to %s: %s",
				icRequest,
				r.IritaHubChain.GetChainID(),
				err,
			)

			response.ErrMsg = "failed to send the interchain request to Hub"
		} else {
			r.Logger.Printf(
				"interchain request %v sent to %s successfully: %s",
				icRequest,
				r.IritaHubChain.GetChainID(),
				icRequestID,
			)

			resp, err := r.IritaHubChain.ResponseListener(icRequestID)
			if err != nil {
				r.Logger.Printf(
					"failed to get the response of the interchain request %s on %s: %s",
					icRequestID,
					r.IritaHubChain.GetChainID(),
					err,
				)

				response.ErrMsg = "failed to receive the response from Hub"
			} else {
				r.Logger.Printf(
					"got the response of the interchain request %s on %s: %v",
					icRequestID,
					r.IritaHubChain.GetChainID(),
					resp,
				)

				response = resp.(ResponseAdaptor)
			}
		}
	}

	err = r.AppChain.SendResponse(icEvent.GetInvocationID(), response)
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
