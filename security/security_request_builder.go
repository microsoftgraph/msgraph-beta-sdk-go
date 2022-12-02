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
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SecurityRequestBuilderGetQueryParameters
}
// SecurityRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurityRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Alerts provides operations to manage the alerts property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) Alerts()(*SecurityAlertsRequestBuilder) {
    return NewSecurityAlertsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Alerts_v2 provides operations to manage the alerts_v2 property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) Alerts_v2()(*SecurityAlerts_v2RequestBuilder) {
    return NewSecurityAlerts_v2RequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Alerts_v2ById provides operations to manage the alerts_v2 property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) Alerts_v2ById(id string)(*SecurityAlerts_v2AlertItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["alert%2Did"] = id
    }
    return NewSecurityAlerts_v2AlertItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AlertsById provides operations to manage the alerts property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) AlertsById(id string)(*SecurityAlertsAlertItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["alert%2Did"] = id
    }
    return NewSecurityAlertsAlertItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AttackSimulation provides operations to manage the attackSimulation property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) AttackSimulation()(*SecurityAttackSimulationRequestBuilder) {
    return NewSecurityAttackSimulationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cases provides operations to manage the cases property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) Cases()(*SecurityCasesRequestBuilder) {
    return NewSecurityCasesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CloudAppSecurityProfiles provides operations to manage the cloudAppSecurityProfiles property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) CloudAppSecurityProfiles()(*SecurityCloudAppSecurityProfilesRequestBuilder) {
    return NewSecurityCloudAppSecurityProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CloudAppSecurityProfilesById provides operations to manage the cloudAppSecurityProfiles property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) CloudAppSecurityProfilesById(id string)(*SecurityCloudAppSecurityProfilesCloudAppSecurityProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudAppSecurityProfile%2Did"] = id
    }
    return NewSecurityCloudAppSecurityProfilesCloudAppSecurityProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSecurityRequestBuilder instantiates a new SecurityRequestBuilder and sets the default values.
func NewSecurityRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSecurityRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get security
func (m *SecurityRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *SecurityRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update security
func (m *SecurityRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Securityable, requestConfiguration *SecurityRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// DomainSecurityProfiles provides operations to manage the domainSecurityProfiles property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) DomainSecurityProfiles()(*SecurityDomainSecurityProfilesRequestBuilder) {
    return NewSecurityDomainSecurityProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DomainSecurityProfilesById provides operations to manage the domainSecurityProfiles property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) DomainSecurityProfilesById(id string)(*SecurityDomainSecurityProfilesDomainSecurityProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["domainSecurityProfile%2Did"] = id
    }
    return NewSecurityDomainSecurityProfilesDomainSecurityProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// FileSecurityProfiles provides operations to manage the fileSecurityProfiles property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) FileSecurityProfiles()(*SecurityFileSecurityProfilesRequestBuilder) {
    return NewSecurityFileSecurityProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FileSecurityProfilesById provides operations to manage the fileSecurityProfiles property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) FileSecurityProfilesById(id string)(*SecurityFileSecurityProfilesFileSecurityProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["fileSecurityProfile%2Did"] = id
    }
    return NewSecurityFileSecurityProfilesFileSecurityProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get security
func (m *SecurityRequestBuilder) Get(ctx context.Context, requestConfiguration *SecurityRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Securityable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSecurityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Securityable), nil
}
// HostSecurityProfiles provides operations to manage the hostSecurityProfiles property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) HostSecurityProfiles()(*SecurityHostSecurityProfilesRequestBuilder) {
    return NewSecurityHostSecurityProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// HostSecurityProfilesById provides operations to manage the hostSecurityProfiles property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) HostSecurityProfilesById(id string)(*SecurityHostSecurityProfilesHostSecurityProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["hostSecurityProfile%2Did"] = id
    }
    return NewSecurityHostSecurityProfilesHostSecurityProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Incidents provides operations to manage the incidents property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) Incidents()(*SecurityIncidentsRequestBuilder) {
    return NewSecurityIncidentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncidentsById provides operations to manage the incidents property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) IncidentsById(id string)(*SecurityIncidentsIncidentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["incident%2Did"] = id
    }
    return NewSecurityIncidentsIncidentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// InformationProtection provides operations to manage the informationProtection property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) InformationProtection()(*SecurityInformationProtectionRequestBuilder) {
    return NewSecurityInformationProtectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IpSecurityProfiles provides operations to manage the ipSecurityProfiles property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) IpSecurityProfiles()(*SecurityIpSecurityProfilesRequestBuilder) {
    return NewSecurityIpSecurityProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IpSecurityProfilesById provides operations to manage the ipSecurityProfiles property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) IpSecurityProfilesById(id string)(*SecurityIpSecurityProfilesIpSecurityProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["ipSecurityProfile%2Did"] = id
    }
    return NewSecurityIpSecurityProfilesIpSecurityProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Labels provides operations to manage the labels property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) Labels()(*SecurityLabelsRequestBuilder) {
    return NewSecurityLabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update security
func (m *SecurityRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Securityable, requestConfiguration *SecurityRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Securityable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSecurityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Securityable), nil
}
// ProviderTenantSettings provides operations to manage the providerTenantSettings property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) ProviderTenantSettings()(*SecurityProviderTenantSettingsRequestBuilder) {
    return NewSecurityProviderTenantSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ProviderTenantSettingsById provides operations to manage the providerTenantSettings property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) ProviderTenantSettingsById(id string)(*SecurityProviderTenantSettingsProviderTenantSettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["providerTenantSetting%2Did"] = id
    }
    return NewSecurityProviderTenantSettingsProviderTenantSettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RunHuntingQuery provides operations to call the runHuntingQuery method.
