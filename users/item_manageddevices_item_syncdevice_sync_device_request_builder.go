package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemManageddevicesItemSyncdeviceSyncDeviceRequestBuilder provides operations to call the syncDevice method.
type ItemManageddevicesItemSyncdeviceSyncDeviceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManageddevicesItemSyncdeviceSyncDeviceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesItemSyncdeviceSyncDeviceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemManageddevicesItemSyncdeviceSyncDeviceRequestBuilderInternal instantiates a new ItemManageddevicesItemSyncdeviceSyncDeviceRequestBuilder and sets the default values.
func NewItemManageddevicesItemSyncdeviceSyncDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemSyncdeviceSyncDeviceRequestBuilder) {
    m := &ItemManageddevicesItemSyncdeviceSyncDeviceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}/syncDevice", pathParameters),
    }
    return m
}
// NewItemManageddevicesItemSyncdeviceSyncDeviceRequestBuilder instantiates a new ItemManageddevicesItemSyncdeviceSyncDeviceRequestBuilder and sets the default values.
func NewItemManageddevicesItemSyncdeviceSyncDeviceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemSyncdeviceSyncDeviceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManageddevicesItemSyncdeviceSyncDeviceRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action syncDevice
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManageddevicesItemSyncdeviceSyncDeviceRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemManageddevicesItemSyncdeviceSyncDeviceRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation invoke action syncDevice
// returns a *RequestInformation when successful
func (m *ItemManageddevicesItemSyncdeviceSyncDeviceRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemManageddevicesItemSyncdeviceSyncDeviceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemManageddevicesItemSyncdeviceSyncDeviceRequestBuilder when successful
func (m *ItemManageddevicesItemSyncdeviceSyncDeviceRequestBuilder) WithUrl(rawUrl string)(*ItemManageddevicesItemSyncdeviceSyncDeviceRequestBuilder) {
    return NewItemManageddevicesItemSyncdeviceSyncDeviceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
