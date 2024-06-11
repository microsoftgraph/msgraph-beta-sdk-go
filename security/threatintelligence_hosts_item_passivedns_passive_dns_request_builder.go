package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ThreatintelligenceHostsItemPassivednsPassiveDnsRequestBuilder provides operations to manage the passiveDns property of the microsoft.graph.security.host entity.
type ThreatintelligenceHostsItemPassivednsPassiveDnsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatintelligenceHostsItemPassivednsPassiveDnsRequestBuilderGetQueryParameters get a list of passiveDnsRecord resources associated with a host. This method is a forward DNS lookup that queries the IP address of the specified host using its hostname. 
type ThreatintelligenceHostsItemPassivednsPassiveDnsRequestBuilderGetQueryParameters struct {
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
// ThreatintelligenceHostsItemPassivednsPassiveDnsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceHostsItemPassivednsPassiveDnsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatintelligenceHostsItemPassivednsPassiveDnsRequestBuilderGetQueryParameters
}
// ByPassiveDnsRecordId provides operations to manage the passiveDns property of the microsoft.graph.security.host entity.
// returns a *ThreatintelligenceHostsItemPassivednsPassiveDnsRecordItemRequestBuilder when successful
func (m *ThreatintelligenceHostsItemPassivednsPassiveDnsRequestBuilder) ByPassiveDnsRecordId(passiveDnsRecordId string)(*ThreatintelligenceHostsItemPassivednsPassiveDnsRecordItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if passiveDnsRecordId != "" {
        urlTplParams["passiveDnsRecord%2Did"] = passiveDnsRecordId
    }
    return NewThreatintelligenceHostsItemPassivednsPassiveDnsRecordItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewThreatintelligenceHostsItemPassivednsPassiveDnsRequestBuilderInternal instantiates a new ThreatintelligenceHostsItemPassivednsPassiveDnsRequestBuilder and sets the default values.
func NewThreatintelligenceHostsItemPassivednsPassiveDnsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceHostsItemPassivednsPassiveDnsRequestBuilder) {
    m := &ThreatintelligenceHostsItemPassivednsPassiveDnsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/threatIntelligence/hosts/{host%2Did}/passiveDns{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewThreatintelligenceHostsItemPassivednsPassiveDnsRequestBuilder instantiates a new ThreatintelligenceHostsItemPassivednsPassiveDnsRequestBuilder and sets the default values.
func NewThreatintelligenceHostsItemPassivednsPassiveDnsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceHostsItemPassivednsPassiveDnsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatintelligenceHostsItemPassivednsPassiveDnsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ThreatintelligenceHostsItemPassivednsCountRequestBuilder when successful
func (m *ThreatintelligenceHostsItemPassivednsPassiveDnsRequestBuilder) Count()(*ThreatintelligenceHostsItemPassivednsCountRequestBuilder) {
    return NewThreatintelligenceHostsItemPassivednsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of passiveDnsRecord resources associated with a host. This method is a forward DNS lookup that queries the IP address of the specified host using its hostname. 
// returns a PassiveDnsRecordCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-host-list-passivedns?view=graph-rest-beta
func (m *ThreatintelligenceHostsItemPassivednsPassiveDnsRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatintelligenceHostsItemPassivednsPassiveDnsRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.PassiveDnsRecordCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// ToGetRequestInformation get a list of passiveDnsRecord resources associated with a host. This method is a forward DNS lookup that queries the IP address of the specified host using its hostname. 
// returns a *RequestInformation when successful
func (m *ThreatintelligenceHostsItemPassivednsPassiveDnsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatintelligenceHostsItemPassivednsPassiveDnsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ThreatintelligenceHostsItemPassivednsPassiveDnsRequestBuilder when successful
func (m *ThreatintelligenceHostsItemPassivednsPassiveDnsRequestBuilder) WithUrl(rawUrl string)(*ThreatintelligenceHostsItemPassivednsPassiveDnsRequestBuilder) {
    return NewThreatintelligenceHostsItemPassivednsPassiveDnsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
