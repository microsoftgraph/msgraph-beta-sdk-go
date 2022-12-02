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
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PoliciesRequestBuilderGetQueryParameters
}
// PoliciesRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PoliciesRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessReviewPolicy provides operations to manage the accessReviewPolicy property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) AccessReviewPolicy()(*PoliciesAccessReviewPolicyRequestBuilder) {
    return NewPoliciesAccessReviewPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivityBasedTimeoutPolicies provides operations to manage the activityBasedTimeoutPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) ActivityBasedTimeoutPolicies()(*PoliciesActivityBasedTimeoutPoliciesRequestBuilder) {
    return NewPoliciesActivityBasedTimeoutPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivityBasedTimeoutPoliciesById provides operations to manage the activityBasedTimeoutPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) ActivityBasedTimeoutPoliciesById(id string)(*PoliciesActivityBasedTimeoutPoliciesActivityBasedTimeoutPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["activityBasedTimeoutPolicy%2Did"] = id
    }
    return NewPoliciesActivityBasedTimeoutPoliciesActivityBasedTimeoutPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AdminConsentRequestPolicy provides operations to manage the adminConsentRequestPolicy property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) AdminConsentRequestPolicy()(*PoliciesAdminConsentRequestPolicyRequestBuilder) {
    return NewPoliciesAdminConsentRequestPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppManagementPolicies provides operations to manage the appManagementPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) AppManagementPolicies()(*PoliciesAppManagementPoliciesRequestBuilder) {
    return NewPoliciesAppManagementPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppManagementPoliciesById provides operations to manage the appManagementPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) AppManagementPoliciesById(id string)(*PoliciesAppManagementPoliciesAppManagementPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appManagementPolicy%2Did"] = id
    }
    return NewPoliciesAppManagementPoliciesAppManagementPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AuthenticationFlowsPolicy provides operations to manage the authenticationFlowsPolicy property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) AuthenticationFlowsPolicy()(*PoliciesAuthenticationFlowsPolicyRequestBuilder) {
    return NewPoliciesAuthenticationFlowsPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AuthenticationMethodsPolicy provides operations to manage the authenticationMethodsPolicy property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) AuthenticationMethodsPolicy()(*PoliciesAuthenticationMethodsPolicyRequestBuilder) {
    return NewPoliciesAuthenticationMethodsPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AuthenticationStrengthPolicies provides operations to manage the authenticationStrengthPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) AuthenticationStrengthPolicies()(*PoliciesAuthenticationStrengthPoliciesRequestBuilder) {
    return NewPoliciesAuthenticationStrengthPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AuthenticationStrengthPoliciesById provides operations to manage the authenticationStrengthPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) AuthenticationStrengthPoliciesById(id string)(*PoliciesAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["authenticationStrengthPolicy%2Did"] = id
    }
    return NewPoliciesAuthenticationStrengthPoliciesAuthenticationStrengthPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AuthorizationPolicy provides operations to manage the authorizationPolicy property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) AuthorizationPolicy()(*PoliciesAuthorizationPolicyRequestBuilder) {
    return NewPoliciesAuthorizationPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AuthorizationPolicyById provides operations to manage the authorizationPolicy property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) AuthorizationPolicyById(id string)(*PoliciesAuthorizationPolicyAuthorizationPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["authorizationPolicy%2Did"] = id
    }
    return NewPoliciesAuthorizationPolicyAuthorizationPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// B2cAuthenticationMethodsPolicy provides operations to manage the b2cAuthenticationMethodsPolicy property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) B2cAuthenticationMethodsPolicy()(*PoliciesB2cAuthenticationMethodsPolicyRequestBuilder) {
    return NewPoliciesB2cAuthenticationMethodsPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ClaimsMappingPolicies provides operations to manage the claimsMappingPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) ClaimsMappingPolicies()(*PoliciesClaimsMappingPoliciesRequestBuilder) {
    return NewPoliciesClaimsMappingPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ClaimsMappingPoliciesById provides operations to manage the claimsMappingPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) ClaimsMappingPoliciesById(id string)(*PoliciesClaimsMappingPoliciesClaimsMappingPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["claimsMappingPolicy%2Did"] = id
    }
    return NewPoliciesClaimsMappingPoliciesClaimsMappingPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ConditionalAccessPolicies provides operations to manage the conditionalAccessPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) ConditionalAccessPolicies()(*PoliciesConditionalAccessPoliciesRequestBuilder) {
    return NewPoliciesConditionalAccessPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConditionalAccessPoliciesById provides operations to manage the conditionalAccessPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) ConditionalAccessPoliciesById(id string)(*PoliciesConditionalAccessPoliciesConditionalAccessPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conditionalAccessPolicy%2Did"] = id
    }
    return NewPoliciesConditionalAccessPoliciesConditionalAccessPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
