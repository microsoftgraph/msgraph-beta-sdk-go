package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder provides operations to manage the hardwareOathDevices property of the microsoft.graph.authenticationMethodDevice entity.
type AuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilderGetQueryParameters read the properties and relationships of a hardwareOathTokenAuthenticationMethodDevice object.
type AuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilderGetQueryParameters
}
// AuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AssignTo provides operations to manage the assignTo property of the microsoft.graph.hardwareOathTokenAuthenticationMethodDevice entity.
// returns a *AuthenticationMethodDevicesHardwareOathDevicesItemAssignToRequestBuilder when successful
func (m *AuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder) AssignTo()(*AuthenticationMethodDevicesHardwareOathDevicesItemAssignToRequestBuilder) {
    return NewAuthenticationMethodDevicesHardwareOathDevicesItemAssignToRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewAuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilderInternal instantiates a new AuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder and sets the default values.
func NewAuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder) {
    m := &AuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/authenticationMethodDevices/hardwareOathDevices/{hardwareOathTokenAuthenticationMethodDevice%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder instantiates a new AuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder and sets the default values.
func NewAuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a Hardware OATH token. Token needs to be unassigned.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/authenticationmethoddevice-delete-hardwareoathdevices?view=graph-rest-beta
func (m *AuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a hardwareOathTokenAuthenticationMethodDevice object.
// returns a HardwareOathTokenAuthenticationMethodDeviceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/hardwareoathtokenauthenticationmethoddevice-get?view=graph-rest-beta
func (m *AuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareOathTokenAuthenticationMethodDeviceable, error) {
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
// Patch update the properties of a hardwareOathTokenAuthenticationMethodDevice object. The token needs to unassigned.
// returns a HardwareOathTokenAuthenticationMethodDeviceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/hardwareoathtokenauthenticationmethoddevice-update?view=graph-rest-beta
func (m *AuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareOathTokenAuthenticationMethodDeviceable, requestConfiguration *AuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareOathTokenAuthenticationMethodDeviceable, error) {
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
// ToDeleteRequestInformation delete a Hardware OATH token. Token needs to be unassigned.
// returns a *RequestInformation when successful
func (m *AuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a hardwareOathTokenAuthenticationMethodDevice object.
// returns a *RequestInformation when successful
func (m *AuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a hardwareOathTokenAuthenticationMethodDevice object. The token needs to unassigned.
// returns a *RequestInformation when successful
func (m *AuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareOathTokenAuthenticationMethodDeviceable, requestConfiguration *AuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder when successful
func (m *AuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder) WithUrl(rawUrl string)(*AuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder) {
    return NewAuthenticationMethodDevicesHardwareOathDevicesHardwareOathTokenAuthenticationMethodDeviceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
