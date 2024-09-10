package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemDevicesItemRegisteredUsersGraphUserRequestBuilder casts the previous resource to user.
type ItemDevicesItemRegisteredUsersGraphUserRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemDevicesItemRegisteredUsersGraphUserRequestBuilderGetQueryParameters get the items of type microsoft.graph.user in the microsoft.graph.directoryObject collection
type ItemDevicesItemRegisteredUsersGraphUserRequestBuilderGetQueryParameters struct {
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
// ItemDevicesItemRegisteredUsersGraphUserRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDevicesItemRegisteredUsersGraphUserRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemDevicesItemRegisteredUsersGraphUserRequestBuilderGetQueryParameters
}
// NewItemDevicesItemRegisteredUsersGraphUserRequestBuilderInternal instantiates a new ItemDevicesItemRegisteredUsersGraphUserRequestBuilder and sets the default values.
func NewItemDevicesItemRegisteredUsersGraphUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDevicesItemRegisteredUsersGraphUserRequestBuilder) {
    m := &ItemDevicesItemRegisteredUsersGraphUserRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/devices/{device%2Did}/registeredUsers/graph.user{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemDevicesItemRegisteredUsersGraphUserRequestBuilder instantiates a new ItemDevicesItemRegisteredUsersGraphUserRequestBuilder and sets the default values.
func NewItemDevicesItemRegisteredUsersGraphUserRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDevicesItemRegisteredUsersGraphUserRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDevicesItemRegisteredUsersGraphUserRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemDevicesItemRegisteredUsersGraphUserCountRequestBuilder when successful
func (m *ItemDevicesItemRegisteredUsersGraphUserRequestBuilder) Count()(*ItemDevicesItemRegisteredUsersGraphUserCountRequestBuilder) {
    return NewItemDevicesItemRegisteredUsersGraphUserCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the items of type microsoft.graph.user in the microsoft.graph.directoryObject collection
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a UserCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemDevicesItemRegisteredUsersGraphUserRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemDevicesItemRegisteredUsersGraphUserRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserCollectionResponseable), nil
}
// ToGetRequestInformation get the items of type microsoft.graph.user in the microsoft.graph.directoryObject collection
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *RequestInformation when successful
func (m *ItemDevicesItemRegisteredUsersGraphUserRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemDevicesItemRegisteredUsersGraphUserRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *ItemDevicesItemRegisteredUsersGraphUserRequestBuilder when successful
func (m *ItemDevicesItemRegisteredUsersGraphUserRequestBuilder) WithUrl(rawUrl string)(*ItemDevicesItemRegisteredUsersGraphUserRequestBuilder) {
    return NewItemDevicesItemRegisteredUsersGraphUserRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
