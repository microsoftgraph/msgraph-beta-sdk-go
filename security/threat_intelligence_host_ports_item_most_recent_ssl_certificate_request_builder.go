package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilder provides operations to manage the mostRecentSslCertificate property of the microsoft.graph.security.hostPort entity.
type ThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilderGetQueryParameters get mostRecentSslCertificate from security
type ThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilderGetQueryParameters
}
// NewThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilderInternal instantiates a new MostRecentSslCertificateRequestBuilder and sets the default values.
func NewThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilder) {
    m := &ThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/threatIntelligence/hostPorts/{hostPort%2Did}/mostRecentSslCertificate{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilder instantiates a new MostRecentSslCertificateRequestBuilder and sets the default values.
func NewThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get mostRecentSslCertificate from security
func (m *ThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SslCertificateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateSslCertificateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SslCertificateable), nil
}
// ToGetRequestInformation get mostRecentSslCertificate from security
func (m *ThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilder) WithUrl(rawUrl string)(*ThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilder) {
    return NewThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
