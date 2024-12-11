package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder provides operations to manage the hardwareOathDevices property of the microsoft.graph.authenticationMethodDevice entity.
type ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilderGetQueryParameters exposes the hardware OATH method in the directory.
type ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilderGetQueryParameters
}
// ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AssignTo provides operations to manage the assignTo property of the microsoft.graph.hardwareOathTokenAuthenticationMethodDevice entity.
// returns a *ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesItemAssignToRequestBuilder when successful
func (m *ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder) AssignTo()(*ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesItemAssignToRequestBuilder) {
    return NewItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesItemAssignToRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilderInternal instantiates a new ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder and sets the default values.
func NewItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder) {
    m := &ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/authentication/hardwareOathMethods/{hardwareOathAuthenticationMethod%2Did}/device/hardwareOathDevices/{hardwareOathTokenAuthenticationMethodDevice%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder instantiates a new ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder and sets the default values.
func NewItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property hardwareOathDevices for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get exposes the hardware OATH method in the directory.
// returns a HardwareOathTokenAuthenticationMethodDeviceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareOathTokenAuthenticationMethodDeviceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property hardwareOathDevices in users
// returns a HardwareOathTokenAuthenticationMethodDeviceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareOathTokenAuthenticationMethodDeviceable, requestConfiguration *ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareOathTokenAuthenticationMethodDeviceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property hardwareOathDevices for users
// returns a *RequestInformation when successful
func (m *ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation exposes the hardware OATH method in the directory.
// returns a *RequestInformation when successful
func (m *ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property hardwareOathDevices in users
// returns a *RequestInformation when successful
func (m *ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareOathTokenAuthenticationMethodDeviceable, requestConfiguration *ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder when successful
func (m *ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder) WithUrl(rawUrl string)(*ItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder) {
    return NewItemAuthenticationHardwareOathMethodsItemDeviceHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
