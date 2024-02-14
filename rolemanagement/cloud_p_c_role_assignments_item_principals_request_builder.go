package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CloudPCRoleAssignmentsItemPrincipalsRequestBuilder provides operations to manage the principals property of the microsoft.graph.unifiedRoleAssignmentMultiple entity.
type CloudPCRoleAssignmentsItemPrincipalsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CloudPCRoleAssignmentsItemPrincipalsRequestBuilderGetQueryParameters read-only collection that references the assigned principals. Provided so that callers can get the principals using $expand at the same time as getting the role assignment. Read-only.  Supports $expand.
type CloudPCRoleAssignmentsItemPrincipalsRequestBuilderGetQueryParameters struct {
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
// CloudPCRoleAssignmentsItemPrincipalsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CloudPCRoleAssignmentsItemPrincipalsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CloudPCRoleAssignmentsItemPrincipalsRequestBuilderGetQueryParameters
}
// ByDirectoryObjectId provides operations to manage the principals property of the microsoft.graph.unifiedRoleAssignmentMultiple entity.
// returns a *CloudPCRoleAssignmentsItemPrincipalsDirectoryObjectItemRequestBuilder when successful
func (m *CloudPCRoleAssignmentsItemPrincipalsRequestBuilder) ByDirectoryObjectId(directoryObjectId string)(*CloudPCRoleAssignmentsItemPrincipalsDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if directoryObjectId != "" {
        urlTplParams["directoryObject%2Did"] = directoryObjectId
    }
    return NewCloudPCRoleAssignmentsItemPrincipalsDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCloudPCRoleAssignmentsItemPrincipalsRequestBuilderInternal instantiates a new CloudPCRoleAssignmentsItemPrincipalsRequestBuilder and sets the default values.
func NewCloudPCRoleAssignmentsItemPrincipalsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudPCRoleAssignmentsItemPrincipalsRequestBuilder) {
    m := &CloudPCRoleAssignmentsItemPrincipalsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/cloudPC/roleAssignments/{unifiedRoleAssignmentMultiple%2Did}/principals{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCloudPCRoleAssignmentsItemPrincipalsRequestBuilder instantiates a new CloudPCRoleAssignmentsItemPrincipalsRequestBuilder and sets the default values.
func NewCloudPCRoleAssignmentsItemPrincipalsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudPCRoleAssignmentsItemPrincipalsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCloudPCRoleAssignmentsItemPrincipalsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CloudPCRoleAssignmentsItemPrincipalsCountRequestBuilder when successful
func (m *CloudPCRoleAssignmentsItemPrincipalsRequestBuilder) Count()(*CloudPCRoleAssignmentsItemPrincipalsCountRequestBuilder) {
    return NewCloudPCRoleAssignmentsItemPrincipalsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read-only collection that references the assigned principals. Provided so that callers can get the principals using $expand at the same time as getting the role assignment. Read-only.  Supports $expand.
// returns a DirectoryObjectCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CloudPCRoleAssignmentsItemPrincipalsRequestBuilder) Get(ctx context.Context, requestConfiguration *CloudPCRoleAssignmentsItemPrincipalsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
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
// ToGetRequestInformation read-only collection that references the assigned principals. Provided so that callers can get the principals using $expand at the same time as getting the role assignment. Read-only.  Supports $expand.
// returns a *RequestInformation when successful
func (m *CloudPCRoleAssignmentsItemPrincipalsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CloudPCRoleAssignmentsItemPrincipalsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CloudPCRoleAssignmentsItemPrincipalsRequestBuilder when successful
func (m *CloudPCRoleAssignmentsItemPrincipalsRequestBuilder) WithUrl(rawUrl string)(*CloudPCRoleAssignmentsItemPrincipalsRequestBuilder) {
    return NewCloudPCRoleAssignmentsItemPrincipalsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
