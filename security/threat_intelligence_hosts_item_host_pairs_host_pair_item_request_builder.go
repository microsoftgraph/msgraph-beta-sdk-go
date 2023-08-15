package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ThreatIntelligenceHostsItemHostPairsHostPairItemRequestBuilder provides operations to manage the hostPairs property of the microsoft.graph.security.host entity.
type ThreatIntelligenceHostsItemHostPairsHostPairItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatIntelligenceHostsItemHostPairsHostPairItemRequestBuilderGetQueryParameters the hostPairs that are associated with this host, where this host is either the parentHost or childHost.
type ThreatIntelligenceHostsItemHostPairsHostPairItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ThreatIntelligenceHostsItemHostPairsHostPairItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatIntelligenceHostsItemHostPairsHostPairItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatIntelligenceHostsItemHostPairsHostPairItemRequestBuilderGetQueryParameters
}
// NewThreatIntelligenceHostsItemHostPairsHostPairItemRequestBuilderInternal instantiates a new HostPairItemRequestBuilder and sets the default values.
func NewThreatIntelligenceHostsItemHostPairsHostPairItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatIntelligenceHostsItemHostPairsHostPairItemRequestBuilder) {
    m := &ThreatIntelligenceHostsItemHostPairsHostPairItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/threatIntelligence/hosts/{host%2Did}/hostPairs/{hostPair%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewThreatIntelligenceHostsItemHostPairsHostPairItemRequestBuilder instantiates a new HostPairItemRequestBuilder and sets the default values.
func NewThreatIntelligenceHostsItemHostPairsHostPairItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatIntelligenceHostsItemHostPairsHostPairItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatIntelligenceHostsItemHostPairsHostPairItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the hostPairs that are associated with this host, where this host is either the parentHost or childHost.
func (m *ThreatIntelligenceHostsItemHostPairsHostPairItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatIntelligenceHostsItemHostPairsHostPairItemRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.HostPairable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateHostPairFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.HostPairable), nil
}
// ToGetRequestInformation the hostPairs that are associated with this host, where this host is either the parentHost or childHost.
func (m *ThreatIntelligenceHostsItemHostPairsHostPairItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatIntelligenceHostsItemHostPairsHostPairItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
