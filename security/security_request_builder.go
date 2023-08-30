package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SecurityRequestBuilder provides operations to manage the security singleton.
type SecurityRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
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
    return NewAlertsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Alerts_v2 provides operations to manage the alerts_v2 property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) Alerts_v2()(*Alerts_v2RequestBuilder) {
    return NewAlerts_v2RequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AttackSimulation provides operations to manage the attackSimulation property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) AttackSimulation()(*AttackSimulationRequestBuilder) {
    return NewAttackSimulationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cases provides operations to manage the cases property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) Cases()(*CasesRequestBuilder) {
    return NewCasesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CloudAppSecurityProfiles provides operations to manage the cloudAppSecurityProfiles property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) CloudAppSecurityProfiles()(*CloudAppSecurityProfilesRequestBuilder) {
    return NewCloudAppSecurityProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewSecurityRequestBuilderInternal instantiates a new SecurityRequestBuilder and sets the default values.
func NewSecurityRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityRequestBuilder) {
    m := &SecurityRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security{?%24select,%24expand}", pathParameters),
    }
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
    return NewDomainSecurityProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FileSecurityProfiles provides operations to manage the fileSecurityProfiles property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) FileSecurityProfiles()(*FileSecurityProfilesRequestBuilder) {
    return NewFileSecurityProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
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
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSecurityFromDiscriminatorValue, errorMapping)
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
    return NewHostSecurityProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Incidents provides operations to manage the incidents property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) Incidents()(*IncidentsRequestBuilder) {
    return NewIncidentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// InformationProtection provides operations to manage the informationProtection property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) InformationProtection()(*InformationProtectionRequestBuilder) {
    return NewInformationProtectionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IpSecurityProfiles provides operations to manage the ipSecurityProfiles property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) IpSecurityProfiles()(*IpSecurityProfilesRequestBuilder) {
    return NewIpSecurityProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Labels provides operations to manage the labels property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) Labels()(*LabelsRequestBuilder) {
    return NewLabelsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphSecurityRunHuntingQuery provides operations to call the runHuntingQuery method.
func (m *SecurityRequestBuilder) MicrosoftGraphSecurityRunHuntingQuery()(*MicrosoftGraphSecurityRunHuntingQueryRequestBuilder) {
    return NewMicrosoftGraphSecurityRunHuntingQueryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
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
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSecurityFromDiscriminatorValue, errorMapping)
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
    return NewProviderTenantSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SecureScoreControlProfiles provides operations to manage the secureScoreControlProfiles property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) SecureScoreControlProfiles()(*SecureScoreControlProfilesRequestBuilder) {
    return NewSecureScoreControlProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SecureScores provides operations to manage the secureScores property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) SecureScores()(*SecureScoresRequestBuilder) {
    return NewSecureScoresRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SecurityActions provides operations to manage the securityActions property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) SecurityActions()(*SecurityActionsRequestBuilder) {
    return NewSecurityActionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SubjectRightsRequests provides operations to manage the subjectRightsRequests property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) SubjectRightsRequests()(*SubjectRightsRequestsRequestBuilder) {
    return NewSubjectRightsRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ThreatIntelligence provides operations to manage the threatIntelligence property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) ThreatIntelligence()(*ThreatIntelligenceRequestBuilder) {
    return NewThreatIntelligenceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ThreatSubmission provides operations to manage the threatSubmission property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) ThreatSubmission()(*ThreatSubmissionRequestBuilder) {
    return NewThreatSubmissionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TiIndicators provides operations to manage the tiIndicators property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) TiIndicators()(*TiIndicatorsRequestBuilder) {
    return NewTiIndicatorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get security
func (m *SecurityRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SecurityRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update security
func (m *SecurityRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Securityable, requestConfiguration *SecurityRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Triggers provides operations to manage the triggers property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) Triggers()(*TriggersRequestBuilder) {
    return NewTriggersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TriggerTypes provides operations to manage the triggerTypes property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) TriggerTypes()(*TriggerTypesRequestBuilder) {
    return NewTriggerTypesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserSecurityProfiles provides operations to manage the userSecurityProfiles property of the microsoft.graph.security entity.
func (m *SecurityRequestBuilder) UserSecurityProfiles()(*UserSecurityProfilesRequestBuilder) {
    return NewUserSecurityProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *SecurityRequestBuilder) WithUrl(rawUrl string)(*SecurityRequestBuilder) {
    return NewSecurityRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
