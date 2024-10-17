package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTransitiveMemberOfGraphDirectoryRoleRequestBuilder casts the previous resource to directoryRole.
type ItemTransitiveMemberOfGraphDirectoryRoleRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTransitiveMemberOfGraphDirectoryRoleRequestBuilderGetQueryParameters get the items of type microsoft.graph.directoryRole in the microsoft.graph.directoryObject collection
type ItemTransitiveMemberOfGraphDirectoryRoleRequestBuilderGetQueryParameters struct {
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
// ItemTransitiveMemberOfGraphDirectoryRoleRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTransitiveMemberOfGraphDirectoryRoleRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTransitiveMemberOfGraphDirectoryRoleRequestBuilderGetQueryParameters
}
// NewItemTransitiveMemberOfGraphDirectoryRoleRequestBuilderInternal instantiates a new ItemTransitiveMemberOfGraphDirectoryRoleRequestBuilder and sets the default values.
func NewItemTransitiveMemberOfGraphDirectoryRoleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTransitiveMemberOfGraphDirectoryRoleRequestBuilder) {
    m := &ItemTransitiveMemberOfGraphDirectoryRoleRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/transitiveMemberOf/graph.directoryRole{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemTransitiveMemberOfGraphDirectoryRoleRequestBuilder instantiates a new ItemTransitiveMemberOfGraphDirectoryRoleRequestBuilder and sets the default values.
func NewItemTransitiveMemberOfGraphDirectoryRoleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTransitiveMemberOfGraphDirectoryRoleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTransitiveMemberOfGraphDirectoryRoleRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemTransitiveMemberOfGraphDirectoryRoleCountRequestBuilder when successful
func (m *ItemTransitiveMemberOfGraphDirectoryRoleRequestBuilder) Count()(*ItemTransitiveMemberOfGraphDirectoryRoleCountRequestBuilder) {
    return NewItemTransitiveMemberOfGraphDirectoryRoleCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the items of type microsoft.graph.directoryRole in the microsoft.graph.directoryObject collection
// returns a DirectoryRoleCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTransitiveMemberOfGraphDirectoryRoleRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTransitiveMemberOfGraphDirectoryRoleRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryRoleCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryRoleCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryRoleCollectionResponseable), nil
}
// ToGetRequestInformation get the items of type microsoft.graph.directoryRole in the microsoft.graph.directoryObject collection
// returns a *RequestInformation when successful
func (m *ItemTransitiveMemberOfGraphDirectoryRoleRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTransitiveMemberOfGraphDirectoryRoleRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTransitiveMemberOfGraphDirectoryRoleRequestBuilder when successful
func (m *ItemTransitiveMemberOfGraphDirectoryRoleRequestBuilder) WithUrl(rawUrl string)(*ItemTransitiveMemberOfGraphDirectoryRoleRequestBuilder) {
    return NewItemTransitiveMemberOfGraphDirectoryRoleRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
