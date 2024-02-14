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
// ThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilderGetQueryParameters the most recent sslCertificate used to communicate on the port.
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
// NewThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilderInternal instantiates a new ThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilder and sets the default values.
func NewThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilder) {
    m := &ThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/threatIntelligence/hostPorts/{hostPort%2Did}/mostRecentSslCertificate{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilder instantiates a new ThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilder and sets the default values.
func NewThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the most recent sslCertificate used to communicate on the port.
// returns a SslCertificateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SslCertificateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// ToGetRequestInformation the most recent sslCertificate used to communicate on the port.
// returns a *RequestInformation when successful
func (m *ThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilder when successful
func (m *ThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilder) WithUrl(rawUrl string)(*ThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilder) {
    return NewThreatIntelligenceHostPortsItemMostRecentSslCertificateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
