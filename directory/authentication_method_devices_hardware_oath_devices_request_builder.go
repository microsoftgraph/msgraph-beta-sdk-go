package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AuthenticationMethodDevicesHardwareOathDevicesRequestBuilder provides operations to manage the hardwareOathDevices property of the microsoft.graph.authenticationMethodDevice entity.
type AuthenticationMethodDevicesHardwareOathDevicesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthenticationMethodDevicesHardwareOathDevicesRequestBuilderGetQueryParameters get hardwareOathDevices from directory
type AuthenticationMethodDevicesHardwareOathDevicesRequestBuilderGetQueryParameters struct {
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
// AuthenticationMethodDevicesHardwareOathDevicesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationMethodDevicesHardwareOathDevicesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationMethodDevicesHardwareOathDevicesRequestBuilderGetQueryParameters
}
// AuthenticationMethodDevicesHardwareOathDevicesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationMethodDevicesHardwareOathDevicesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByHardwareOathTokenAuthenticationMethodDeviceId provides operations to manage the hardwareOathDevices property of the microsoft.graph.authenticationMethodDevice entity.
// returns a *AuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder when successful
func (m *AuthenticationMethodDevicesHardwareOathDevicesRequestBuilder) ByHardwareOathTokenAuthenticationMethodDeviceId(hardwareOathTokenAuthenticationMethodDeviceId string)(*AuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if hardwareOathTokenAuthenticationMethodDeviceId != "" {
        urlTplParams["hardwareOathTokenAuthenticationMethodDevice%2Did"] = hardwareOathTokenAuthenticationMethodDeviceId
    }
    return NewAuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAuthenticationMethodDevicesHardwareOathDevicesRequestBuilderInternal instantiates a new AuthenticationMethodDevicesHardwareOathDevicesRequestBuilder and sets the default values.
func NewAuthenticationMethodDevicesHardwareOathDevicesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationMethodDevicesHardwareOathDevicesRequestBuilder) {
    m := &AuthenticationMethodDevicesHardwareOathDevicesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/authenticationMethodDevices/hardwareOathDevices{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewAuthenticationMethodDevicesHardwareOathDevicesRequestBuilder instantiates a new AuthenticationMethodDevicesHardwareOathDevicesRequestBuilder and sets the default values.
func NewAuthenticationMethodDevicesHardwareOathDevicesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationMethodDevicesHardwareOathDevicesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationMethodDevicesHardwareOathDevicesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *AuthenticationMethodDevicesHardwareOathDevicesCountRequestBuilder when successful
func (m *AuthenticationMethodDevicesHardwareOathDevicesRequestBuilder) Count()(*AuthenticationMethodDevicesHardwareOathDevicesCountRequestBuilder) {
    return NewAuthenticationMethodDevicesHardwareOathDevicesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get hardwareOathDevices from directory
// returns a HardwareOathTokenAuthenticationMethodDeviceCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthenticationMethodDevicesHardwareOathDevicesRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationMethodDevicesHardwareOathDevicesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareOathTokenAuthenticationMethodDeviceCollectionResponseable, error) {
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
// Post create new navigation property to hardwareOathDevices for directory
// returns a HardwareOathTokenAuthenticationMethodDeviceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthenticationMethodDevicesHardwareOathDevicesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareOathTokenAuthenticationMethodDeviceable, requestConfiguration *AuthenticationMethodDevicesHardwareOathDevicesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareOathTokenAuthenticationMethodDeviceable, error) {
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
// ToGetRequestInformation get hardwareOathDevices from directory
// returns a *RequestInformation when successful
func (m *AuthenticationMethodDevicesHardwareOathDevicesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationMethodDevicesHardwareOathDevicesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to hardwareOathDevices for directory
// returns a *RequestInformation when successful
func (m *AuthenticationMethodDevicesHardwareOathDevicesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareOathTokenAuthenticationMethodDeviceable, requestConfiguration *AuthenticationMethodDevicesHardwareOathDevicesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AuthenticationMethodDevicesHardwareOathDevicesRequestBuilder when successful
func (m *AuthenticationMethodDevicesHardwareOathDevicesRequestBuilder) WithUrl(rawUrl string)(*AuthenticationMethodDevicesHardwareOathDevicesRequestBuilder) {
    return NewAuthenticationMethodDevicesHardwareOathDevicesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
