// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ThreatIntelligenceHostsHostItemRequestBuilder provides operations to manage the hosts property of the microsoft.graph.security.threatIntelligence entity.
type ThreatIntelligenceHostsHostItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatIntelligenceHostsHostItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatIntelligenceHostsHostItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ThreatIntelligenceHostsHostItemRequestBuilderGetQueryParameters read the properties and relationships of a host object. The host resource is the abstract base type that returns an implementation. A host can be of one of the following types:
type ThreatIntelligenceHostsHostItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ThreatIntelligenceHostsHostItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatIntelligenceHostsHostItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatIntelligenceHostsHostItemRequestBuilderGetQueryParameters
}
// ThreatIntelligenceHostsHostItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatIntelligenceHostsHostItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ChildHostPairs provides operations to manage the childHostPairs property of the microsoft.graph.security.host entity.
// returns a *ThreatIntelligenceHostsItemChildHostPairsRequestBuilder when successful
func (m *ThreatIntelligenceHostsHostItemRequestBuilder) ChildHostPairs()(*ThreatIntelligenceHostsItemChildHostPairsRequestBuilder) {
    return NewThreatIntelligenceHostsItemChildHostPairsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Components provides operations to manage the components property of the microsoft.graph.security.host entity.
// returns a *ThreatIntelligenceHostsItemComponentsRequestBuilder when successful
func (m *ThreatIntelligenceHostsHostItemRequestBuilder) Components()(*ThreatIntelligenceHostsItemComponentsRequestBuilder) {
    return NewThreatIntelligenceHostsItemComponentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewThreatIntelligenceHostsHostItemRequestBuilderInternal instantiates a new ThreatIntelligenceHostsHostItemRequestBuilder and sets the default values.
func NewThreatIntelligenceHostsHostItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatIntelligenceHostsHostItemRequestBuilder) {
    m := &ThreatIntelligenceHostsHostItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/threatIntelligence/hosts/{host%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewThreatIntelligenceHostsHostItemRequestBuilder instantiates a new ThreatIntelligenceHostsHostItemRequestBuilder and sets the default values.
func NewThreatIntelligenceHostsHostItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatIntelligenceHostsHostItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatIntelligenceHostsHostItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Cookies provides operations to manage the cookies property of the microsoft.graph.security.host entity.
// returns a *ThreatIntelligenceHostsItemCookiesRequestBuilder when successful
func (m *ThreatIntelligenceHostsHostItemRequestBuilder) Cookies()(*ThreatIntelligenceHostsItemCookiesRequestBuilder) {
    return NewThreatIntelligenceHostsItemCookiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property hosts for security
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatIntelligenceHostsHostItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ThreatIntelligenceHostsHostItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get read the properties and relationships of a host object. The host resource is the abstract base type that returns an implementation. A host can be of one of the following types:
// returns a Hostable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-host-get?view=graph-rest-beta
func (m *ThreatIntelligenceHostsHostItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatIntelligenceHostsHostItemRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.Hostable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateHostFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.Hostable), nil
}
// HostPairs provides operations to manage the hostPairs property of the microsoft.graph.security.host entity.
// returns a *ThreatIntelligenceHostsItemHostPairsRequestBuilder when successful
func (m *ThreatIntelligenceHostsHostItemRequestBuilder) HostPairs()(*ThreatIntelligenceHostsItemHostPairsRequestBuilder) {
    return NewThreatIntelligenceHostsItemHostPairsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ParentHostPairs provides operations to manage the parentHostPairs property of the microsoft.graph.security.host entity.
// returns a *ThreatIntelligenceHostsItemParentHostPairsRequestBuilder when successful
func (m *ThreatIntelligenceHostsHostItemRequestBuilder) ParentHostPairs()(*ThreatIntelligenceHostsItemParentHostPairsRequestBuilder) {
    return NewThreatIntelligenceHostsItemParentHostPairsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PassiveDns provides operations to manage the passiveDns property of the microsoft.graph.security.host entity.
// returns a *ThreatIntelligenceHostsItemPassiveDnsRequestBuilder when successful
func (m *ThreatIntelligenceHostsHostItemRequestBuilder) PassiveDns()(*ThreatIntelligenceHostsItemPassiveDnsRequestBuilder) {
    return NewThreatIntelligenceHostsItemPassiveDnsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PassiveDnsReverse provides operations to manage the passiveDnsReverse property of the microsoft.graph.security.host entity.
// returns a *ThreatIntelligenceHostsItemPassiveDnsReverseRequestBuilder when successful
func (m *ThreatIntelligenceHostsHostItemRequestBuilder) PassiveDnsReverse()(*ThreatIntelligenceHostsItemPassiveDnsReverseRequestBuilder) {
    return NewThreatIntelligenceHostsItemPassiveDnsReverseRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property hosts in security
// returns a Hostable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatIntelligenceHostsHostItemRequestBuilder) Patch(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.Hostable, requestConfiguration *ThreatIntelligenceHostsHostItemRequestBuilderPatchRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.Hostable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateHostFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.Hostable), nil
}
// Ports provides operations to manage the ports property of the microsoft.graph.security.host entity.
// returns a *ThreatIntelligenceHostsItemPortsRequestBuilder when successful
func (m *ThreatIntelligenceHostsHostItemRequestBuilder) Ports()(*ThreatIntelligenceHostsItemPortsRequestBuilder) {
    return NewThreatIntelligenceHostsItemPortsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Reputation provides operations to manage the reputation property of the microsoft.graph.security.host entity.
// returns a *ThreatIntelligenceHostsItemReputationRequestBuilder when successful
func (m *ThreatIntelligenceHostsHostItemRequestBuilder) Reputation()(*ThreatIntelligenceHostsItemReputationRequestBuilder) {
    return NewThreatIntelligenceHostsItemReputationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SslCertificates provides operations to manage the sslCertificates property of the microsoft.graph.security.host entity.
// returns a *ThreatIntelligenceHostsItemSslCertificatesRequestBuilder when successful
func (m *ThreatIntelligenceHostsHostItemRequestBuilder) SslCertificates()(*ThreatIntelligenceHostsItemSslCertificatesRequestBuilder) {
    return NewThreatIntelligenceHostsItemSslCertificatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Subdomains provides operations to manage the subdomains property of the microsoft.graph.security.host entity.
// returns a *ThreatIntelligenceHostsItemSubdomainsRequestBuilder when successful
func (m *ThreatIntelligenceHostsHostItemRequestBuilder) Subdomains()(*ThreatIntelligenceHostsItemSubdomainsRequestBuilder) {
    return NewThreatIntelligenceHostsItemSubdomainsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property hosts for security
// returns a *RequestInformation when successful
func (m *ThreatIntelligenceHostsHostItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ThreatIntelligenceHostsHostItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a host object. The host resource is the abstract base type that returns an implementation. A host can be of one of the following types:
// returns a *RequestInformation when successful
func (m *ThreatIntelligenceHostsHostItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatIntelligenceHostsHostItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property hosts in security
// returns a *RequestInformation when successful
func (m *ThreatIntelligenceHostsHostItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.Hostable, requestConfiguration *ThreatIntelligenceHostsHostItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// Trackers provides operations to manage the trackers property of the microsoft.graph.security.host entity.
// returns a *ThreatIntelligenceHostsItemTrackersRequestBuilder when successful
func (m *ThreatIntelligenceHostsHostItemRequestBuilder) Trackers()(*ThreatIntelligenceHostsItemTrackersRequestBuilder) {
    return NewThreatIntelligenceHostsItemTrackersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Whois provides operations to manage the whois property of the microsoft.graph.security.host entity.
// returns a *ThreatIntelligenceHostsItemWhoisRequestBuilder when successful
func (m *ThreatIntelligenceHostsHostItemRequestBuilder) Whois()(*ThreatIntelligenceHostsItemWhoisRequestBuilder) {
    return NewThreatIntelligenceHostsItemWhoisRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ThreatIntelligenceHostsHostItemRequestBuilder when successful
func (m *ThreatIntelligenceHostsHostItemRequestBuilder) WithUrl(rawUrl string)(*ThreatIntelligenceHostsHostItemRequestBuilder) {
    return NewThreatIntelligenceHostsHostItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
