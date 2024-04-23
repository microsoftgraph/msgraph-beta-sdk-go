package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DirectoryRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdRequestBuilder provides operations to call the assignedPrincipals method.
type DirectoryRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DirectoryRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdRequestBuilderGetQueryParameters get the list of security principals (users, groups, and service principals) that are assigned to a specific role for different scopes either directly or transitively. You can use the $count query parameter to also get the count. To list the direct and transitive role assignments for a specific principal, use the List transitiveRoleAssignments API.
type DirectoryRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Usage: directoryScopeId='@directoryScopeId'
    DirectoryScopeId *string `uriparametername:"directoryScopeId"`
    // Usage: directoryScopeType='@directoryScopeType'
    DirectoryScopeType *string `uriparametername:"directoryScopeType"`
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
    // Usage: transitive=@transitive
    Transitive *bool `uriparametername:"transitive"`
}
// DirectoryRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdRequestBuilderGetQueryParameters
}
// NewDirectoryRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdRequestBuilderInternal instantiates a new DirectoryRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdRequestBuilder and sets the default values.
func NewDirectoryRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdRequestBuilder) {
    m := &DirectoryRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/directory/roleDefinitions/{unifiedRoleDefinition%2Did}/assignedPrincipals(transitive=@transitive,directoryScopeType='@directoryScopeType',directoryScopeId='@directoryScopeId'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top,directoryScopeId*,directoryScopeType*,transitive*}", pathParameters),
    }
    return m
}
// NewDirectoryRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdRequestBuilder instantiates a new DirectoryRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdRequestBuilder and sets the default values.
func NewDirectoryRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the list of security principals (users, groups, and service principals) that are assigned to a specific role for different scopes either directly or transitively. You can use the $count query parameter to also get the count. To list the direct and transitive role assignments for a specific principal, use the List transitiveRoleAssignments API.
// Deprecated: This method is obsolete. Use GetAsAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdGetResponse instead.
// returns a DirectoryRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/unifiedroledefinition-assignedprincipals?view=graph-rest-beta
func (m *DirectoryRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectoryRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdRequestBuilderGetRequestConfiguration)(DirectoryRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDirectoryRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DirectoryRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdResponseable), nil
}
// GetAsAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdGetResponse get the list of security principals (users, groups, and service principals) that are assigned to a specific role for different scopes either directly or transitively. You can use the $count query parameter to also get the count. To list the direct and transitive role assignments for a specific principal, use the List transitiveRoleAssignments API.
// returns a DirectoryRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/unifiedroledefinition-assignedprincipals?view=graph-rest-beta
func (m *DirectoryRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdRequestBuilder) GetAsAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdGetResponse(ctx context.Context, requestConfiguration *DirectoryRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdRequestBuilderGetRequestConfiguration)(DirectoryRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDirectoryRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DirectoryRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdGetResponseable), nil
}
// ToGetRequestInformation get the list of security principals (users, groups, and service principals) that are assigned to a specific role for different scopes either directly or transitively. You can use the $count query parameter to also get the count. To list the direct and transitive role assignments for a specific principal, use the List transitiveRoleAssignments API.
// returns a *RequestInformation when successful
func (m *DirectoryRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DirectoryRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdRequestBuilder when successful
func (m *DirectoryRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdRequestBuilder) WithUrl(rawUrl string)(*DirectoryRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdRequestBuilder) {
    return NewDirectoryRoleDefinitionsItemAssignedPrincipalsWithTransitivedirectoryScopeTypeDirectoryScopeTypeDirectoryScopeIdDirectoryScopeIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
