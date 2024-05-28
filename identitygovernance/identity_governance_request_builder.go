package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentityGovernanceRequestBuilder provides operations to manage the identityGovernance singleton.
type IdentityGovernanceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IdentityGovernanceRequestBuilderGetQueryParameters get identityGovernance
type IdentityGovernanceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityGovernanceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityGovernanceRequestBuilderGetQueryParameters
}
// IdentityGovernanceRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessReviews provides operations to manage the accessReviews property of the microsoft.graph.identityGovernance entity.
// returns a *AccessreviewsAccessReviewsRequestBuilder when successful
func (m *IdentityGovernanceRequestBuilder) AccessReviews()(*AccessreviewsAccessReviewsRequestBuilder) {
    return NewAccessreviewsAccessReviewsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AppConsent provides operations to manage the appConsent property of the microsoft.graph.identityGovernance entity.
// returns a *AppconsentAppConsentRequestBuilder when successful
func (m *IdentityGovernanceRequestBuilder) AppConsent()(*AppconsentAppConsentRequestBuilder) {
    return NewAppconsentAppConsentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewIdentityGovernanceRequestBuilderInternal instantiates a new IdentityGovernanceRequestBuilder and sets the default values.
func NewIdentityGovernanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceRequestBuilder) {
    m := &IdentityGovernanceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewIdentityGovernanceRequestBuilder instantiates a new IdentityGovernanceRequestBuilder and sets the default values.
func NewIdentityGovernanceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceRequestBuilderInternal(urlParams, requestAdapter)
}
// EntitlementManagement provides operations to manage the entitlementManagement property of the microsoft.graph.identityGovernance entity.
// returns a *EntitlementmanagementEntitlementManagementRequestBuilder when successful
func (m *IdentityGovernanceRequestBuilder) EntitlementManagement()(*EntitlementmanagementEntitlementManagementRequestBuilder) {
    return NewEntitlementmanagementEntitlementManagementRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get identityGovernance
// returns a IdentityGovernanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IdentityGovernanceRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityGovernanceRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityGovernanceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentityGovernanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityGovernanceable), nil
}
// LifecycleWorkflows provides operations to manage the lifecycleWorkflows property of the microsoft.graph.identityGovernance entity.
// returns a *LifecycleworkflowsLifecycleWorkflowsRequestBuilder when successful
func (m *IdentityGovernanceRequestBuilder) LifecycleWorkflows()(*LifecycleworkflowsLifecycleWorkflowsRequestBuilder) {
    return NewLifecycleworkflowsLifecycleWorkflowsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update identityGovernance
// returns a IdentityGovernanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IdentityGovernanceRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityGovernanceable, requestConfiguration *IdentityGovernanceRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityGovernanceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentityGovernanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityGovernanceable), nil
}
// PermissionsAnalytics provides operations to manage the permissionsAnalytics property of the microsoft.graph.identityGovernance entity.
// returns a *PermissionsanalyticsPermissionsAnalyticsRequestBuilder when successful
func (m *IdentityGovernanceRequestBuilder) PermissionsAnalytics()(*PermissionsanalyticsPermissionsAnalyticsRequestBuilder) {
    return NewPermissionsanalyticsPermissionsAnalyticsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PermissionsManagement provides operations to manage the permissionsManagement property of the microsoft.graph.identityGovernance entity.
// returns a *PermissionsmanagementPermissionsManagementRequestBuilder when successful
func (m *IdentityGovernanceRequestBuilder) PermissionsManagement()(*PermissionsmanagementPermissionsManagementRequestBuilder) {
    return NewPermissionsmanagementPermissionsManagementRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PrivilegedAccess provides operations to manage the privilegedAccess property of the microsoft.graph.identityGovernance entity.
// returns a *PrivilegedaccessPrivilegedAccessRequestBuilder when successful
func (m *IdentityGovernanceRequestBuilder) PrivilegedAccess()(*PrivilegedaccessPrivilegedAccessRequestBuilder) {
    return NewPrivilegedaccessPrivilegedAccessRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleManagementAlerts provides operations to manage the roleManagementAlerts property of the microsoft.graph.identityGovernance entity.
// returns a *RolemanagementalertsRoleManagementAlertsRequestBuilder when successful
func (m *IdentityGovernanceRequestBuilder) RoleManagementAlerts()(*RolemanagementalertsRoleManagementAlertsRequestBuilder) {
    return NewRolemanagementalertsRoleManagementAlertsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TermsOfUse provides operations to manage the termsOfUse property of the microsoft.graph.identityGovernance entity.
// returns a *TermsofuseTermsOfUseRequestBuilder when successful
func (m *IdentityGovernanceRequestBuilder) TermsOfUse()(*TermsofuseTermsOfUseRequestBuilder) {
    return NewTermsofuseTermsOfUseRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get identityGovernance
// returns a *RequestInformation when successful
func (m *IdentityGovernanceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update identityGovernance
// returns a *RequestInformation when successful
func (m *IdentityGovernanceRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityGovernanceable, requestConfiguration *IdentityGovernanceRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *IdentityGovernanceRequestBuilder when successful
func (m *IdentityGovernanceRequestBuilder) WithUrl(rawUrl string)(*IdentityGovernanceRequestBuilder) {
    return NewIdentityGovernanceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