// CreateGetRequestInformation get policies
func (m *PoliciesRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *PoliciesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update policies
func (m *PoliciesRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicyRootable, requestConfiguration *PoliciesRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CrossTenantAccessPolicy provides operations to manage the crossTenantAccessPolicy property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) CrossTenantAccessPolicy()(*PoliciesCrossTenantAccessPolicyRequestBuilder) {
    return NewPoliciesCrossTenantAccessPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DefaultAppManagementPolicy provides operations to manage the defaultAppManagementPolicy property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) DefaultAppManagementPolicy()(*PoliciesDefaultAppManagementPolicyRequestBuilder) {
    return NewPoliciesDefaultAppManagementPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceRegistrationPolicy provides operations to manage the deviceRegistrationPolicy property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) DeviceRegistrationPolicy()(*PoliciesDeviceRegistrationPolicyRequestBuilder) {
    return NewPoliciesDeviceRegistrationPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DirectoryRoleAccessReviewPolicy provides operations to manage the directoryRoleAccessReviewPolicy property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) DirectoryRoleAccessReviewPolicy()(*PoliciesDirectoryRoleAccessReviewPolicyRequestBuilder) {
    return NewPoliciesDirectoryRoleAccessReviewPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExternalIdentitiesPolicy provides operations to manage the externalIdentitiesPolicy property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) ExternalIdentitiesPolicy()(*PoliciesExternalIdentitiesPolicyRequestBuilder) {
    return NewPoliciesExternalIdentitiesPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FeatureRolloutPolicies provides operations to manage the featureRolloutPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) FeatureRolloutPolicies()(*PoliciesFeatureRolloutPoliciesRequestBuilder) {
    return NewPoliciesFeatureRolloutPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FeatureRolloutPoliciesById provides operations to manage the featureRolloutPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) FeatureRolloutPoliciesById(id string)(*PoliciesFeatureRolloutPoliciesFeatureRolloutPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["featureRolloutPolicy%2Did"] = id
    }
    return NewPoliciesFeatureRolloutPoliciesFeatureRolloutPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get policies
func (m *PoliciesRequestBuilder) Get(ctx context.Context, requestConfiguration *PoliciesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicyRootable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePolicyRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicyRootable), nil
}
// HomeRealmDiscoveryPolicies provides operations to manage the homeRealmDiscoveryPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) HomeRealmDiscoveryPolicies()(*PoliciesHomeRealmDiscoveryPoliciesRequestBuilder) {
    return NewPoliciesHomeRealmDiscoveryPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// HomeRealmDiscoveryPoliciesById provides operations to manage the homeRealmDiscoveryPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) HomeRealmDiscoveryPoliciesById(id string)(*PoliciesHomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["homeRealmDiscoveryPolicy%2Did"] = id
    }
    return NewPoliciesHomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// IdentitySecurityDefaultsEnforcementPolicy provides operations to manage the identitySecurityDefaultsEnforcementPolicy property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) IdentitySecurityDefaultsEnforcementPolicy()(*PoliciesIdentitySecurityDefaultsEnforcementPolicyRequestBuilder) {
    return NewPoliciesIdentitySecurityDefaultsEnforcementPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MobileAppManagementPolicies provides operations to manage the mobileAppManagementPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) MobileAppManagementPolicies()(*PoliciesMobileAppManagementPoliciesRequestBuilder) {
    return NewPoliciesMobileAppManagementPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MobileAppManagementPoliciesById provides operations to manage the mobileAppManagementPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) MobileAppManagementPoliciesById(id string)(*PoliciesMobileAppManagementPoliciesMobilityManagementPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobilityManagementPolicy%2Did"] = id
    }
    return NewPoliciesMobileAppManagementPoliciesMobilityManagementPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MobileDeviceManagementPolicies provides operations to manage the mobileDeviceManagementPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) MobileDeviceManagementPolicies()(*PoliciesMobileDeviceManagementPoliciesRequestBuilder) {
    return NewPoliciesMobileDeviceManagementPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MobileDeviceManagementPoliciesById provides operations to manage the mobileDeviceManagementPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) MobileDeviceManagementPoliciesById(id string)(*PoliciesMobileDeviceManagementPoliciesMobilityManagementPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobilityManagementPolicy%2Did"] = id
    }
    return NewPoliciesMobileDeviceManagementPoliciesMobilityManagementPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update policies
