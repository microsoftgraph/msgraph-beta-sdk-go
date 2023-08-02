package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ThreatIntelligenceHostsItemPassiveDnsReverseRequestBuilder provides operations to manage the passiveDnsReverse property of the microsoft.graph.security.host entity.
type ThreatIntelligenceHostsItemPassiveDnsReverseRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatIntelligenceHostsItemPassiveDnsReverseRequestBuilderGetQueryParameters get a list of passiveDnsRecord resources.
type ThreatIntelligenceHostsItemPassiveDnsReverseRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ThreatIntelligenceHostsItemPassiveDnsReverseRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatIntelligenceHostsItemPassiveDnsReverseRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatIntelligenceHostsItemPassiveDnsReverseRequestBuilderGetQueryParameters
}
// ByPassiveDnsRecordId provides operations to manage the passiveDnsReverse property of the microsoft.graph.security.host entity.
func (m *ThreatIntelligenceHostsItemPassiveDnsReverseRequestBuilder) ByPassiveDnsRecordId(passiveDnsRecordId string)(*ThreatIntelligenceHostsItemPassiveDnsReversePassiveDnsRecordItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if passiveDnsRecordId != "" {
        urlTplParams["passiveDnsRecord%2Did"] = passiveDnsRecordId
    }
    return NewThreatIntelligenceHostsItemPassiveDnsReversePassiveDnsRecordItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewThreatIntelligenceHostsItemPassiveDnsReverseRequestBuilderInternal instantiates a new PassiveDnsReverseRequestBuilder and sets the default values.
func NewThreatIntelligenceHostsItemPassiveDnsReverseRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatIntelligenceHostsItemPassiveDnsReverseRequestBuilder) {
    m := &ThreatIntelligenceHostsItemPassiveDnsReverseRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/threatIntelligence/hosts/{host%2Did}/passiveDnsReverse{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewThreatIntelligenceHostsItemPassiveDnsReverseRequestBuilder instantiates a new PassiveDnsReverseRequestBuilder and sets the default values.
func NewThreatIntelligenceHostsItemPassiveDnsReverseRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatIntelligenceHostsItemPassiveDnsReverseRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatIntelligenceHostsItemPassiveDnsReverseRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ThreatIntelligenceHostsItemPassiveDnsReverseRequestBuilder) Count()(*ThreatIntelligenceHostsItemPassiveDnsReverseCountRequestBuilder) {
    return NewThreatIntelligenceHostsItemPassiveDnsReverseCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of passiveDnsRecord resources.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-host-list-passivednsreverse?view=graph-rest-1.0
func (m *ThreatIntelligenceHostsItemPassiveDnsReverseRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatIntelligenceHostsItemPassiveDnsReverseRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.PassiveDnsRecordCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreatePassiveDnsRecordCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.PassiveDnsRecordCollectionResponseable), nil
}
// ToGetRequestInformation get a list of passiveDnsRecord resources.
func (m *ThreatIntelligenceHostsItemPassiveDnsReverseRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatIntelligenceHostsItemPassiveDnsReverseRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
