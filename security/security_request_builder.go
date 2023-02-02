package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SecurityRequestBuilder provides operations to manage the security singleton.
type SecurityRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SecurityRequestBuilderGetQueryParameters get security
type SecurityRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SecurityRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurityRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SecurityRequestBuilderGetQueryParameters
}
// SecurityRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurityRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Alerts provides operations to manage the alerts property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) Alerts()(*AlertsRequestBuilder) {
    return NewAlertsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Alerts_v2 provides operations to manage the alerts_v2 property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) Alerts_v2()(*Alerts_v2RequestBuilder) {
    return NewAlerts_v2RequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Alerts_v2ById provides operations to manage the alerts_v2 property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) Alerts_v2ById(id string)(*Alerts_v2AlertItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["alert%2Did"] = id
    }
    return NewAlerts_v2AlertItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// AlertsById provides operations to manage the alerts property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) AlertsById(id string)(*AlertsAlertItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["alert%2Did"] = id
    }
    return NewAlertsAlertItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// AttackSimulation provides operations to manage the attackSimulation property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) AttackSimulation()(*AttackSimulationRequestBuilder) {
    return NewAttackSimulationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Cases provides operations to manage the cases property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) Cases()(*CasesRequestBuilder) {
    return NewCasesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CloudAppSecurityProfiles provides operations to manage the cloudAppSecurityProfiles property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) CloudAppSecurityProfiles()(*CloudAppSecurityProfilesRequestBuilder) {
    return NewCloudAppSecurityProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CloudAppSecurityProfilesById provides operations to manage the cloudAppSecurityProfiles property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) CloudAppSecurityProfilesById(id string)(*CloudAppSecurityProfilesCloudAppSecurityProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudAppSecurityProfile%2Did"] = id
    }
    return NewCloudAppSecurityProfilesCloudAppSecurityProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewSecurityRequestBuilderInternal instantiates a new SecurityRequestBuilder and sets the default values.
func NewSecurityRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityRequestBuilder) {
    m := &SecurityRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewSecurityRequestBuilder instantiates a new SecurityRequestBuilder and sets the default values.
func NewSecurityRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSecurityRequestBuilderInternal(urlParams, requestAdapter)
}
// DomainSecurityProfiles provides operations to manage the domainSecurityProfiles property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) DomainSecurityProfiles()(*DomainSecurityProfilesRequestBuilder) {
    return NewDomainSecurityProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DomainSecurityProfilesById provides operations to manage the domainSecurityProfiles property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) DomainSecurityProfilesById(id string)(*DomainSecurityProfilesDomainSecurityProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["domainSecurityProfile%2Did"] = id
    }
    return NewDomainSecurityProfilesDomainSecurityProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// FileSecurityProfiles provides operations to manage the fileSecurityProfiles property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) FileSecurityProfiles()(*FileSecurityProfilesRequestBuilder) {
    return NewFileSecurityProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// FileSecurityProfilesById provides operations to manage the fileSecurityProfiles property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) FileSecurityProfilesById(id string)(*FileSecurityProfilesFileSecurityProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["fileSecurityProfile%2Did"] = id
    }
    return NewFileSecurityProfilesFileSecurityProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get get security
