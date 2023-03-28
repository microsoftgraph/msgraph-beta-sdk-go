package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ThreatIntelligenceRequestBuilder provides operations to manage the threatIntelligence property of the microsoft.graph.security entity.
type ThreatIntelligenceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatIntelligenceRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatIntelligenceRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ThreatIntelligenceRequestBuilderGetQueryParameters get threatIntelligence from security
type ThreatIntelligenceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ThreatIntelligenceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatIntelligenceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatIntelligenceRequestBuilderGetQueryParameters
}
// ThreatIntelligenceRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatIntelligenceRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ArticleIndicators provides operations to manage the articleIndicators property of the microsoft.graph.security.threatIntelligence entity.
func (m *ThreatIntelligenceRequestBuilder) ArticleIndicators()(*ThreatIntelligenceArticleIndicatorsRequestBuilder) {
    return NewThreatIntelligenceArticleIndicatorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ArticleIndicatorsById provides operations to manage the articleIndicators property of the microsoft.graph.security.threatIntelligence entity.
func (m *ThreatIntelligenceRequestBuilder) ArticleIndicatorsById(id string)(*ThreatIntelligenceArticleIndicatorsArticleIndicatorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["articleIndicator%2Did"] = id
    }
    return NewThreatIntelligenceArticleIndicatorsArticleIndicatorItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Articles provides operations to manage the articles property of the microsoft.graph.security.threatIntelligence entity.
func (m *ThreatIntelligenceRequestBuilder) Articles()(*ThreatIntelligenceArticlesRequestBuilder) {
    return NewThreatIntelligenceArticlesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ArticlesById provides operations to manage the articles property of the microsoft.graph.security.threatIntelligence entity.
func (m *ThreatIntelligenceRequestBuilder) ArticlesById(id string)(*ThreatIntelligenceArticlesArticleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["article%2Did"] = id
    }
    return NewThreatIntelligenceArticlesArticleItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewThreatIntelligenceRequestBuilderInternal instantiates a new ThreatIntelligenceRequestBuilder and sets the default values.
func NewThreatIntelligenceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatIntelligenceRequestBuilder) {
    m := &ThreatIntelligenceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/threatIntelligence{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewThreatIntelligenceRequestBuilder instantiates a new ThreatIntelligenceRequestBuilder and sets the default values.
func NewThreatIntelligenceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatIntelligenceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatIntelligenceRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property threatIntelligence for security
func (m *ThreatIntelligenceRequestBuilder) Delete(ctx context.Context, requestConfiguration *ThreatIntelligenceRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get threatIntelligence from security
func (m *ThreatIntelligenceRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatIntelligenceRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ThreatIntelligenceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
func (m *ThreatIntelligenceRequestBuilder) HostComponents()(*ThreatIntelligenceHostComponentsRequestBuilder) {
    return NewThreatIntelligenceHostComponentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// HostComponentsById provides operations to manage the hostComponents property of the microsoft.graph.security.threatIntelligence entity.
func (m *ThreatIntelligenceRequestBuilder) HostComponentsById(id string)(*ThreatIntelligenceHostComponentsHostComponentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["hostComponent%2Did"] = id
    }
    return NewThreatIntelligenceHostComponentsHostComponentItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// HostCookies provides operations to manage the hostCookies property of the microsoft.graph.security.threatIntelligence entity.
func (m *ThreatIntelligenceRequestBuilder) HostCookies()(*ThreatIntelligenceHostCookiesRequestBuilder) {
    return NewThreatIntelligenceHostCookiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// HostCookiesById provides operations to manage the hostCookies property of the microsoft.graph.security.threatIntelligence entity.
func (m *ThreatIntelligenceRequestBuilder) HostCookiesById(id string)(*ThreatIntelligenceHostCookiesHostCookieItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["hostCookie%2Did"] = id
    }
    return NewThreatIntelligenceHostCookiesHostCookieItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Hosts provides operations to manage the hosts property of the microsoft.graph.security.threatIntelligence entity.
func (m *ThreatIntelligenceRequestBuilder) Hosts()(*ThreatIntelligenceHostsRequestBuilder) {
    return NewThreatIntelligenceHostsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// HostsById provides operations to manage the hosts property of the microsoft.graph.security.threatIntelligence entity.
func (m *ThreatIntelligenceRequestBuilder) HostsById(id string)(*ThreatIntelligenceHostsHostItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["host%2Did"] = id
    }
    return NewThreatIntelligenceHostsHostItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// HostTrackers provides operations to manage the hostTrackers property of the microsoft.graph.security.threatIntelligence entity.
func (m *ThreatIntelligenceRequestBuilder) HostTrackers()(*ThreatIntelligenceHostTrackersRequestBuilder) {
    return NewThreatIntelligenceHostTrackersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// HostTrackersById provides operations to manage the hostTrackers property of the microsoft.graph.security.threatIntelligence entity.
func (m *ThreatIntelligenceRequestBuilder) HostTrackersById(id string)(*ThreatIntelligenceHostTrackersHostTrackerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["hostTracker%2Did"] = id
    }
    return NewThreatIntelligenceHostTrackersHostTrackerItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// IntelligenceProfileIndicators provides operations to manage the intelligenceProfileIndicators property of the microsoft.graph.security.threatIntelligence entity.
func (m *ThreatIntelligenceRequestBuilder) IntelligenceProfileIndicators()(*ThreatIntelligenceIntelligenceProfileIndicatorsRequestBuilder) {
    return NewThreatIntelligenceIntelligenceProfileIndicatorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IntelligenceProfileIndicatorsById provides operations to manage the intelligenceProfileIndicators property of the microsoft.graph.security.threatIntelligence entity.
func (m *ThreatIntelligenceRequestBuilder) IntelligenceProfileIndicatorsById(id string)(*ThreatIntelligenceIntelligenceProfileIndicatorsIntelligenceProfileIndicatorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["intelligenceProfileIndicator%2Did"] = id
    }
    return NewThreatIntelligenceIntelligenceProfileIndicatorsIntelligenceProfileIndicatorItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// IntelProfiles provides operations to manage the intelProfiles property of the microsoft.graph.security.threatIntelligence entity.
func (m *ThreatIntelligenceRequestBuilder) IntelProfiles()(*ThreatIntelligenceIntelProfilesRequestBuilder) {
    return NewThreatIntelligenceIntelProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IntelProfilesById provides operations to manage the intelProfiles property of the microsoft.graph.security.threatIntelligence entity.
func (m *ThreatIntelligenceRequestBuilder) IntelProfilesById(id string)(*ThreatIntelligenceIntelProfilesIntelligenceProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["intelligenceProfile%2Did"] = id
    }
    return NewThreatIntelligenceIntelProfilesIntelligenceProfileItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// PassiveDnsRecords provides operations to manage the passiveDnsRecords property of the microsoft.graph.security.threatIntelligence entity.
func (m *ThreatIntelligenceRequestBuilder) PassiveDnsRecords()(*ThreatIntelligencePassiveDnsRecordsRequestBuilder) {
    return NewThreatIntelligencePassiveDnsRecordsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PassiveDnsRecordsById provides operations to manage the passiveDnsRecords property of the microsoft.graph.security.threatIntelligence entity.
func (m *ThreatIntelligenceRequestBuilder) PassiveDnsRecordsById(id string)(*ThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["passiveDnsRecord%2Did"] = id
    }
    return NewThreatIntelligencePassiveDnsRecordsPassiveDnsRecordItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property threatIntelligence in security
func (m *ThreatIntelligenceRequestBuilder) Patch(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ThreatIntelligenceable, requestConfiguration *ThreatIntelligenceRequestBuilderPatchRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ThreatIntelligenceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// ToDeleteRequestInformation delete navigation property threatIntelligence for security
func (m *ThreatIntelligenceRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ThreatIntelligenceRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation get threatIntelligence from security
func (m *ThreatIntelligenceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatIntelligenceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property threatIntelligence in security
func (m *ThreatIntelligenceRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ThreatIntelligenceable, requestConfiguration *ThreatIntelligenceRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Vulnerabilities provides operations to manage the vulnerabilities property of the microsoft.graph.security.threatIntelligence entity.
func (m *ThreatIntelligenceRequestBuilder) Vulnerabilities()(*ThreatIntelligenceVulnerabilitiesRequestBuilder) {
    return NewThreatIntelligenceVulnerabilitiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// VulnerabilitiesById provides operations to manage the vulnerabilities property of the microsoft.graph.security.threatIntelligence entity.
func (m *ThreatIntelligenceRequestBuilder) VulnerabilitiesById(id string)(*ThreatIntelligenceVulnerabilitiesVulnerabilityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["vulnerability%2Did"] = id
    }
    return NewThreatIntelligenceVulnerabilitiesVulnerabilityItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
