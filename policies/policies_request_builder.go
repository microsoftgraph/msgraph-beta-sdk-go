package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PoliciesRequestBuilder provides operations to manage the policyRoot singleton.
type PoliciesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PoliciesRequestBuilderGetQueryParameters get policies
type PoliciesRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PoliciesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PoliciesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PoliciesRequestBuilderGetQueryParameters
}
// PoliciesRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PoliciesRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessReviewPolicy provides operations to manage the accessReviewPolicy property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) AccessReviewPolicy()(*AccessReviewPolicyRequestBuilder) {
    return NewAccessReviewPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivityBasedTimeoutPolicies provides operations to manage the activityBasedTimeoutPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) ActivityBasedTimeoutPolicies()(*ActivityBasedTimeoutPoliciesRequestBuilder) {
    return NewActivityBasedTimeoutPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivityBasedTimeoutPoliciesById provides operations to manage the activityBasedTimeoutPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) ActivityBasedTimeoutPoliciesById(id string)(*ActivityBasedTimeoutPoliciesActivityBasedTimeoutPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["activityBasedTimeoutPolicy%2Did"] = id
    }
    return NewActivityBasedTimeoutPoliciesActivityBasedTimeoutPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AdminConsentRequestPolicy provides operations to manage the adminConsentRequestPolicy property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) AdminConsentRequestPolicy()(*AdminConsentRequestPolicyRequestBuilder) {
    return NewAdminConsentRequestPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppManagementPolicies provides operations to manage the appManagementPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) AppManagementPolicies()(*AppManagementPoliciesRequestBuilder) {
    return NewAppManagementPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppManagementPoliciesById provides operations to manage the appManagementPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) AppManagementPoliciesById(id string)(*AppManagementPoliciesAppManagementPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appManagementPolicy%2Did"] = id
    }
    return NewAppManagementPoliciesAppManagementPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AuthenticationFlowsPolicy provides operations to manage the authenticationFlowsPolicy property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) AuthenticationFlowsPolicy()(*AuthenticationFlowsPolicyRequestBuilder) {
    return NewAuthenticationFlowsPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AuthenticationMethodsPolicy provides operations to manage the authenticationMethodsPolicy property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) AuthenticationMethodsPolicy()(*AuthenticationMethodsPolicyRequestBuilder) {
    return NewAuthenticationMethodsPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AuthenticationStrengthPolicies provides operations to manage the authenticationStrengthPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) AuthenticationStrengthPolicies()(*AuthenticationStrengthPoliciesRequestBuilder) {
    return NewAuthenticationStrengthPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AuthenticationStrengthPoliciesById provides operations to manage the authenticationStrengthPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) AuthenticationStrengthPoliciesById(id string)(*AuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["authenticationStrengthPolicy%2Did"] = id
    }
    return NewAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AuthorizationPolicy provides operations to manage the authorizationPolicy property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) AuthorizationPolicy()(*AuthorizationPolicyRequestBuilder) {
    return NewAuthorizationPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AuthorizationPolicyById provides operations to manage the authorizationPolicy property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) AuthorizationPolicyById(id string)(*AuthorizationPolicyAuthorizationPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["authorizationPolicy%2Did"] = id
    }
    return NewAuthorizationPolicyAuthorizationPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// B2cAuthenticationMethodsPolicy provides operations to manage the b2cAuthenticationMethodsPolicy property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) B2cAuthenticationMethodsPolicy()(*B2cAuthenticationMethodsPolicyRequestBuilder) {
    return NewB2cAuthenticationMethodsPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ClaimsMappingPolicies provides operations to manage the claimsMappingPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) ClaimsMappingPolicies()(*ClaimsMappingPoliciesRequestBuilder) {
    return NewClaimsMappingPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ClaimsMappingPoliciesById provides operations to manage the claimsMappingPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) ClaimsMappingPoliciesById(id string)(*ClaimsMappingPoliciesClaimsMappingPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["claimsMappingPolicy%2Did"] = id
    }
    return NewClaimsMappingPoliciesClaimsMappingPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ConditionalAccessPolicies provides operations to manage the conditionalAccessPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) ConditionalAccessPolicies()(*ConditionalAccessPoliciesRequestBuilder) {
    return NewConditionalAccessPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConditionalAccessPoliciesById provides operations to manage the conditionalAccessPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) ConditionalAccessPoliciesById(id string)(*ConditionalAccessPoliciesConditionalAccessPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conditionalAccessPolicy%2Did"] = id
    }
    return NewConditionalAccessPoliciesConditionalAccessPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewPoliciesRequestBuilderInternal instantiates a new PoliciesRequestBuilder and sets the default values.
func NewPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PoliciesRequestBuilder) {
    m := &PoliciesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/policies{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPoliciesRequestBuilder instantiates a new PoliciesRequestBuilder and sets the default values.
func NewPoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// CrossTenantAccessPolicy provides operations to manage the crossTenantAccessPolicy property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) CrossTenantAccessPolicy()(*CrossTenantAccessPolicyRequestBuilder) {
    return NewCrossTenantAccessPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DefaultAppManagementPolicy provides operations to manage the defaultAppManagementPolicy property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) DefaultAppManagementPolicy()(*DefaultAppManagementPolicyRequestBuilder) {
    return NewDefaultAppManagementPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceRegistrationPolicy provides operations to manage the deviceRegistrationPolicy property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) DeviceRegistrationPolicy()(*DeviceRegistrationPolicyRequestBuilder) {
    return NewDeviceRegistrationPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DirectoryRoleAccessReviewPolicy provides operations to manage the directoryRoleAccessReviewPolicy property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) DirectoryRoleAccessReviewPolicy()(*DirectoryRoleAccessReviewPolicyRequestBuilder) {
    return NewDirectoryRoleAccessReviewPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExternalIdentitiesPolicy provides operations to manage the externalIdentitiesPolicy property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) ExternalIdentitiesPolicy()(*ExternalIdentitiesPolicyRequestBuilder) {
    return NewExternalIdentitiesPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FeatureRolloutPolicies provides operations to manage the featureRolloutPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) FeatureRolloutPolicies()(*FeatureRolloutPoliciesRequestBuilder) {
    return NewFeatureRolloutPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FeatureRolloutPoliciesById provides operations to manage the featureRolloutPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) FeatureRolloutPoliciesById(id string)(*FeatureRolloutPoliciesFeatureRolloutPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["featureRolloutPolicy%2Did"] = id
    }
    return NewFeatureRolloutPoliciesFeatureRolloutPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get policies
func (m *PoliciesRequestBuilder) Get(ctx context.Context, requestConfiguration *PoliciesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicyRootable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePolicyRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicyRootable), nil
}
// HomeRealmDiscoveryPolicies provides operations to manage the homeRealmDiscoveryPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) HomeRealmDiscoveryPolicies()(*HomeRealmDiscoveryPoliciesRequestBuilder) {
    return NewHomeRealmDiscoveryPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// HomeRealmDiscoveryPoliciesById provides operations to manage the homeRealmDiscoveryPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) HomeRealmDiscoveryPoliciesById(id string)(*HomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["homeRealmDiscoveryPolicy%2Did"] = id
    }
    return NewHomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// IdentitySecurityDefaultsEnforcementPolicy provides operations to manage the identitySecurityDefaultsEnforcementPolicy property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) IdentitySecurityDefaultsEnforcementPolicy()(*IdentitySecurityDefaultsEnforcementPolicyRequestBuilder) {
    return NewIdentitySecurityDefaultsEnforcementPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MobileAppManagementPolicies provides operations to manage the mobileAppManagementPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) MobileAppManagementPolicies()(*MobileAppManagementPoliciesRequestBuilder) {
    return NewMobileAppManagementPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MobileAppManagementPoliciesById provides operations to manage the mobileAppManagementPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) MobileAppManagementPoliciesById(id string)(*MobileAppManagementPoliciesMobilityManagementPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobilityManagementPolicy%2Did"] = id
    }
    return NewMobileAppManagementPoliciesMobilityManagementPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MobileDeviceManagementPolicies provides operations to manage the mobileDeviceManagementPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) MobileDeviceManagementPolicies()(*MobileDeviceManagementPoliciesRequestBuilder) {
    return NewMobileDeviceManagementPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MobileDeviceManagementPoliciesById provides operations to manage the mobileDeviceManagementPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) MobileDeviceManagementPoliciesById(id string)(*MobileDeviceManagementPoliciesMobilityManagementPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobilityManagementPolicy%2Did"] = id
    }
    return NewMobileDeviceManagementPoliciesMobilityManagementPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update policies
