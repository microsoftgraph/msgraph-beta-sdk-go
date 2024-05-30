package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RoleassignmentsItemRolescopetagsRoleScopeTagsRequestBuilder provides operations to manage the roleScopeTags property of the microsoft.graph.deviceAndAppManagementRoleAssignment entity.
type RoleassignmentsItemRolescopetagsRoleScopeTagsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RoleassignmentsItemRolescopetagsRoleScopeTagsRequestBuilderGetQueryParameters the set of Role Scope Tags defined on the Role Assignment.
type RoleassignmentsItemRolescopetagsRoleScopeTagsRequestBuilderGetQueryParameters struct {
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
// RoleassignmentsItemRolescopetagsRoleScopeTagsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleassignmentsItemRolescopetagsRoleScopeTagsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RoleassignmentsItemRolescopetagsRoleScopeTagsRequestBuilderGetQueryParameters
}
// ByRoleScopeTagId provides operations to manage the roleScopeTags property of the microsoft.graph.deviceAndAppManagementRoleAssignment entity.
// returns a *RoleassignmentsItemRolescopetagsRoleScopeTagItemRequestBuilder when successful
func (m *RoleassignmentsItemRolescopetagsRoleScopeTagsRequestBuilder) ByRoleScopeTagId(roleScopeTagId string)(*RoleassignmentsItemRolescopetagsRoleScopeTagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if roleScopeTagId != "" {
        urlTplParams["roleScopeTag%2Did"] = roleScopeTagId
    }
    return NewRoleassignmentsItemRolescopetagsRoleScopeTagItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewRoleassignmentsItemRolescopetagsRoleScopeTagsRequestBuilderInternal instantiates a new RoleassignmentsItemRolescopetagsRoleScopeTagsRequestBuilder and sets the default values.
func NewRoleassignmentsItemRolescopetagsRoleScopeTagsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleassignmentsItemRolescopetagsRoleScopeTagsRequestBuilder) {
    m := &RoleassignmentsItemRolescopetagsRoleScopeTagsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/roleAssignments/{deviceAndAppManagementRoleAssignment%2Did}/roleScopeTags{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewRoleassignmentsItemRolescopetagsRoleScopeTagsRequestBuilder instantiates a new RoleassignmentsItemRolescopetagsRoleScopeTagsRequestBuilder and sets the default values.
func NewRoleassignmentsItemRolescopetagsRoleScopeTagsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleassignmentsItemRolescopetagsRoleScopeTagsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleassignmentsItemRolescopetagsRoleScopeTagsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *RoleassignmentsItemRolescopetagsCountRequestBuilder when successful
func (m *RoleassignmentsItemRolescopetagsRoleScopeTagsRequestBuilder) Count()(*RoleassignmentsItemRolescopetagsCountRequestBuilder) {
    return NewRoleassignmentsItemRolescopetagsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the set of Role Scope Tags defined on the Role Assignment.
// returns a RoleScopeTagCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *RoleassignmentsItemRolescopetagsRoleScopeTagsRequestBuilder) Get(ctx context.Context, requestConfiguration *RoleassignmentsItemRolescopetagsRoleScopeTagsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RoleScopeTagCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateRoleScopeTagCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RoleScopeTagCollectionResponseable), nil
}
// ToGetRequestInformation the set of Role Scope Tags defined on the Role Assignment.
// returns a *RequestInformation when successful
func (m *RoleassignmentsItemRolescopetagsRoleScopeTagsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RoleassignmentsItemRolescopetagsRoleScopeTagsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *RoleassignmentsItemRolescopetagsRoleScopeTagsRequestBuilder when successful
func (m *RoleassignmentsItemRolescopetagsRoleScopeTagsRequestBuilder) WithUrl(rawUrl string)(*RoleassignmentsItemRolescopetagsRoleScopeTagsRequestBuilder) {
    return NewRoleassignmentsItemRolescopetagsRoleScopeTagsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
