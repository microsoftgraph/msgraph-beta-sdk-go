package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemManagedDevicesItemPauseConfigurationRefreshRequestBuilder provides operations to call the pauseConfigurationRefresh method.
type ItemManagedDevicesItemPauseConfigurationRefreshRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManagedDevicesItemPauseConfigurationRefreshRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManagedDevicesItemPauseConfigurationRefreshRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemManagedDevicesItemPauseConfigurationRefreshRequestBuilderInternal instantiates a new PauseConfigurationRefreshRequestBuilder and sets the default values.
func NewItemManagedDevicesItemPauseConfigurationRefreshRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesItemPauseConfigurationRefreshRequestBuilder) {
    m := &ItemManagedDevicesItemPauseConfigurationRefreshRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}/pauseConfigurationRefresh", pathParameters),
    }
    return m
}
// NewItemManagedDevicesItemPauseConfigurationRefreshRequestBuilder instantiates a new PauseConfigurationRefreshRequestBuilder and sets the default values.
func NewItemManagedDevicesItemPauseConfigurationRefreshRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesItemPauseConfigurationRefreshRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManagedDevicesItemPauseConfigurationRefreshRequestBuilderInternal(urlParams, requestAdapter)
}
// Post initiates a command to pause config refresh for the device.
func (m *ItemManagedDevicesItemPauseConfigurationRefreshRequestBuilder) Post(ctx context.Context, body ItemManagedDevicesItemPauseConfigurationRefreshPostRequestBodyable, requestConfiguration *ItemManagedDevicesItemPauseConfigurationRefreshRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation initiates a command to pause config refresh for the device.
func (m *ItemManagedDevicesItemPauseConfigurationRefreshRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemManagedDevicesItemPauseConfigurationRefreshPostRequestBodyable, requestConfiguration *ItemManagedDevicesItemPauseConfigurationRefreshRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemManagedDevicesItemPauseConfigurationRefreshRequestBuilder) WithUrl(rawUrl string)(*ItemManagedDevicesItemPauseConfigurationRefreshRequestBuilder) {
    return NewItemManagedDevicesItemPauseConfigurationRefreshRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