func (m *PoliciesRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicyRootable, requestConfiguration *PoliciesRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicyRootable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePolicyRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicyRootable), nil
}
// PermissionGrantPolicies provides operations to manage the permissionGrantPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) PermissionGrantPolicies()(*PermissionGrantPoliciesRequestBuilder) {
    return NewPermissionGrantPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionGrantPoliciesById provides operations to manage the permissionGrantPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) PermissionGrantPoliciesById(id string)(*PermissionGrantPoliciesPermissionGrantPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permissionGrantPolicy%2Did"] = id
    }
    return NewPermissionGrantPoliciesPermissionGrantPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleManagementPolicies provides operations to manage the roleManagementPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) RoleManagementPolicies()(*RoleManagementPoliciesRequestBuilder) {
    return NewRoleManagementPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleManagementPoliciesById provides operations to manage the roleManagementPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) RoleManagementPoliciesById(id string)(*RoleManagementPoliciesUnifiedRoleManagementPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleManagementPolicy%2Did"] = id
    }
    return NewRoleManagementPoliciesUnifiedRoleManagementPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleManagementPolicyAssignments provides operations to manage the roleManagementPolicyAssignments property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) RoleManagementPolicyAssignments()(*RoleManagementPolicyAssignmentsRequestBuilder) {
    return NewRoleManagementPolicyAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleManagementPolicyAssignmentsById provides operations to manage the roleManagementPolicyAssignments property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) RoleManagementPolicyAssignmentsById(id string)(*RoleManagementPolicyAssignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleManagementPolicyAssignment%2Did"] = id
    }
    return NewRoleManagementPolicyAssignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ServicePrincipalCreationPolicies provides operations to manage the servicePrincipalCreationPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) ServicePrincipalCreationPolicies()(*ServicePrincipalCreationPoliciesRequestBuilder) {
    return NewServicePrincipalCreationPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipalCreationPoliciesById provides operations to manage the servicePrincipalCreationPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) ServicePrincipalCreationPoliciesById(id string)(*ServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["servicePrincipalCreationPolicy%2Did"] = id
    }
    return NewServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ToGetRequestInformation get policies
func (m *PoliciesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PoliciesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// TokenIssuancePolicies provides operations to manage the tokenIssuancePolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) TokenIssuancePolicies()(*TokenIssuancePoliciesRequestBuilder) {
    return NewTokenIssuancePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TokenIssuancePoliciesById provides operations to manage the tokenIssuancePolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) TokenIssuancePoliciesById(id string)(*TokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tokenIssuancePolicy%2Did"] = id
    }
    return NewTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TokenLifetimePolicies provides operations to manage the tokenLifetimePolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) TokenLifetimePolicies()(*TokenLifetimePoliciesRequestBuilder) {
    return NewTokenLifetimePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TokenLifetimePoliciesById provides operations to manage the tokenLifetimePolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) TokenLifetimePoliciesById(id string)(*TokenLifetimePoliciesTokenLifetimePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tokenLifetimePolicy%2Did"] = id
    }
    return NewTokenLifetimePoliciesTokenLifetimePolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ToPatchRequestInformation update policies
func (m *PoliciesRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicyRootable, requestConfiguration *PoliciesRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
