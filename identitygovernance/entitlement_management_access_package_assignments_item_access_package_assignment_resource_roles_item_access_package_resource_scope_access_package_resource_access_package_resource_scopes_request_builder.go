package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopesRequestBuilder provides operations to manage the accessPackageResourceScopes property of the microsoft.graph.accessPackageResource entity.
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopesRequestBuilderGetQueryParameters read-only. Nullable. Supports $expand.
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopesRequestBuilderGetQueryParameters
}
// EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAccessPackageResourceScopeId provides operations to manage the accessPackageResourceScopes property of the microsoft.graph.accessPackageResource entity.
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopesRequestBuilder) ByAccessPackageResourceScopeId(accessPackageResourceScopeId string)(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if accessPackageResourceScopeId != "" {
        urlTplParams["accessPackageResourceScope%2Did"] = accessPackageResourceScopeId
    }
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopesRequestBuilderInternal instantiates a new AccessPackageResourceScopesRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopesRequestBuilder) {
    m := &EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/{accessPackageAssignment%2Did}/accessPackageAssignmentResourceRoles/{accessPackageAssignmentResourceRole%2Did}/accessPackageResourceScope/accessPackageResource/accessPackageResourceScopes{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopesRequestBuilder instantiates a new AccessPackageResourceScopesRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopesRequestBuilder) Count()(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopesCountRequestBuilder) {
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read-only. Nullable. Supports $expand.
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopesRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceScopeCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceScopeCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceScopeCollectionResponseable), nil
}
// Post create new navigation property to accessPackageResourceScopes for identityGovernance
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceScopeable, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceScopeable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceScopeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceScopeable), nil
}
// ToGetRequestInformation read-only. Nullable. Supports $expand.
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to accessPackageResourceScopes for identityGovernance
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceScopeable, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopesRequestBuilder) WithUrl(rawUrl string)(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopesRequestBuilder) {
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
