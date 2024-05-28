package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DirectoryTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder provides operations to manage the transitiveRoleAssignments property of the microsoft.graph.rbacApplication entity.
type DirectoryTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DirectoryTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilderGetQueryParameters get the list of direct and transitive unifiedRoleAssignment objects for a specific principal. For example, if a user is assigned a Microsoft Entra role through group membership, the role assignment is transitive, and this request will list the group's ID as the principalId. Results can also be filtered by the roleDefinitionId and directoryScopeId. Supported only for directory (Microsoft Entra ID) provider. For more information, see Use Microsoft Entra groups to manage role assignments.
type DirectoryTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilderGetQueryParameters struct {
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
// DirectoryTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilderGetQueryParameters
}
// DirectoryTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUnifiedRoleAssignmentId provides operations to manage the transitiveRoleAssignments property of the microsoft.graph.rbacApplication entity.
// returns a *DirectoryTransitiveroleassignmentsUnifiedRoleAssignmentItemRequestBuilder when successful
func (m *DirectoryTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder) ByUnifiedRoleAssignmentId(unifiedRoleAssignmentId string)(*DirectoryTransitiveroleassignmentsUnifiedRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if unifiedRoleAssignmentId != "" {
        urlTplParams["unifiedRoleAssignment%2Did"] = unifiedRoleAssignmentId
    }
    return NewDirectoryTransitiveroleassignmentsUnifiedRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDirectoryTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilderInternal instantiates a new DirectoryTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder and sets the default values.
func NewDirectoryTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder) {
    m := &DirectoryTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/directory/transitiveRoleAssignments{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDirectoryTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder instantiates a new DirectoryTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder and sets the default values.
func NewDirectoryTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DirectoryTransitiveroleassignmentsCountRequestBuilder when successful
func (m *DirectoryTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder) Count()(*DirectoryTransitiveroleassignmentsCountRequestBuilder) {
    return NewDirectoryTransitiveroleassignmentsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the list of direct and transitive unifiedRoleAssignment objects for a specific principal. For example, if a user is assigned a Microsoft Entra role through group membership, the role assignment is transitive, and this request will list the group's ID as the principalId. Results can also be filtered by the roleDefinitionId and directoryScopeId. Supported only for directory (Microsoft Entra ID) provider. For more information, see Use Microsoft Entra groups to manage role assignments.
// returns a UnifiedRoleAssignmentCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/rbacapplication-list-transitiveroleassignments?view=graph-rest-beta
func (m *DirectoryTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectoryTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleAssignmentCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentCollectionResponseable), nil
}
// Post create new navigation property to transitiveRoleAssignments for roleManagement
// returns a UnifiedRoleAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentable, requestConfiguration *DirectoryTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRoleAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentable), nil
}
// ToGetRequestInformation get the list of direct and transitive unifiedRoleAssignment objects for a specific principal. For example, if a user is assigned a Microsoft Entra role through group membership, the role assignment is transitive, and this request will list the group's ID as the principalId. Results can also be filtered by the roleDefinitionId and directoryScopeId. Supported only for directory (Microsoft Entra ID) provider. For more information, see Use Microsoft Entra groups to manage role assignments.
// returns a *RequestInformation when successful
func (m *DirectoryTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DirectoryTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to transitiveRoleAssignments for roleManagement
// returns a *RequestInformation when successful
func (m *DirectoryTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRoleAssignmentable, requestConfiguration *DirectoryTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DirectoryTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder when successful
func (m *DirectoryTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder) WithUrl(rawUrl string)(*DirectoryTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder) {
    return NewDirectoryTransitiveroleassignmentsTransitiveRoleAssignmentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
