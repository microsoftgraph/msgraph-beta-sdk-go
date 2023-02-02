package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilder provides operations to manage the accessPackageResourceScopes property of the microsoft.graph.accessPackageResource entity.
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderGetQueryParameters read-only. Nullable. Supports $expand.
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderGetQueryParameters
}
// EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderInternal instantiates a new AccessPackageResourceScopeItemRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, accessPackageResourceScopeId *string)(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilder) {
    m := &EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/{accessPackageAssignment%2Did}/accessPackageAssignmentResourceRoles/{accessPackageAssignmentResourceRole%2Did}/accessPackageResourceScope/accessPackageResource/accessPackageResourceRoles/{accessPackageResourceRole%2Did}/accessPackageResource/accessPackageResourceScopes/{accessPackageResourceScope%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if accessPackageResourceScopeId != nil {
        urlTplParams["accessPackageResourceScope%2Did"] = *accessPackageResourceScopeId
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilder instantiates a new AccessPackageResourceScopeItemRequestBuilder and sets the default values.
func NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Delete delete navigation property accessPackageResourceScopes for identityGovernance
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get read-only. Nullable. Supports $expand.
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceScopeable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceScopeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceScopeable), nil
}
// Patch update the navigation property accessPackageResourceScopes in identityGovernance
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceScopeable, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceScopeable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccessPackageResourceScopeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceScopeable), nil
}
// ToDeleteRequestInformation delete navigation property accessPackageResourceScopes for identityGovernance
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation read-only. Nullable. Supports $expand.
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property accessPackageResourceScopes in identityGovernance
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AccessPackageResourceScopeable, requestConfiguration *EntitlementManagementAccessPackageAssignmentsItemAccessPackageAssignmentResourceRolesItemAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRolesItemAccessPackageResourceAccessPackageResourceScopesAccessPackageResourceScopeItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
