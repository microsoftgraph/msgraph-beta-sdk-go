package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ThreatIntelligenceHostsItemSslCertificatesHostSslCertificateItemRequestBuilder provides operations to manage the sslCertificates property of the microsoft.graph.security.host entity.
type ThreatIntelligenceHostsItemSslCertificatesHostSslCertificateItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatIntelligenceHostsItemSslCertificatesHostSslCertificateItemRequestBuilderGetQueryParameters get sslCertificates from security
type ThreatIntelligenceHostsItemSslCertificatesHostSslCertificateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ThreatIntelligenceHostsItemSslCertificatesHostSslCertificateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatIntelligenceHostsItemSslCertificatesHostSslCertificateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatIntelligenceHostsItemSslCertificatesHostSslCertificateItemRequestBuilderGetQueryParameters
}
// NewThreatIntelligenceHostsItemSslCertificatesHostSslCertificateItemRequestBuilderInternal instantiates a new HostSslCertificateItemRequestBuilder and sets the default values.
func NewThreatIntelligenceHostsItemSslCertificatesHostSslCertificateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatIntelligenceHostsItemSslCertificatesHostSslCertificateItemRequestBuilder) {
    m := &ThreatIntelligenceHostsItemSslCertificatesHostSslCertificateItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/threatIntelligence/hosts/{host%2Did}/sslCertificates/{hostSslCertificate%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewThreatIntelligenceHostsItemSslCertificatesHostSslCertificateItemRequestBuilder instantiates a new HostSslCertificateItemRequestBuilder and sets the default values.
func NewThreatIntelligenceHostsItemSslCertificatesHostSslCertificateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatIntelligenceHostsItemSslCertificatesHostSslCertificateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatIntelligenceHostsItemSslCertificatesHostSslCertificateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get sslCertificates from security
func (m *ThreatIntelligenceHostsItemSslCertificatesHostSslCertificateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatIntelligenceHostsItemSslCertificatesHostSslCertificateItemRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.HostSslCertificateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateHostSslCertificateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.HostSslCertificateable), nil
}
// ToGetRequestInformation get sslCertificates from security
func (m *ThreatIntelligenceHostsItemSslCertificatesHostSslCertificateItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatIntelligenceHostsItemSslCertificatesHostSslCertificateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
