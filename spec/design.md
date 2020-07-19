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
    GetProvider()     string // provider getter
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
// IServiceMarketI defines the interface for the iService market on the application chain
type IServiceMarketI interface {
    // add a service binding to the iService market
    AddServiceBinding(serviceName, schemas, provider, serviceFee string, qos uint64) error

    // get the service fee by the service name and provider
    GetServiceFee(serviceName string, provider string) (fee string, err error)
}
```

```go
// AppChainI defines the interface to interact with the application chain
type AppChainI interface {
    ChainI

    // listen to the interchain events with an event handler
    InterchainEventListener(cb func(icEvent InterchainEventI)) error

    // send the response to the application chain
    SendResponse(invocationID string, response ResponseI) error

    // iService market interface
    IServiceMarketI
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

    // send the interchain request and handle the response with the given callback
    SendInterchainRequest(request InterchainRequestI, cb func(icRequestID string, response ResponseI)) error
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
    Provider     string
    Input        string
    Timeout      uint64
}

func (e InterchainEvent) GetInvocationID() string {
    return e.InvocationID
}

func (e InterchainEvent) GetServiceName() string {
    return e.ServiceName
}

func (e InterchainEvent) GetProvider() string {
    return e.Provider
}

func (e InterchainEvent) GetInput() string {
    return e.Input
}

func (e InterchainEvent) GetTimeout() uint64 {
    return e.Timeout
}

type FabricChain struct {
    ChannelID      string
    PeerRPCAddrs   []string
    OrdererRPCAddr string
    Client         interface{}
    Key            string
    Passphrase     string
}

func NewFabricChain(
    channelID string,
    peerRPCAddrs []string,
    ordererRPCAddr string,
    client interface{},
    key string,
    Passphrase string,
) FabricChain {
    return FabricChain{
        ChannelID:      channelID,
        PeerRPCAddrs:   peerRPCAddrs,
        OrdererRPCAddr: ordererRPCAddr,
        Client:         client,
        Key:            key,
        Passphrase:     passphrase,
    }
}

func (fc FabricChain) GetChainID() string {
    return fc.ChannelID
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
    ChainID    string
    RPCAddr    string
    Client     interface{}
    Key        string
    Passphrase string
}

func NewIritaHubChain(chainID string, rpcAddr string, client interface{}, key string, passphrase string) IritaHubChain {
    return IritaHubChain{
        ChainID:    chainID,
        RPCAddr:    rpcAddr,
        Client:     client,
        Key:        key,
        Passphrase: passphrase,
        }
}

func (ic IritaHubChain) GetChainID() string {
    return ic.ChainID
}

func (ic IritaHubChain) SendInterchainRequest(request InterchainRequestI, cb func(icRequestID string, response ResponseI)) error {
    return nil
}
```
