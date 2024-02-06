package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ThreatIntelligenceHostsItemParentHostPairsHostPairItemRequestBuilder provides operations to manage the parentHostPairs property of the microsoft.graph.security.host entity.
type ThreatIntelligenceHostsItemParentHostPairsHostPairItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatIntelligenceHostsItemParentHostPairsHostPairItemRequestBuilderGetQueryParameters the hostPairs that are associated with a host, where that host is the childHost and has an incoming pairing with a parentHost.
type ThreatIntelligenceHostsItemParentHostPairsHostPairItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ThreatIntelligenceHostsItemParentHostPairsHostPairItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatIntelligenceHostsItemParentHostPairsHostPairItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatIntelligenceHostsItemParentHostPairsHostPairItemRequestBuilderGetQueryParameters
}
// NewThreatIntelligenceHostsItemParentHostPairsHostPairItemRequestBuilderInternal instantiates a new HostPairItemRequestBuilder and sets the default values.
func NewThreatIntelligenceHostsItemParentHostPairsHostPairItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatIntelligenceHostsItemParentHostPairsHostPairItemRequestBuilder) {
    m := &ThreatIntelligenceHostsItemParentHostPairsHostPairItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/threatIntelligence/hosts/{host%2Did}/parentHostPairs/{hostPair%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewThreatIntelligenceHostsItemParentHostPairsHostPairItemRequestBuilder instantiates a new HostPairItemRequestBuilder and sets the default values.
func NewThreatIntelligenceHostsItemParentHostPairsHostPairItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatIntelligenceHostsItemParentHostPairsHostPairItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatIntelligenceHostsItemParentHostPairsHostPairItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the hostPairs that are associated with a host, where that host is the childHost and has an incoming pairing with a parentHost.
func (m *ThreatIntelligenceHostsItemParentHostPairsHostPairItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatIntelligenceHostsItemParentHostPairsHostPairItemRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.HostPairable, error) {
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
// ToGetRequestInformation the hostPairs that are associated with a host, where that host is the childHost and has an incoming pairing with a parentHost.
func (m *ThreatIntelligenceHostsItemParentHostPairsHostPairItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatIntelligenceHostsItemParentHostPairsHostPairItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ThreatIntelligenceHostsItemParentHostPairsHostPairItemRequestBuilder) WithUrl(rawUrl string)(*ThreatIntelligenceHostsItemParentHostPairsHostPairItemRequestBuilder) {
    return NewThreatIntelligenceHostsItemParentHostPairsHostPairItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
