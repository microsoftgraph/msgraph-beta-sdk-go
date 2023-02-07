package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemDevicesItemRegisteredOwnersRequestBuilder provides operations to manage the registeredOwners property of the microsoft.graph.device entity.
type ItemDevicesItemRegisteredOwnersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemDevicesItemRegisteredOwnersRequestBuilderGetQueryParameters the user that cloud joined the device or registered their personal device. The registered owner is set at the time of registration. Currently, there can be only one owner. Read-only. Nullable. Supports $expand.
type ItemDevicesItemRegisteredOwnersRequestBuilderGetQueryParameters struct {
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
// ItemDevicesItemRegisteredOwnersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDevicesItemRegisteredOwnersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemDevicesItemRegisteredOwnersRequestBuilderGetQueryParameters
}
// NewItemDevicesItemRegisteredOwnersRequestBuilderInternal instantiates a new RegisteredOwnersRequestBuilder and sets the default values.
func NewItemDevicesItemRegisteredOwnersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDevicesItemRegisteredOwnersRequestBuilder) {
    m := &ItemDevicesItemRegisteredOwnersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/devices/{device%2Did}/registeredOwners{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemDevicesItemRegisteredOwnersRequestBuilder instantiates a new RegisteredOwnersRequestBuilder and sets the default values.
func NewItemDevicesItemRegisteredOwnersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDevicesItemRegisteredOwnersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDevicesItemRegisteredOwnersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ItemDevicesItemRegisteredOwnersRequestBuilder) Count()(*ItemDevicesItemRegisteredOwnersCountRequestBuilder) {
    return NewItemDevicesItemRegisteredOwnersCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get the user that cloud joined the device or registered their personal device. The registered owner is set at the time of registration. Currently, there can be only one owner. Read-only. Nullable. Supports $expand.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/device-list-registeredowners?view=graph-rest-1.0
func (m *ItemDevicesItemRegisteredOwnersRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemDevicesItemRegisteredOwnersRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable), nil
}
// MicrosoftGraphEndpoint casts the previous resource to endpoint.
func (m *ItemDevicesItemRegisteredOwnersRequestBuilder) MicrosoftGraphEndpoint()(*ItemDevicesItemRegisteredOwnersMicrosoftGraphEndpointRequestBuilder) {
    return NewItemDevicesItemRegisteredOwnersMicrosoftGraphEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphServicePrincipal casts the previous resource to servicePrincipal.
func (m *ItemDevicesItemRegisteredOwnersRequestBuilder) MicrosoftGraphServicePrincipal()(*ItemDevicesItemRegisteredOwnersMicrosoftGraphServicePrincipalRequestBuilder) {
    return NewItemDevicesItemRegisteredOwnersMicrosoftGraphServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUser casts the previous resource to user.
func (m *ItemDevicesItemRegisteredOwnersRequestBuilder) MicrosoftGraphUser()(*ItemDevicesItemRegisteredOwnersMicrosoftGraphUserRequestBuilder) {
    return NewItemDevicesItemRegisteredOwnersMicrosoftGraphUserRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Ref provides operations to manage the collection of user entities.
func (m *ItemDevicesItemRegisteredOwnersRequestBuilder) Ref()(*ItemDevicesItemRegisteredOwnersRefRequestBuilder) {
    return NewItemDevicesItemRegisteredOwnersRefRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToGetRequestInformation the user that cloud joined the device or registered their personal device. The registered owner is set at the time of registration. Currently, there can be only one owner. Read-only. Nullable. Supports $expand.
func (m *ItemDevicesItemRegisteredOwnersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemDevicesItemRegisteredOwnersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
