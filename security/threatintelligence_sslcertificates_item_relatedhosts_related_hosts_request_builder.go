package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ThreatintelligenceSslcertificatesItemRelatedhostsRelatedHostsRequestBuilder provides operations to manage the relatedHosts property of the microsoft.graph.security.sslCertificate entity.
type ThreatintelligenceSslcertificatesItemRelatedhostsRelatedHostsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatintelligenceSslcertificatesItemRelatedhostsRelatedHostsRequestBuilderGetQueryParameters the hosts related with this sslCertificate.
type ThreatintelligenceSslcertificatesItemRelatedhostsRelatedHostsRequestBuilderGetQueryParameters struct {
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
// ThreatintelligenceSslcertificatesItemRelatedhostsRelatedHostsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceSslcertificatesItemRelatedhostsRelatedHostsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatintelligenceSslcertificatesItemRelatedhostsRelatedHostsRequestBuilderGetQueryParameters
}
// ByHostId provides operations to manage the relatedHosts property of the microsoft.graph.security.sslCertificate entity.
// returns a *ThreatintelligenceSslcertificatesItemRelatedhostsHostItemRequestBuilder when successful
func (m *ThreatintelligenceSslcertificatesItemRelatedhostsRelatedHostsRequestBuilder) ByHostId(hostId string)(*ThreatintelligenceSslcertificatesItemRelatedhostsHostItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if hostId != "" {
        urlTplParams["host%2Did"] = hostId
    }
    return NewThreatintelligenceSslcertificatesItemRelatedhostsHostItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewThreatintelligenceSslcertificatesItemRelatedhostsRelatedHostsRequestBuilderInternal instantiates a new ThreatintelligenceSslcertificatesItemRelatedhostsRelatedHostsRequestBuilder and sets the default values.
func NewThreatintelligenceSslcertificatesItemRelatedhostsRelatedHostsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceSslcertificatesItemRelatedhostsRelatedHostsRequestBuilder) {
    m := &ThreatintelligenceSslcertificatesItemRelatedhostsRelatedHostsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/threatIntelligence/sslCertificates/{sslCertificate%2Did}/relatedHosts{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewThreatintelligenceSslcertificatesItemRelatedhostsRelatedHostsRequestBuilder instantiates a new ThreatintelligenceSslcertificatesItemRelatedhostsRelatedHostsRequestBuilder and sets the default values.
func NewThreatintelligenceSslcertificatesItemRelatedhostsRelatedHostsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceSslcertificatesItemRelatedhostsRelatedHostsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatintelligenceSslcertificatesItemRelatedhostsRelatedHostsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ThreatintelligenceSslcertificatesItemRelatedhostsCountRequestBuilder when successful
func (m *ThreatintelligenceSslcertificatesItemRelatedhostsRelatedHostsRequestBuilder) Count()(*ThreatintelligenceSslcertificatesItemRelatedhostsCountRequestBuilder) {
    return NewThreatintelligenceSslcertificatesItemRelatedhostsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the hosts related with this sslCertificate.
// returns a HostCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatintelligenceSslcertificatesItemRelatedhostsRelatedHostsRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatintelligenceSslcertificatesItemRelatedhostsRelatedHostsRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.HostCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateHostCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.HostCollectionResponseable), nil
}
// ToGetRequestInformation the hosts related with this sslCertificate.
// returns a *RequestInformation when successful
func (m *ThreatintelligenceSslcertificatesItemRelatedhostsRelatedHostsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatintelligenceSslcertificatesItemRelatedhostsRelatedHostsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ThreatintelligenceSslcertificatesItemRelatedhostsRelatedHostsRequestBuilder when successful
func (m *ThreatintelligenceSslcertificatesItemRelatedhostsRelatedHostsRequestBuilder) WithUrl(rawUrl string)(*ThreatintelligenceSslcertificatesItemRelatedhostsRelatedHostsRequestBuilder) {
    return NewThreatintelligenceSslcertificatesItemRelatedhostsRelatedHostsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
