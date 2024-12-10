package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemAuthenticationHardwareOathMethodsHardwareOathAuthenticationMethodItemRequestBuilder provides operations to manage the hardwareOathMethods property of the microsoft.graph.authentication entity.
type ItemAuthenticationHardwareOathMethodsHardwareOathAuthenticationMethodItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAuthenticationHardwareOathMethodsHardwareOathAuthenticationMethodItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationHardwareOathMethodsHardwareOathAuthenticationMethodItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemAuthenticationHardwareOathMethodsHardwareOathAuthenticationMethodItemRequestBuilderGetQueryParameters the hardware OATH time-based one-time password (TOTP) devices assigned to a user for authentication.
type ItemAuthenticationHardwareOathMethodsHardwareOathAuthenticationMethodItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemAuthenticationHardwareOathMethodsHardwareOathAuthenticationMethodItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationHardwareOathMethodsHardwareOathAuthenticationMethodItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAuthenticationHardwareOathMethodsHardwareOathAuthenticationMethodItemRequestBuilderGetQueryParameters
}
// Activate provides operations to call the activate method.
// returns a *ItemAuthenticationHardwareOathMethodsItemActivateRequestBuilder when successful
func (m *ItemAuthenticationHardwareOathMethodsHardwareOathAuthenticationMethodItemRequestBuilder) Activate()(*ItemAuthenticationHardwareOathMethodsItemActivateRequestBuilder) {
    return NewItemAuthenticationHardwareOathMethodsItemActivateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemAuthenticationHardwareOathMethodsHardwareOathAuthenticationMethodItemRequestBuilderInternal instantiates a new ItemAuthenticationHardwareOathMethodsHardwareOathAuthenticationMethodItemRequestBuilder and sets the default values.
func NewItemAuthenticationHardwareOathMethodsHardwareOathAuthenticationMethodItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationHardwareOathMethodsHardwareOathAuthenticationMethodItemRequestBuilder) {
    m := &ItemAuthenticationHardwareOathMethodsHardwareOathAuthenticationMethodItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/authentication/hardwareOathMethods/{hardwareOathAuthenticationMethod%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemAuthenticationHardwareOathMethodsHardwareOathAuthenticationMethodItemRequestBuilder instantiates a new ItemAuthenticationHardwareOathMethodsHardwareOathAuthenticationMethodItemRequestBuilder and sets the default values.
func NewItemAuthenticationHardwareOathMethodsHardwareOathAuthenticationMethodItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationHardwareOathMethodsHardwareOathAuthenticationMethodItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAuthenticationHardwareOathMethodsHardwareOathAuthenticationMethodItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Deactivate provides operations to call the deactivate method.
// returns a *ItemAuthenticationHardwareOathMethodsItemDeactivateRequestBuilder when successful
func (m *ItemAuthenticationHardwareOathMethodsHardwareOathAuthenticationMethodItemRequestBuilder) Deactivate()(*ItemAuthenticationHardwareOathMethodsItemDeactivateRequestBuilder) {
    return NewItemAuthenticationHardwareOathMethodsItemDeactivateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property hardwareOathMethods for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAuthenticationHardwareOathMethodsHardwareOathAuthenticationMethodItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemAuthenticationHardwareOathMethodsHardwareOathAuthenticationMethodItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Device provides operations to manage the device property of the microsoft.graph.hardwareOathAuthenticationMethod entity.
// returns a *ItemAuthenticationHardwareOathMethodsItemDeviceRequestBuilder when successful
func (m *ItemAuthenticationHardwareOathMethodsHardwareOathAuthenticationMethodItemRequestBuilder) Device()(*ItemAuthenticationHardwareOathMethodsItemDeviceRequestBuilder) {
    return NewItemAuthenticationHardwareOathMethodsItemDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the hardware OATH time-based one-time password (TOTP) devices assigned to a user for authentication.
// returns a HardwareOathAuthenticationMethodable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAuthenticationHardwareOathMethodsHardwareOathAuthenticationMethodItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAuthenticationHardwareOathMethodsHardwareOathAuthenticationMethodItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareOathAuthenticationMethodable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateHardwareOathAuthenticationMethodFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareOathAuthenticationMethodable), nil
}
// ToDeleteRequestInformation delete navigation property hardwareOathMethods for users
// returns a *RequestInformation when successful
func (m *ItemAuthenticationHardwareOathMethodsHardwareOathAuthenticationMethodItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemAuthenticationHardwareOathMethodsHardwareOathAuthenticationMethodItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the hardware OATH time-based one-time password (TOTP) devices assigned to a user for authentication.
// returns a *RequestInformation when successful
func (m *ItemAuthenticationHardwareOathMethodsHardwareOathAuthenticationMethodItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAuthenticationHardwareOathMethodsHardwareOathAuthenticationMethodItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemAuthenticationHardwareOathMethodsHardwareOathAuthenticationMethodItemRequestBuilder when successful
func (m *ItemAuthenticationHardwareOathMethodsHardwareOathAuthenticationMethodItemRequestBuilder) WithUrl(rawUrl string)(*ItemAuthenticationHardwareOathMethodsHardwareOathAuthenticationMethodItemRequestBuilder) {
    return NewItemAuthenticationHardwareOathMethodsHardwareOathAuthenticationMethodItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
