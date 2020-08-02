package core

// InterchainRequestAdaptor bridges the interchain event and Irita-Hub interchain request
type InterchainRequestAdaptor struct {
	ServiceName   string
	Provider      string
	Input         string
	Timeout       uint64
	ServiceFeeCap string
}

// GetServiceName implements InterchainRequestI
func (ir InterchainRequestAdaptor) GetServiceName() string {
	return ir.ServiceName
}

// GetProvider implements InterchainRequestI
func (ir InterchainRequestAdaptor) GetProvider() string {
	return ir.Provider
}

// GetInput implements InterchainRequestI
func (ir InterchainRequestAdaptor) GetInput() string {
	return ir.Input
}

// GetTimeout implements InterchainRequestI
func (ir InterchainRequestAdaptor) GetTimeout() uint64 {
	return ir.Timeout
}

// GetServiceFeeCap implements InterchainRequestI
func (ir InterchainRequestAdaptor) GetServiceFeeCap() string {
	return ir.ServiceFeeCap
}

// ResponseAdaptor is the wrapped response struct of Irita-Hub
type ResponseAdaptor struct {
	StatusCode  int
	Result      string
	Output      string
	ICRequestID string
}

// GetErrMsg implements ResponseI
func (r ResponseAdaptor) GetErrMsg() string {
	switch r.StatusCode {
	case 200:
		return ""

	case 400, 500:
		return r.Result

	default:
		return ""
	}
}

// GetOutput implements ResponseI
func (r ResponseAdaptor) GetOutput() string {
	switch r.StatusCode {
	case 200:
		return r.Output

	case 400, 500:
		return r.Result

	default:
		return ""
	}
}

// GetInterchainRequestID implements ResponseI
func (r ResponseAdaptor) GetInterchainRequestID() string {
	return r.ICRequestID
}
