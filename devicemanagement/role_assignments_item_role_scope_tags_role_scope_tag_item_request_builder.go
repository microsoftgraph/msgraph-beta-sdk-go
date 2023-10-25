package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RoleAssignmentsItemRoleScopeTagsRoleScopeTagItemRequestBuilder provides operations to manage the roleScopeTags property of the microsoft.graph.deviceAndAppManagementRoleAssignment entity.
type RoleAssignmentsItemRoleScopeTagsRoleScopeTagItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RoleAssignmentsItemRoleScopeTagsRoleScopeTagItemRequestBuilderGetQueryParameters the set of Role Scope Tags defined on the Role Assignment.
type RoleAssignmentsItemRoleScopeTagsRoleScopeTagItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RoleAssignmentsItemRoleScopeTagsRoleScopeTagItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleAssignmentsItemRoleScopeTagsRoleScopeTagItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RoleAssignmentsItemRoleScopeTagsRoleScopeTagItemRequestBuilderGetQueryParameters
}
// NewRoleAssignmentsItemRoleScopeTagsRoleScopeTagItemRequestBuilderInternal instantiates a new RoleScopeTagItemRequestBuilder and sets the default values.
func NewRoleAssignmentsItemRoleScopeTagsRoleScopeTagItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleAssignmentsItemRoleScopeTagsRoleScopeTagItemRequestBuilder) {
    m := &RoleAssignmentsItemRoleScopeTagsRoleScopeTagItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/roleAssignments/{deviceAndAppManagementRoleAssignment%2Did}/roleScopeTags/{roleScopeTag%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewRoleAssignmentsItemRoleScopeTagsRoleScopeTagItemRequestBuilder instantiates a new RoleScopeTagItemRequestBuilder and sets the default values.
func NewRoleAssignmentsItemRoleScopeTagsRoleScopeTagItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleAssignmentsItemRoleScopeTagsRoleScopeTagItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleAssignmentsItemRoleScopeTagsRoleScopeTagItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the set of Role Scope Tags defined on the Role Assignment.
func (m *RoleAssignmentsItemRoleScopeTagsRoleScopeTagItemRequestBuilder) Get(ctx context.Context, requestConfiguration *RoleAssignmentsItemRoleScopeTagsRoleScopeTagItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RoleScopeTagable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateRoleScopeTagFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RoleScopeTagable), nil
}
// ToGetRequestInformation the set of Role Scope Tags defined on the Role Assignment.
func (m *RoleAssignmentsItemRoleScopeTagsRoleScopeTagItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RoleAssignmentsItemRoleScopeTagsRoleScopeTagItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *RoleAssignmentsItemRoleScopeTagsRoleScopeTagItemRequestBuilder) WithUrl(rawUrl string)(*RoleAssignmentsItemRoleScopeTagsRoleScopeTagItemRequestBuilder) {
    return NewRoleAssignmentsItemRoleScopeTagsRoleScopeTagItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