func (m *PoliciesRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicyRootable, requestConfiguration *PoliciesRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicyRootable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePolicyRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicyRootable), nil
}
// PermissionGrantPolicies provides operations to manage the permissionGrantPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) PermissionGrantPolicies()(*PoliciesPermissionGrantPoliciesRequestBuilder) {
    return NewPoliciesPermissionGrantPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionGrantPoliciesById provides operations to manage the permissionGrantPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) PermissionGrantPoliciesById(id string)(*PoliciesPermissionGrantPoliciesPermissionGrantPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permissionGrantPolicy%2Did"] = id
    }
    return NewPoliciesPermissionGrantPoliciesPermissionGrantPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleManagementPolicies provides operations to manage the roleManagementPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) RoleManagementPolicies()(*PoliciesRoleManagementPoliciesRequestBuilder) {
    return NewPoliciesRoleManagementPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleManagementPoliciesById provides operations to manage the roleManagementPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) RoleManagementPoliciesById(id string)(*PoliciesRoleManagementPoliciesUnifiedRoleManagementPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleManagementPolicy%2Did"] = id
    }
    return NewPoliciesRoleManagementPoliciesUnifiedRoleManagementPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleManagementPolicyAssignments provides operations to manage the roleManagementPolicyAssignments property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) RoleManagementPolicyAssignments()(*PoliciesRoleManagementPolicyAssignmentsRequestBuilder) {
    return NewPoliciesRoleManagementPolicyAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleManagementPolicyAssignmentsById provides operations to manage the roleManagementPolicyAssignments property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) RoleManagementPolicyAssignmentsById(id string)(*PoliciesRoleManagementPolicyAssignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unifiedRoleManagementPolicyAssignment%2Did"] = id
    }
    return NewPoliciesRoleManagementPolicyAssignmentsUnifiedRoleManagementPolicyAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ServicePrincipalCreationPolicies provides operations to manage the servicePrincipalCreationPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) ServicePrincipalCreationPolicies()(*PoliciesServicePrincipalCreationPoliciesRequestBuilder) {
    return NewPoliciesServicePrincipalCreationPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipalCreationPoliciesById provides operations to manage the servicePrincipalCreationPolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) ServicePrincipalCreationPoliciesById(id string)(*PoliciesServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["servicePrincipalCreationPolicy%2Did"] = id
    }
    return NewPoliciesServicePrincipalCreationPoliciesServicePrincipalCreationPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TokenIssuancePolicies provides operations to manage the tokenIssuancePolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) TokenIssuancePolicies()(*PoliciesTokenIssuancePoliciesRequestBuilder) {
    return NewPoliciesTokenIssuancePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TokenIssuancePoliciesById provides operations to manage the tokenIssuancePolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) TokenIssuancePoliciesById(id string)(*PoliciesTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tokenIssuancePolicy%2Did"] = id
    }
    return NewPoliciesTokenIssuancePoliciesTokenIssuancePolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TokenLifetimePolicies provides operations to manage the tokenLifetimePolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) TokenLifetimePolicies()(*PoliciesTokenLifetimePoliciesRequestBuilder) {
    return NewPoliciesTokenLifetimePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TokenLifetimePoliciesById provides operations to manage the tokenLifetimePolicies property of the microsoft.graph.policyRoot entity.
func (m *PoliciesRequestBuilder) TokenLifetimePoliciesById(id string)(*PoliciesTokenLifetimePoliciesTokenLifetimePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tokenLifetimePolicy%2Did"] = id
    }
    return NewPoliciesTokenLifetimePoliciesTokenLifetimePolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
