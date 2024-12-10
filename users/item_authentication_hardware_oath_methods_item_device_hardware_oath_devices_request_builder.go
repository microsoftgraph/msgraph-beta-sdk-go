package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesRequestBuilder provides operations to manage the hardwareOathDevices property of the microsoft.graph.authenticationMethodDevice entity.
type ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesRequestBuilderGetQueryParameters exposes the hardware OATH method in the directory.
type ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesRequestBuilderGetQueryParameters struct {
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
// ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesRequestBuilderGetQueryParameters
}
// ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByHardwareOathTokenAuthenticationMethodDeviceId provides operations to manage the hardwareOathDevices property of the microsoft.graph.authenticationMethodDevice entity.
// returns a *ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder when successful
func (m *ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesRequestBuilder) ByHardwareOathTokenAuthenticationMethodDeviceId(hardwareOathTokenAuthenticationMethodDeviceId string)(*ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if hardwareOathTokenAuthenticationMethodDeviceId != "" {
        urlTplParams["hardwareOathTokenAuthenticationMethodDevice%2Did"] = hardwareOathTokenAuthenticationMethodDeviceId
    }
    return NewItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesRequestBuilderInternal instantiates a new ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesRequestBuilder and sets the default values.
func NewItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesRequestBuilder) {
    m := &ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/authentication/hardwareOathMethods/{hardwareOathAuthenticationMethod%2Did}/device/hardwareOathDevices{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesRequestBuilder instantiates a new ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesRequestBuilder and sets the default values.
func NewItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesCountRequestBuilder when successful
func (m *ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesRequestBuilder) Count()(*ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesCountRequestBuilder) {
    return NewItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get exposes the hardware OATH method in the directory.
// returns a HardwareOathTokenAuthenticationMethodDeviceCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareOathTokenAuthenticationMethodDeviceCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateHardwareOathTokenAuthenticationMethodDeviceCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareOathTokenAuthenticationMethodDeviceCollectionResponseable), nil
}
// Post create new navigation property to hardwareOathDevices for users
// returns a HardwareOathTokenAuthenticationMethodDeviceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareOathTokenAuthenticationMethodDeviceable, requestConfiguration *ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareOathTokenAuthenticationMethodDeviceable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateHardwareOathTokenAuthenticationMethodDeviceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareOathTokenAuthenticationMethodDeviceable), nil
}
// ToGetRequestInformation exposes the hardware OATH method in the directory.
// returns a *RequestInformation when successful
func (m *ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to hardwareOathDevices for users
// returns a *RequestInformation when successful
func (m *ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareOathTokenAuthenticationMethodDeviceable, requestConfiguration *ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesRequestBuilder when successful
func (m *ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesRequestBuilder) WithUrl(rawUrl string)(*ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesRequestBuilder) {
    return NewItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
