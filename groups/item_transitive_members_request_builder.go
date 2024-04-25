package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTransitiveMembersRequestBuilder provides operations to manage the transitiveMembers property of the microsoft.graph.group entity.
type ItemTransitiveMembersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTransitiveMembersRequestBuilderGetQueryParameters the direct and transitive members of a group. Nullable.
type ItemTransitiveMembersRequestBuilderGetQueryParameters struct {
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
// ItemTransitiveMembersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTransitiveMembersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTransitiveMembersRequestBuilderGetQueryParameters
}
// ByDirectoryObjectId provides operations to manage the transitiveMembers property of the microsoft.graph.group entity.
// returns a *ItemTransitiveMembersDirectoryObjectItemRequestBuilder when successful
func (m *ItemTransitiveMembersRequestBuilder) ByDirectoryObjectId(directoryObjectId string)(*ItemTransitiveMembersDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if directoryObjectId != "" {
        urlTplParams["directoryObject%2Did"] = directoryObjectId
    }
    return NewItemTransitiveMembersDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemTransitiveMembersRequestBuilderInternal instantiates a new ItemTransitiveMembersRequestBuilder and sets the default values.
func NewItemTransitiveMembersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTransitiveMembersRequestBuilder) {
    m := &ItemTransitiveMembersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/transitiveMembers{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemTransitiveMembersRequestBuilder instantiates a new ItemTransitiveMembersRequestBuilder and sets the default values.
func NewItemTransitiveMembersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTransitiveMembersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTransitiveMembersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemTransitiveMembersCountRequestBuilder when successful
func (m *ItemTransitiveMembersRequestBuilder) Count()(*ItemTransitiveMembersCountRequestBuilder) {
    return NewItemTransitiveMembersCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the direct and transitive members of a group. Nullable.
// returns a DirectoryObjectCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTransitiveMembersRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTransitiveMembersRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
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
// GraphApplication casts the previous resource to application.
// returns a *ItemTransitiveMembersGraphApplicationRequestBuilder when successful
func (m *ItemTransitiveMembersRequestBuilder) GraphApplication()(*ItemTransitiveMembersGraphApplicationRequestBuilder) {
    return NewItemTransitiveMembersGraphApplicationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphDevice casts the previous resource to device.
// returns a *ItemTransitiveMembersGraphDeviceRequestBuilder when successful
func (m *ItemTransitiveMembersRequestBuilder) GraphDevice()(*ItemTransitiveMembersGraphDeviceRequestBuilder) {
    return NewItemTransitiveMembersGraphDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphGroup casts the previous resource to group.
// returns a *ItemTransitiveMembersGraphGroupRequestBuilder when successful
func (m *ItemTransitiveMembersRequestBuilder) GraphGroup()(*ItemTransitiveMembersGraphGroupRequestBuilder) {
    return NewItemTransitiveMembersGraphGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphOrgContact casts the previous resource to orgContact.
// returns a *ItemTransitiveMembersGraphOrgContactRequestBuilder when successful
func (m *ItemTransitiveMembersRequestBuilder) GraphOrgContact()(*ItemTransitiveMembersGraphOrgContactRequestBuilder) {
    return NewItemTransitiveMembersGraphOrgContactRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphServicePrincipal casts the previous resource to servicePrincipal.
// returns a *ItemTransitiveMembersGraphServicePrincipalRequestBuilder when successful
func (m *ItemTransitiveMembersRequestBuilder) GraphServicePrincipal()(*ItemTransitiveMembersGraphServicePrincipalRequestBuilder) {
    return NewItemTransitiveMembersGraphServicePrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphUser casts the previous resource to user.
// returns a *ItemTransitiveMembersGraphUserRequestBuilder when successful
func (m *ItemTransitiveMembersRequestBuilder) GraphUser()(*ItemTransitiveMembersGraphUserRequestBuilder) {
    return NewItemTransitiveMembersGraphUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the direct and transitive members of a group. Nullable.
// returns a *RequestInformation when successful
func (m *ItemTransitiveMembersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTransitiveMembersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTransitiveMembersRequestBuilder when successful
func (m *ItemTransitiveMembersRequestBuilder) WithUrl(rawUrl string)(*ItemTransitiveMembersRequestBuilder) {
    return NewItemTransitiveMembersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
