# Relayer 接口设计

## 基本接口

```go
// ChainI defines the basic chain interface
type ChainI interface {
    GetChainID() string // chain ID getter
}
```

## 应用链接口

```go
// InterchainEventI abstracts the event which is triggered for an interchain service invocation
type InterchainEventI interface {
    GetInvocationID() string // invocation ID getter
    GetServiceName()  string // service name getter
    GetInput()        string // request input getter
    GetTimeout()      uint64 // request timeout getter
}
```

```go
// ResponseI exposes the response related interfaces
type ResponseI interface {
    GetICRequestID()  string // interchain request ID getter
    GetResponse()     string // service response getter
    GetError()        string // error msg getter
}
```

```go
// AppChainI defines the interface to interact with the application chain
type AppChainI interface {
    ChainI
    InterchainEventListener(cb func(icEvent InterchainEventI)) error
    SendResponse(invocationID string, response ResponseI) error
}
```

## IRITA-HUB 链接口

```go
// InterchainRequestI is an interface for the interchain request
type InterchainRequestI interface {
    GetServiceName()   string // service name getter
    GetProvider()      string // provider getter
    GetInput()         string // request input getter
    GetTimeout()       uint64 // request timeout getter
    GetServiceFeeCap() string // service fee cap getter
}
```

```go
// IritaHubChainI defines the interface to interact with IRITA-HUB
type IritaHubChainI interface {
    ChainI
    SendInterchainRequest(request InterchainRequestI) (icRequestID string, err error)
    ResponseListener(icRequestID string) (ResponseI, error)
}
```

## 接口适配器

### 跨链请求适配器

此适配器为 `InterchainRequestI` 的内置实现，用于将跨链事件适配到 IRITA-HUB 的跨链请求。

```go
// InterchainRequestAdaptor bridges the interchain event and Irita-Hub interchain request
type InterchainRequestAdaptor struct {
    ServiceName   string
    Provider      string
    Input         string
    Timeout       uint64
    ServiceFeeCap string
}
```

### 响应适配器

此适配器为 `ResponseI` 的内置实现，用于将 IRITA-HUB 的原生响应适配到应用链所需的形式。

```go
// ResponseAdaptor is the wrapped response struct of Irita-Hub
type ResponseAdaptor struct {
    StatusCode   int
    Result       string
    Output       string
    ICRequestID  string
    ErrMsg       string
}
```

## 链示例实现

### 应用链

```go
type InterchainEvent struct {
    InvocationID string
    ServiceName  string
    Input        string
    Timeout      uint64
}

func (e InterchainEvent) GetInvocationID() string {
    return e.InvocationID
}

func (e InterchainEvent) GetServiceName() string {
    return e.ServiceName
}

func (e InterchainEvent) GetInput() string {
    return e.Input
}

func (e InterchainEvent) GetTimeout() uint64 {
    return e.Timeout
}

type FabricChain struct {
    ChainID        string
    ChannelID      string
    PeerRPCAddrs   []string
    OrdererRPCAddr string
    Client         interface{}
    Key            interface{}
}

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

func (fc FabricChain) GetChainID() string {
    return fc.ChainID
}

func (fc FabricChain) InterchainEventListener(cb func(icEvent InterchainEventI)) error {
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

    return nil
}

func (fc FabricChain) SendResponse(invocationID string, response ResponseI) error {
    return nil
}
```

### IRITA-HUB

```go
type IritaHubChain struct {
    ChainID string
    RPCAddr string
    Client  interface{}
    Key     interface{}
}

func NewIritaHubChain(chainID string, rpcAddr string, client interface{}, key interface{}) IritaHubChain {
    return IritaHubChain{
        ChainID: chainID,
        RPCAddr: rpcAddr,
        Client:  client,
        Key:     key,
        }
}

func (ic IritaHubChain) GetChainID() string {
    return ic.ChainID
}

func (ic IritaHubChain) SendInterchainRequest(request InterchainRequestI) (icRequestID string, err error) {
    return fmt.Sprintf("interchain-request-id-%d", rand.Int()), nil
}

func (ic IritaHubChain) ResponseListener(icRequestID string) (ResponseI, error) {
    response := ResponseAdaptor{
        StatusCode: 200,
        Result: `{"code":"200","message":""}`,
        Output: fmt.Sprintf(`{"rate":"%s"}`, rand.Int()),
        ICRequestID: icRequestID,
        ErrMsg: "",
    }

    return response, nil
}
```