func (m *SecurityRequestBuilder) Get(ctx context.Context, requestConfiguration *SecurityRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Securityable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSecurityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Securityable), nil
}
// HostSecurityProfiles provides operations to manage the hostSecurityProfiles property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) HostSecurityProfiles()(*HostSecurityProfilesRequestBuilder) {
    return NewHostSecurityProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// HostSecurityProfilesById provides operations to manage the hostSecurityProfiles property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) HostSecurityProfilesById(id string)(*HostSecurityProfilesHostSecurityProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["hostSecurityProfile%2Did"] = id
    }
    return NewHostSecurityProfilesHostSecurityProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Incidents provides operations to manage the incidents property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) Incidents()(*IncidentsRequestBuilder) {
    return NewIncidentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// IncidentsById provides operations to manage the incidents property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) IncidentsById(id string)(*IncidentsIncidentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["incident%2Did"] = id
    }
    return NewIncidentsIncidentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// InformationProtection provides operations to manage the informationProtection property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) InformationProtection()(*InformationProtectionRequestBuilder) {
    return NewInformationProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// IpSecurityProfiles provides operations to manage the ipSecurityProfiles property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) IpSecurityProfiles()(*IpSecurityProfilesRequestBuilder) {
    return NewIpSecurityProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// IpSecurityProfilesById provides operations to manage the ipSecurityProfiles property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) IpSecurityProfilesById(id string)(*IpSecurityProfilesIpSecurityProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["ipSecurityProfile%2Did"] = id
    }
    return NewIpSecurityProfilesIpSecurityProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Labels provides operations to manage the labels property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) Labels()(*LabelsRequestBuilder) {
    return NewLabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSecurityRunHuntingQuery provides operations to call the runHuntingQuery method.
func (m *SecurityRequestBuilder) MicrosoftGraphSecurityRunHuntingQuery()(*MicrosoftGraphSecurityRunHuntingQueryRunHuntingQueryRequestBuilder) {
    return NewMicrosoftGraphSecurityRunHuntingQueryRunHuntingQueryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update security
func (m *SecurityRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Securityable, requestConfiguration *SecurityRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Securityable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSecurityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Securityable), nil
}
// ProviderTenantSettings provides operations to manage the providerTenantSettings property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) ProviderTenantSettings()(*ProviderTenantSettingsRequestBuilder) {
    return NewProviderTenantSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ProviderTenantSettingsById provides operations to manage the providerTenantSettings property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) ProviderTenantSettingsById(id string)(*ProviderTenantSettingsProviderTenantSettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["providerTenantSetting%2Did"] = id
    }
    return NewProviderTenantSettingsProviderTenantSettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// SecureScoreControlProfiles provides operations to manage the secureScoreControlProfiles property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) SecureScoreControlProfiles()(*SecureScoreControlProfilesRequestBuilder) {
    return NewSecureScoreControlProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SecureScoreControlProfilesById provides operations to manage the secureScoreControlProfiles property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) SecureScoreControlProfilesById(id string)(*SecureScoreControlProfilesSecureScoreControlProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["secureScoreControlProfile%2Did"] = id
    }
    return NewSecureScoreControlProfilesSecureScoreControlProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// SecureScores provides operations to manage the secureScores property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) SecureScores()(*SecureScoresRequestBuilder) {
    return NewSecureScoresRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SecureScoresById provides operations to manage the secureScores property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) SecureScoresById(id string)(*SecureScoresSecureScoreItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["secureScore%2Did"] = id
    }
    return NewSecureScoresSecureScoreItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// SecurityActions provides operations to manage the securityActions property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) SecurityActions()(*SecurityActionsRequestBuilder) {
    return NewSecurityActionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SecurityActionsById provides operations to manage the securityActions property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) SecurityActionsById(id string)(*SecurityActionsSecurityActionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["securityAction%2Did"] = id
    }
    return NewSecurityActionsSecurityActionItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// SubjectRightsRequests provides operations to manage the subjectRightsRequests property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) SubjectRightsRequests()(*SubjectRightsRequestsRequestBuilder) {
    return NewSubjectRightsRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SubjectRightsRequestsById provides operations to manage the subjectRightsRequests property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) SubjectRightsRequestsById(id string)(*SubjectRightsRequestsSubjectRightsRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subjectRightsRequest%2Did"] = id
    }
    return NewSubjectRightsRequestsSubjectRightsRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ThreatSubmission provides operations to manage the threatSubmission property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) ThreatSubmission()(*ThreatSubmissionRequestBuilder) {
    return NewThreatSubmissionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TiIndicators provides operations to manage the tiIndicators property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) TiIndicators()(*TiIndicatorsRequestBuilder) {
    return NewTiIndicatorsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TiIndicatorsById provides operations to manage the tiIndicators property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) TiIndicatorsById(id string)(*TiIndicatorsTiIndicatorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tiIndicator%2Did"] = id
    }
    return NewTiIndicatorsTiIndicatorItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToGetRequestInformation get security
func (m *SecurityRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SecurityRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
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
// ToPatchRequestInformation update security
func (m *SecurityRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Securityable, requestConfiguration *SecurityRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Triggers provides operations to manage the triggers property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) Triggers()(*TriggersRequestBuilder) {
    return NewTriggersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TriggerTypes provides operations to manage the triggerTypes property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) TriggerTypes()(*TriggerTypesRequestBuilder) {
    return NewTriggerTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserSecurityProfiles provides operations to manage the userSecurityProfiles property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) UserSecurityProfiles()(*UserSecurityProfilesRequestBuilder) {
    return NewUserSecurityProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserSecurityProfilesById provides operations to manage the userSecurityProfiles property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) UserSecurityProfilesById(id string)(*UserSecurityProfilesUserSecurityProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userSecurityProfile%2Did"] = id
    }
    return NewUserSecurityProfilesUserSecurityProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
