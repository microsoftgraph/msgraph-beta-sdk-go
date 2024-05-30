package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilder provides operations to manage the directoryScopes property of the microsoft.graph.unifiedRoleAssignmentMultiple entity.
type DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilderGetQueryParameters read-only collection that references the directory objects that are scope of the assignment. Provided so that callers can get the directory objects using $expand at the same time as getting the role assignment. Read-only.  Supports $expand.
type DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilderGetQueryParameters struct {
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
// DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilderGetQueryParameters
}
// ByDirectoryObjectId provides operations to manage the directoryScopes property of the microsoft.graph.unifiedRoleAssignmentMultiple entity.
// returns a *DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryObjectItemRequestBuilder when successful
func (m *DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilder) ByDirectoryObjectId(directoryObjectId string)(*DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if directoryObjectId != "" {
        urlTplParams["directoryObject%2Did"] = directoryObjectId
    }
    return NewDevicemanagementRoleassignmentsItemDirectoryscopesDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDevicemanagementRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilderInternal instantiates a new DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilder and sets the default values.
func NewDevicemanagementRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilder) {
    m := &DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/deviceManagement/roleAssignments/{unifiedRoleAssignmentMultiple%2Did}/directoryScopes{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDevicemanagementRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilder instantiates a new DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilder and sets the default values.
func NewDevicemanagementRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicemanagementRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DevicemanagementRoleassignmentsItemDirectoryscopesCountRequestBuilder when successful
func (m *DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilder) Count()(*DevicemanagementRoleassignmentsItemDirectoryscopesCountRequestBuilder) {
    return NewDevicemanagementRoleassignmentsItemDirectoryscopesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read-only collection that references the directory objects that are scope of the assignment. Provided so that callers can get the directory objects using $expand at the same time as getting the role assignment. Read-only.  Supports $expand.
// returns a DirectoryObjectCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilder) Get(ctx context.Context, requestConfiguration *DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
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
func (m *DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilder when successful
func (m *DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilder) WithUrl(rawUrl string)(*DevicemanagementRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilder) {
    return NewDevicemanagementRoleassignmentsItemDirectoryscopesDirectoryScopesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
