package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemDevicesItemDeviceTemplateRequestBuilder provides operations to manage the deviceTemplate property of the microsoft.graph.device entity.
type ItemDevicesItemDeviceTemplateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemDevicesItemDeviceTemplateRequestBuilderGetQueryParameters device template used to instantiate this device. Nullable. Read-only.
type ItemDevicesItemDeviceTemplateRequestBuilderGetQueryParameters struct {
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
// ItemDevicesItemDeviceTemplateRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDevicesItemDeviceTemplateRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemDevicesItemDeviceTemplateRequestBuilderGetQueryParameters
}
// ByDeviceTemplateId provides operations to manage the deviceTemplate property of the microsoft.graph.device entity.
// returns a *ItemDevicesItemDeviceTemplateDeviceTemplateItemRequestBuilder when successful
func (m *ItemDevicesItemDeviceTemplateRequestBuilder) ByDeviceTemplateId(deviceTemplateId string)(*ItemDevicesItemDeviceTemplateDeviceTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if deviceTemplateId != "" {
        urlTplParams["deviceTemplate%2Did"] = deviceTemplateId
    }
    return NewItemDevicesItemDeviceTemplateDeviceTemplateItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemDevicesItemDeviceTemplateRequestBuilderInternal instantiates a new ItemDevicesItemDeviceTemplateRequestBuilder and sets the default values.
func NewItemDevicesItemDeviceTemplateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDevicesItemDeviceTemplateRequestBuilder) {
    m := &ItemDevicesItemDeviceTemplateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/devices/{device%2Did}/deviceTemplate{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemDevicesItemDeviceTemplateRequestBuilder instantiates a new ItemDevicesItemDeviceTemplateRequestBuilder and sets the default values.
func NewItemDevicesItemDeviceTemplateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDevicesItemDeviceTemplateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDevicesItemDeviceTemplateRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemDevicesItemDeviceTemplateCountRequestBuilder when successful
func (m *ItemDevicesItemDeviceTemplateRequestBuilder) Count()(*ItemDevicesItemDeviceTemplateCountRequestBuilder) {
    return NewItemDevicesItemDeviceTemplateCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get device template used to instantiate this device. Nullable. Read-only.
// returns a DeviceTemplateCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemDevicesItemDeviceTemplateRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemDevicesItemDeviceTemplateRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceTemplateCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceTemplateCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceTemplateCollectionResponseable), nil
}
// ToGetRequestInformation device template used to instantiate this device. Nullable. Read-only.
// returns a *RequestInformation when successful
func (m *ItemDevicesItemDeviceTemplateRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemDevicesItemDeviceTemplateRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemDevicesItemDeviceTemplateRequestBuilder when successful
func (m *ItemDevicesItemDeviceTemplateRequestBuilder) WithUrl(rawUrl string)(*ItemDevicesItemDeviceTemplateRequestBuilder) {
    return NewItemDevicesItemDeviceTemplateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
