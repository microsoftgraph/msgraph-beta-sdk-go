package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ThreatintelligenceThreatIntelligenceRequestBuilder provides operations to manage the threatIntelligence property of the microsoft.graph.security entity.
type ThreatintelligenceThreatIntelligenceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatintelligenceThreatIntelligenceRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceThreatIntelligenceRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ThreatintelligenceThreatIntelligenceRequestBuilderGetQueryParameters get threatIntelligence from security
type ThreatintelligenceThreatIntelligenceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ThreatintelligenceThreatIntelligenceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceThreatIntelligenceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatintelligenceThreatIntelligenceRequestBuilderGetQueryParameters
}
// ThreatintelligenceThreatIntelligenceRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceThreatIntelligenceRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ArticleIndicators provides operations to manage the articleIndicators property of the microsoft.graph.security.threatIntelligence entity.
// returns a *ThreatintelligenceArticleindicatorsArticleIndicatorsRequestBuilder when successful
func (m *ThreatintelligenceThreatIntelligenceRequestBuilder) ArticleIndicators()(*ThreatintelligenceArticleindicatorsArticleIndicatorsRequestBuilder) {
    return NewThreatintelligenceArticleindicatorsArticleIndicatorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Articles provides operations to manage the articles property of the microsoft.graph.security.threatIntelligence entity.
// returns a *ThreatintelligenceArticlesRequestBuilder when successful
func (m *ThreatintelligenceThreatIntelligenceRequestBuilder) Articles()(*ThreatintelligenceArticlesRequestBuilder) {
    return NewThreatintelligenceArticlesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewThreatintelligenceThreatIntelligenceRequestBuilderInternal instantiates a new ThreatintelligenceThreatIntelligenceRequestBuilder and sets the default values.
func NewThreatintelligenceThreatIntelligenceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceThreatIntelligenceRequestBuilder) {
    m := &ThreatintelligenceThreatIntelligenceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/threatIntelligence{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewThreatintelligenceThreatIntelligenceRequestBuilder instantiates a new ThreatintelligenceThreatIntelligenceRequestBuilder and sets the default values.
func NewThreatintelligenceThreatIntelligenceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceThreatIntelligenceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatintelligenceThreatIntelligenceRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property threatIntelligence for security
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatintelligenceThreatIntelligenceRequestBuilder) Delete(ctx context.Context, requestConfiguration *ThreatintelligenceThreatIntelligenceRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get threatIntelligence from security
// returns a ThreatIntelligenceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatintelligenceThreatIntelligenceRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatintelligenceThreatIntelligenceRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ThreatIntelligenceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateThreatIntelligenceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ThreatIntelligenceable), nil
}
// HostComponents provides operations to manage the hostComponents property of the microsoft.graph.security.threatIntelligence entity.
// returns a *ThreatintelligenceHostcomponentsHostComponentsRequestBuilder when successful
func (m *ThreatintelligenceThreatIntelligenceRequestBuilder) HostComponents()(*ThreatintelligenceHostcomponentsHostComponentsRequestBuilder) {
    return NewThreatintelligenceHostcomponentsHostComponentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// HostCookies provides operations to manage the hostCookies property of the microsoft.graph.security.threatIntelligence entity.
// returns a *ThreatintelligenceHostcookiesHostCookiesRequestBuilder when successful
func (m *ThreatintelligenceThreatIntelligenceRequestBuilder) HostCookies()(*ThreatintelligenceHostcookiesHostCookiesRequestBuilder) {
    return NewThreatintelligenceHostcookiesHostCookiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// HostPairs provides operations to manage the hostPairs property of the microsoft.graph.security.threatIntelligence entity.
// returns a *ThreatintelligenceHostpairsHostPairsRequestBuilder when successful
func (m *ThreatintelligenceThreatIntelligenceRequestBuilder) HostPairs()(*ThreatintelligenceHostpairsHostPairsRequestBuilder) {
    return NewThreatintelligenceHostpairsHostPairsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// HostPorts provides operations to manage the hostPorts property of the microsoft.graph.security.threatIntelligence entity.
// returns a *ThreatintelligenceHostportsHostPortsRequestBuilder when successful
func (m *ThreatintelligenceThreatIntelligenceRequestBuilder) HostPorts()(*ThreatintelligenceHostportsHostPortsRequestBuilder) {
    return NewThreatintelligenceHostportsHostPortsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Hosts provides operations to manage the hosts property of the microsoft.graph.security.threatIntelligence entity.
// returns a *ThreatintelligenceHostsRequestBuilder when successful
func (m *ThreatintelligenceThreatIntelligenceRequestBuilder) Hosts()(*ThreatintelligenceHostsRequestBuilder) {
    return NewThreatintelligenceHostsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// HostSslCertificates provides operations to manage the hostSslCertificates property of the microsoft.graph.security.threatIntelligence entity.
// returns a *ThreatintelligenceHostsslcertificatesHostSslCertificatesRequestBuilder when successful
func (m *ThreatintelligenceThreatIntelligenceRequestBuilder) HostSslCertificates()(*ThreatintelligenceHostsslcertificatesHostSslCertificatesRequestBuilder) {
    return NewThreatintelligenceHostsslcertificatesHostSslCertificatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// HostTrackers provides operations to manage the hostTrackers property of the microsoft.graph.security.threatIntelligence entity.
// returns a *ThreatintelligenceHosttrackersHostTrackersRequestBuilder when successful
func (m *ThreatintelligenceThreatIntelligenceRequestBuilder) HostTrackers()(*ThreatintelligenceHosttrackersHostTrackersRequestBuilder) {
    return NewThreatintelligenceHosttrackersHostTrackersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IntelligenceProfileIndicators provides operations to manage the intelligenceProfileIndicators property of the microsoft.graph.security.threatIntelligence entity.
// returns a *ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorsRequestBuilder when successful
func (m *ThreatintelligenceThreatIntelligenceRequestBuilder) IntelligenceProfileIndicators()(*ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorsRequestBuilder) {
    return NewThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IntelProfiles provides operations to manage the intelProfiles property of the microsoft.graph.security.threatIntelligence entity.
// returns a *ThreatintelligenceIntelprofilesIntelProfilesRequestBuilder when successful
func (m *ThreatintelligenceThreatIntelligenceRequestBuilder) IntelProfiles()(*ThreatintelligenceIntelprofilesIntelProfilesRequestBuilder) {
    return NewThreatintelligenceIntelprofilesIntelProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PassiveDnsRecords provides operations to manage the passiveDnsRecords property of the microsoft.graph.security.threatIntelligence entity.
// returns a *ThreatintelligencePassivednsrecordsPassiveDnsRecordsRequestBuilder when successful
func (m *ThreatintelligenceThreatIntelligenceRequestBuilder) PassiveDnsRecords()(*ThreatintelligencePassivednsrecordsPassiveDnsRecordsRequestBuilder) {
    return NewThreatintelligencePassivednsrecordsPassiveDnsRecordsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property threatIntelligence in security
// returns a ThreatIntelligenceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatintelligenceThreatIntelligenceRequestBuilder) Patch(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ThreatIntelligenceable, requestConfiguration *ThreatintelligenceThreatIntelligenceRequestBuilderPatchRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ThreatIntelligenceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateThreatIntelligenceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ThreatIntelligenceable), nil
}
// SslCertificates provides operations to manage the sslCertificates property of the microsoft.graph.security.threatIntelligence entity.
// returns a *ThreatintelligenceSslcertificatesSslCertificatesRequestBuilder when successful
func (m *ThreatintelligenceThreatIntelligenceRequestBuilder) SslCertificates()(*ThreatintelligenceSslcertificatesSslCertificatesRequestBuilder) {
    return NewThreatintelligenceSslcertificatesSslCertificatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Subdomains provides operations to manage the subdomains property of the microsoft.graph.security.threatIntelligence entity.
// returns a *ThreatintelligenceSubdomainsRequestBuilder when successful
func (m *ThreatintelligenceThreatIntelligenceRequestBuilder) Subdomains()(*ThreatintelligenceSubdomainsRequestBuilder) {
    return NewThreatintelligenceSubdomainsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property threatIntelligence for security
// returns a *RequestInformation when successful
func (m *ThreatintelligenceThreatIntelligenceRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ThreatintelligenceThreatIntelligenceRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get threatIntelligence from security
// returns a *RequestInformation when successful
func (m *ThreatintelligenceThreatIntelligenceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatintelligenceThreatIntelligenceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property threatIntelligence in security
// returns a *RequestInformation when successful
func (m *ThreatintelligenceThreatIntelligenceRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ThreatIntelligenceable, requestConfiguration *ThreatintelligenceThreatIntelligenceRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Vulnerabilities provides operations to manage the vulnerabilities property of the microsoft.graph.security.threatIntelligence entity.
// returns a *ThreatintelligenceVulnerabilitiesRequestBuilder when successful
func (m *ThreatintelligenceThreatIntelligenceRequestBuilder) Vulnerabilities()(*ThreatintelligenceVulnerabilitiesRequestBuilder) {
    return NewThreatintelligenceVulnerabilitiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WhoisHistoryRecords provides operations to manage the whoisHistoryRecords property of the microsoft.graph.security.threatIntelligence entity.
// returns a *ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordsRequestBuilder when successful
func (m *ThreatintelligenceThreatIntelligenceRequestBuilder) WhoisHistoryRecords()(*ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordsRequestBuilder) {
    return NewThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WhoisRecords provides operations to manage the whoisRecords property of the microsoft.graph.security.threatIntelligence entity.
// returns a *ThreatintelligenceWhoisrecordsWhoisRecordsRequestBuilder when successful
func (m *ThreatintelligenceThreatIntelligenceRequestBuilder) WhoisRecords()(*ThreatintelligenceWhoisrecordsWhoisRecordsRequestBuilder) {
    return NewThreatintelligenceWhoisrecordsWhoisRecordsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ThreatintelligenceThreatIntelligenceRequestBuilder when successful
func (m *ThreatintelligenceThreatIntelligenceRequestBuilder) WithUrl(rawUrl string)(*ThreatintelligenceThreatIntelligenceRequestBuilder) {
    return NewThreatintelligenceThreatIntelligenceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
