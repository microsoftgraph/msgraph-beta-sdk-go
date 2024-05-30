package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CloudpcRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilder provides operations to manage the directoryScopes property of the microsoft.graph.unifiedRoleAssignmentMultiple entity.
type CloudpcRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CloudpcRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilderGetQueryParameters read-only collection that references the directory objects that are scope of the assignment. Provided so that callers can get the directory objects using $expand at the same time as getting the role assignment. Read-only.  Supports $expand.
type CloudpcRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilderGetQueryParameters struct {
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
// CloudpcRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CloudpcRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CloudpcRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilderGetQueryParameters
}
// ByDirectoryObjectId provides operations to manage the directoryScopes property of the microsoft.graph.unifiedRoleAssignmentMultiple entity.
// returns a *CloudpcRoleassignmentsItemDirectoryscopesDirectoryObjectItemRequestBuilder when successful
func (m *CloudpcRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilder) ByDirectoryObjectId(directoryObjectId string)(*CloudpcRoleassignmentsItemDirectoryscopesDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if directoryObjectId != "" {
        urlTplParams["directoryObject%2Did"] = directoryObjectId
    }
    return NewCloudpcRoleassignmentsItemDirectoryscopesDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCloudpcRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilderInternal instantiates a new CloudpcRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilder and sets the default values.
func NewCloudpcRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudpcRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilder) {
    m := &CloudpcRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/cloudPC/roleAssignments/{unifiedRoleAssignmentMultiple%2Did}/directoryScopes{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCloudpcRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilder instantiates a new CloudpcRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilder and sets the default values.
func NewCloudpcRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudpcRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCloudpcRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CloudpcRoleassignmentsItemDirectoryscopesCountRequestBuilder when successful
func (m *CloudpcRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilder) Count()(*CloudpcRoleassignmentsItemDirectoryscopesCountRequestBuilder) {
    return NewCloudpcRoleassignmentsItemDirectoryscopesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read-only collection that references the directory objects that are scope of the assignment. Provided so that callers can get the directory objects using $expand at the same time as getting the role assignment. Read-only.  Supports $expand.
// returns a DirectoryObjectCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CloudpcRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilder) Get(ctx context.Context, requestConfiguration *CloudpcRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable), nil
}
// ToGetRequestInformation read-only collection that references the directory objects that are scope of the assignment. Provided so that callers can get the directory objects using $expand at the same time as getting the role assignment. Read-only.  Supports $expand.
// returns a *RequestInformation when successful
func (m *CloudpcRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CloudpcRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CloudpcRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilder when successful
func (m *CloudpcRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilder) WithUrl(rawUrl string)(*CloudpcRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilder) {
    return NewCloudpcRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