func (m *SecurityRequestBuilder) RunHuntingQuery()(*SecurityRunHuntingQueryRequestBuilder) {
    return NewSecurityRunHuntingQueryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SecureScoreControlProfiles provides operations to manage the secureScoreControlProfiles property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) SecureScoreControlProfiles()(*SecuritySecureScoreControlProfilesRequestBuilder) {
    return NewSecuritySecureScoreControlProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SecureScoreControlProfilesById provides operations to manage the secureScoreControlProfiles property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) SecureScoreControlProfilesById(id string)(*SecuritySecureScoreControlProfilesSecureScoreControlProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["secureScoreControlProfile%2Did"] = id
    }
    return NewSecuritySecureScoreControlProfilesSecureScoreControlProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SecureScores provides operations to manage the secureScores property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) SecureScores()(*SecuritySecureScoresRequestBuilder) {
    return NewSecuritySecureScoresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SecureScoresById provides operations to manage the secureScores property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) SecureScoresById(id string)(*SecuritySecureScoresSecureScoreItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["secureScore%2Did"] = id
    }
    return NewSecuritySecureScoresSecureScoreItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SecurityActions provides operations to manage the securityActions property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) SecurityActions()(*SecuritySecurityActionsRequestBuilder) {
    return NewSecuritySecurityActionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SecurityActionsById provides operations to manage the securityActions property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) SecurityActionsById(id string)(*SecuritySecurityActionsSecurityActionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["securityAction%2Did"] = id
    }
    return NewSecuritySecurityActionsSecurityActionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SubjectRightsRequests provides operations to manage the subjectRightsRequests property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) SubjectRightsRequests()(*SecuritySubjectRightsRequestsRequestBuilder) {
    return NewSecuritySubjectRightsRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubjectRightsRequestsById provides operations to manage the subjectRightsRequests property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) SubjectRightsRequestsById(id string)(*SecuritySubjectRightsRequestsSubjectRightsRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subjectRightsRequest%2Did"] = id
    }
    return NewSecuritySubjectRightsRequestsSubjectRightsRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ThreatSubmission provides operations to manage the threatSubmission property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) ThreatSubmission()(*SecurityThreatSubmissionRequestBuilder) {
    return NewSecurityThreatSubmissionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TiIndicators provides operations to manage the tiIndicators property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) TiIndicators()(*SecurityTiIndicatorsRequestBuilder) {
    return NewSecurityTiIndicatorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TiIndicatorsById provides operations to manage the tiIndicators property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) TiIndicatorsById(id string)(*SecurityTiIndicatorsTiIndicatorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tiIndicator%2Did"] = id
    }
    return NewSecurityTiIndicatorsTiIndicatorItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Triggers provides operations to manage the triggers property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) Triggers()(*SecurityTriggersRequestBuilder) {
    return NewSecurityTriggersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TriggerTypes provides operations to manage the triggerTypes property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) TriggerTypes()(*SecurityTriggerTypesRequestBuilder) {
    return NewSecurityTriggerTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserSecurityProfiles provides operations to manage the userSecurityProfiles property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) UserSecurityProfiles()(*SecurityUserSecurityProfilesRequestBuilder) {
    return NewSecurityUserSecurityProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserSecurityProfilesById provides operations to manage the userSecurityProfiles property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) UserSecurityProfilesById(id string)(*SecurityUserSecurityProfilesUserSecurityProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userSecurityProfile%2Did"] = id
    }
    return NewSecurityUserSecurityProfilesUserSecurityProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
