package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceconfigurationsGettargetedusersanddevicesGetTargetedUsersAndDevicesRequestBuilder provides operations to call the getTargetedUsersAndDevices method.
type DeviceconfigurationsGettargetedusersanddevicesGetTargetedUsersAndDevicesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceconfigurationsGettargetedusersanddevicesGetTargetedUsersAndDevicesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceconfigurationsGettargetedusersanddevicesGetTargetedUsersAndDevicesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceconfigurationsGettargetedusersanddevicesGetTargetedUsersAndDevicesRequestBuilderInternal instantiates a new DeviceconfigurationsGettargetedusersanddevicesGetTargetedUsersAndDevicesRequestBuilder and sets the default values.
func NewDeviceconfigurationsGettargetedusersanddevicesGetTargetedUsersAndDevicesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceconfigurationsGettargetedusersanddevicesGetTargetedUsersAndDevicesRequestBuilder) {
    m := &DeviceconfigurationsGettargetedusersanddevicesGetTargetedUsersAndDevicesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceConfigurations/getTargetedUsersAndDevices", pathParameters),
    }
    return m
}
// NewDeviceconfigurationsGettargetedusersanddevicesGetTargetedUsersAndDevicesRequestBuilder instantiates a new DeviceconfigurationsGettargetedusersanddevicesGetTargetedUsersAndDevicesRequestBuilder and sets the default values.
func NewDeviceconfigurationsGettargetedusersanddevicesGetTargetedUsersAndDevicesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceconfigurationsGettargetedusersanddevicesGetTargetedUsersAndDevicesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceconfigurationsGettargetedusersanddevicesGetTargetedUsersAndDevicesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getTargetedUsersAndDevices
// Deprecated: This method is obsolete. Use PostAsGetTargetedUsersAndDevicesPostResponse instead.
// returns a DeviceconfigurationsGettargetedusersanddevicesGetTargetedUsersAndDevicesResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceconfigurationsGettargetedusersanddevicesGetTargetedUsersAndDevicesRequestBuilder) Post(ctx context.Context, body DeviceconfigurationsGettargetedusersanddevicesGetTargetedUsersAndDevicesPostRequestBodyable, requestConfiguration *DeviceconfigurationsGettargetedusersanddevicesGetTargetedUsersAndDevicesRequestBuilderPostRequestConfiguration)(DeviceconfigurationsGettargetedusersanddevicesGetTargetedUsersAndDevicesResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeviceconfigurationsGettargetedusersanddevicesGetTargetedUsersAndDevicesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeviceconfigurationsGettargetedusersanddevicesGetTargetedUsersAndDevicesResponseable), nil
}
// PostAsGetTargetedUsersAndDevicesPostResponse invoke action getTargetedUsersAndDevices
// returns a DeviceconfigurationsGettargetedusersanddevicesGetTargetedUsersAndDevicesPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceconfigurationsGettargetedusersanddevicesGetTargetedUsersAndDevicesRequestBuilder) PostAsGetTargetedUsersAndDevicesPostResponse(ctx context.Context, body DeviceconfigurationsGettargetedusersanddevicesGetTargetedUsersAndDevicesPostRequestBodyable, requestConfiguration *DeviceconfigurationsGettargetedusersanddevicesGetTargetedUsersAndDevicesRequestBuilderPostRequestConfiguration)(DeviceconfigurationsGettargetedusersanddevicesGetTargetedUsersAndDevicesPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeviceconfigurationsGettargetedusersanddevicesGetTargetedUsersAndDevicesPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeviceconfigurationsGettargetedusersanddevicesGetTargetedUsersAndDevicesPostResponseable), nil
}
// ToPostRequestInformation invoke action getTargetedUsersAndDevices
// returns a *RequestInformation when successful
func (m *DeviceconfigurationsGettargetedusersanddevicesGetTargetedUsersAndDevicesRequestBuilder) ToPostRequestInformation(ctx context.Context, body DeviceconfigurationsGettargetedusersanddevicesGetTargetedUsersAndDevicesPostRequestBodyable, requestConfiguration *DeviceconfigurationsGettargetedusersanddevicesGetTargetedUsersAndDevicesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DeviceconfigurationsGettargetedusersanddevicesGetTargetedUsersAndDevicesRequestBuilder when successful
func (m *DeviceconfigurationsGettargetedusersanddevicesGetTargetedUsersAndDevicesRequestBuilder) WithUrl(rawUrl string)(*DeviceconfigurationsGettargetedusersanddevicesGetTargetedUsersAndDevicesRequestBuilder) {
    return NewDeviceconfigurationsGettargetedusersanddevicesGetTargetedUsersAndDevicesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
